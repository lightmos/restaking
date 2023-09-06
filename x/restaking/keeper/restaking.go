package keeper

import (
	"errors"
	"github.com/lightmos/restaking/x/restaking/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	staketype "github.com/cosmos/cosmos-sdk/x/staking/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

func (k Keeper) RestakeUndelegate(ctx sdk.Context) []abci.ValidatorUpdate {
	matureUnbonds := k.DequeueAllMatureUBDQueue(ctx, ctx.BlockHeader().Time)
	for _, dvPair := range matureUnbonds {
		addr, err := sdk.ValAddressFromBech32(dvPair.ValidatorAddress)
		if err != nil {
			panic(err)
		}
		delegatorAddress := sdk.MustAccAddressFromBech32(dvPair.DelegatorAddress)

		shareVal, found := k.GetValidatorToken(ctx, delegatorAddress.String())

		if !found {
			continue
		}

		balances, err := k.CompleteShareUnbonding(ctx, delegatorAddress, addr)
		if err != nil {
			continue
		}

		for _, balance := range balances {
			if balance.Denom != k.stakingKeeper.BondDenom(ctx) {
				continue
			}
			del, retireToken := k.DescHistory(ctx, balance.Denom, "token", delegatorAddress.String(), int32(balance.Amount.Int64()))
			if !del {
				continue
			}

			shareVal.Total.Amount = shareVal.Total.Amount.Sub(balance.Amount)
			shareVal.Retire.Amount = shareVal.Retire.Amount.Add(balance.Amount)
			shareVal.Available.Amount = shareVal.Available.Amount.Add(sdk.NewInt(int64(retireToken)))
			k.SetValidatorToken(ctx, shareVal)
		}
	}
	return []abci.ValidatorUpdate{}
}

// TransmitRestakingPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitRestakingPacket(
	ctx sdk.Context,
	packetData types.RestakePacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()

	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvRestakePacket processes packet reception
func (k Keeper) OnRecvRestakePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RestakePacketData) (packetAck types.RestakePacketDataAck, err error) {
	goctx := sdk.UnwrapSDKContext(ctx)
	logger := k.Logger(goctx)

	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	var pk cryptotypes.PubKey
	if err := k.cdc.UnmarshalInterfaceJSON([]byte(data.Pubkey), &pk); err != nil {
		return packetAck, err
	}

	//var pkAny *codectypes.Any
	//if pk != nil {
	//	var err error
	//	if pkAny, err = codectypes.NewAnyWithValue(pk); err != nil {
	//		return packetAck, err
	//	}
	//}

	packetAck.Succeed = false

	// mint token
	destDenomFromVocher, flg := k.OriginalDenom(ctx, packet.DestinationPort, packet.DestinationChannel, data.Value.Denom)
	if !flg {
		return packetAck, errors.New("invalid denom")
	}
	restaker, err := sdk.AccAddressFromBech32(data.Restaker)
	if err != nil {
		return packetAck, err
	}
	logger.Info("carver|recv restake packet", "restaker", restaker, "denom", data.Value.Denom,
		"destDenomFromVocher", destDenomFromVocher, "data", data)

	err = k.MintTokens(ctx, restaker, sdk.NewCoin(destDenomFromVocher, data.Value.Amount))
	if err != nil {
		return packetAck, err
	}

	// restake validator
	//cosmosValidator := &stakingtypes.MsgCreateValidator{
	//	Description:       stakingtypes.Description(data.Description),
	//	Commission:        stakingtypes.CommissionRates(data.Commission),
	//	MinSelfDelegation: data.MinSelfDelegation,
	//	DelegatorAddress:  data.DelegatorAddress,
	//	ValidatorAddress:  data.ValidatorAddress,
	//	Pubkey:            pkAny,
	//	Value:             sdk.NewCoin(destDenomFromVocher, data.Value.Amount),
	//}

	// ## simple test restakeValidator ##
	//_, err = k.stakingKeeper.CreateValidator(ctx, cosmosValidator)
	if data.Commission.Rate.LT(k.stakingKeeper.MinCommissionRate(ctx)) {
		return packetAck, sdkerrors.Wrapf(staketype.ErrCommissionLTMinRate, "cannot set validator commission to less than minimum rate of %s", k.stakingKeeper.MinCommissionRate(ctx))
	}

	valAddr, _ := sdk.ValAddressFromBech32(data.ValidatorAddress)
	// check to see if the pubkey or sender has been registered before
	if _, found := k.stakingKeeper.GetValidator(ctx, valAddr); found {
		return packetAck, staketype.ErrValidatorOwnerExists
	}

	if _, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(pk)); found {
		return packetAck, staketype.ErrValidatorPubKeyExists
	}

	bondDenom := k.stakingKeeper.BondDenom(ctx)
	if destDenomFromVocher != bondDenom {
		return packetAck, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest, "invalid coin denomination: got %s, expected %s", data.Value.Denom, bondDenom,
		)
	}

	if _, err := data.Description.EnsureLength(); err != nil {
		return packetAck, err
	}

	cp := ctx.ConsensusParams()
	if cp != nil && cp.Validator != nil {
		pkType := pk.Type()
		hasKeyType := false
		for _, keyType := range cp.Validator.PubKeyTypes {
			if pkType == keyType {
				hasKeyType = true
				break
			}
		}
		if !hasKeyType {
			return packetAck, sdkerrors.Wrapf(
				staketype.ErrValidatorPubKeyTypeNotSupported,
				"got: %s, expected: %s", pk.Type(), cp.Validator.PubKeyTypes,
			)
		}
	}

	des := staketype.Description{
		Moniker:         data.Description.Moniker,
		Identity:        data.Description.Identity,
		Website:         data.Description.Website,
		SecurityContact: data.Description.SecurityContact,
		Details:         data.Description.Details,
	}
	validator, err := staketype.NewValidator(valAddr, pk, des)
	if err != nil {
		return packetAck, err
	}

	commission := staketype.NewCommissionWithTime(
		data.Commission.Rate, data.Commission.MaxRate,
		data.Commission.MaxChangeRate, ctx.BlockHeader().Time,
	)

	validator, err = validator.SetInitialCommission(commission)
	if err != nil {
		return packetAck, err
	}

	validator.MinSelfDelegation = data.MinSelfDelegation

	k.stakingKeeper.SetValidator(ctx, validator)
	k.stakingKeeper.SetValidatorByConsAddr(ctx, validator)
	k.stakingKeeper.SetNewValidatorByPowerIndex(ctx, validator)

	// call the after-creation hook
	if err := k.stakingKeeper.Hooks().AfterValidatorCreated(ctx, validator.GetOperator()); err != nil {
		return packetAck, err
	}

	// move coins from the msg.Address account to a (self-delegation) delegator account
	// the validator account and global shares are updated within here
	// NOTE source will always be from a wallet which are unbonded
	_, err = k.stakingKeeper.Delegate(ctx, restaker, data.Value.Amount, staketype.Unbonded, validator, true)
	if err != nil {
		return packetAck, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			staketype.EventTypeCreateValidator,
			sdk.NewAttribute(staketype.AttributeKeyValidator, data.ValidatorAddress),
			sdk.NewAttribute(sdk.AttributeKeyAmount, data.Value.String()),
		),
	})

	if err != nil {
		logger.Error("carver|restake Validator err", "err", err.Error())
		// if restake fail, burn tokens
		k.BurnTokens(ctx, restaker, sdk.NewCoin(data.Value.Denom, data.Value.Amount))
		return packetAck, err
	}

	totalCoin := sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), data.Value.Amount)
	retireCoin := sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), sdk.ZeroInt())
	avaCoin := sdk.NewCoin("token", sdk.ZeroInt())
	vt := types.ValidatorToken{
		Address:   data.Restaker,
		Total:     totalCoin,
		Retire:    retireCoin,
		Available: avaCoin,
	}
	k.SetValidatorToken(ctx, vt)
	logger.Info("carver|recv restake handle succeed", "restaker", restaker, "denom", data.Value.Denom)
	packetAck.Succeed = true
	return packetAck, nil
}

// OnAcknowledgementRestakePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementRestakePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RestakePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		ctx.Logger().Info("caver|OnAcknowledgementRestakePacket err")
		return k.refundPacketToken(ctx, packet, data)
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.RestakePacketDataAck
		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		ctx.Logger().Info("caver|OnAcknowledgementRestakePacket succeed")
		// save restake validator trace
		k.SetRestakeValidatorTrace(ctx, data.Restaker, data.DestinationChainId)
		return nil
	default:
		return nil
	}
}

func (k Keeper) refundPacketToken(ctx sdk.Context, packet channeltypes.Packet, data types.RestakePacketData) error {
	// In case of error we unlock the native token
	receiver, err := sdk.AccAddressFromBech32(data.Restaker)
	if err != nil {
		return err
	}

	if err := k.UnlockTokens(
		ctx,
		packet.SourcePort,
		packet.SourceChannel,
		receiver,
		sdk.Coin(data.Value),
	); err != nil {
		return err
	}

	return nil
}

// OnTimeoutRestakePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutRestakePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RestakePacketData) error {
	return k.refundPacketToken(ctx, packet, data)
}
