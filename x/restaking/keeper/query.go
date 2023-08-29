package keeper

import (
	"github.com/lightmos/restaking/x/restaking/types"
)

var _ types.QueryServer = Keeper{}
