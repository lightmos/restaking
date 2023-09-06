// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package restaking

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ValidatorToken           protoreflect.MessageDescriptor
	fd_ValidatorToken_address   protoreflect.FieldDescriptor
	fd_ValidatorToken_total     protoreflect.FieldDescriptor
	fd_ValidatorToken_retire    protoreflect.FieldDescriptor
	fd_ValidatorToken_available protoreflect.FieldDescriptor
)

func init() {
	file_lightmos_restaking_validator_token_proto_init()
	md_ValidatorToken = File_lightmos_restaking_validator_token_proto.Messages().ByName("ValidatorToken")
	fd_ValidatorToken_address = md_ValidatorToken.Fields().ByName("address")
	fd_ValidatorToken_total = md_ValidatorToken.Fields().ByName("total")
	fd_ValidatorToken_retire = md_ValidatorToken.Fields().ByName("retire")
	fd_ValidatorToken_available = md_ValidatorToken.Fields().ByName("available")
}

var _ protoreflect.Message = (*fastReflection_ValidatorToken)(nil)

type fastReflection_ValidatorToken ValidatorToken

func (x *ValidatorToken) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorToken)(x)
}

func (x *ValidatorToken) slowProtoReflect() protoreflect.Message {
	mi := &file_lightmos_restaking_validator_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorToken_messageType fastReflection_ValidatorToken_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorToken_messageType{}

type fastReflection_ValidatorToken_messageType struct{}

func (x fastReflection_ValidatorToken_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorToken)(nil)
}
func (x fastReflection_ValidatorToken_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorToken)
}
func (x fastReflection_ValidatorToken_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorToken
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorToken) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorToken
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorToken) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorToken_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorToken) New() protoreflect.Message {
	return new(fastReflection_ValidatorToken)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorToken) Interface() protoreflect.ProtoMessage {
	return (*ValidatorToken)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorToken) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_ValidatorToken_address, value) {
			return
		}
	}
	if x.Total != "" {
		value := protoreflect.ValueOfString(x.Total)
		if !f(fd_ValidatorToken_total, value) {
			return
		}
	}
	if x.Retire != "" {
		value := protoreflect.ValueOfString(x.Retire)
		if !f(fd_ValidatorToken_retire, value) {
			return
		}
	}
	if x.Available != "" {
		value := protoreflect.ValueOfString(x.Available)
		if !f(fd_ValidatorToken_available, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorToken) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "lightmos.restaking.ValidatorToken.address":
		return x.Address != ""
	case "lightmos.restaking.ValidatorToken.total":
		return x.Total != ""
	case "lightmos.restaking.ValidatorToken.retire":
		return x.Retire != ""
	case "lightmos.restaking.ValidatorToken.available":
		return x.Available != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.ValidatorToken"))
		}
		panic(fmt.Errorf("message lightmos.restaking.ValidatorToken does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorToken) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "lightmos.restaking.ValidatorToken.address":
		x.Address = ""
	case "lightmos.restaking.ValidatorToken.total":
		x.Total = ""
	case "lightmos.restaking.ValidatorToken.retire":
		x.Retire = ""
	case "lightmos.restaking.ValidatorToken.available":
		x.Available = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.ValidatorToken"))
		}
		panic(fmt.Errorf("message lightmos.restaking.ValidatorToken does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorToken) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "lightmos.restaking.ValidatorToken.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "lightmos.restaking.ValidatorToken.total":
		value := x.Total
		return protoreflect.ValueOfString(value)
	case "lightmos.restaking.ValidatorToken.retire":
		value := x.Retire
		return protoreflect.ValueOfString(value)
	case "lightmos.restaking.ValidatorToken.available":
		value := x.Available
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.ValidatorToken"))
		}
		panic(fmt.Errorf("message lightmos.restaking.ValidatorToken does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorToken) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "lightmos.restaking.ValidatorToken.address":
		x.Address = value.Interface().(string)
	case "lightmos.restaking.ValidatorToken.total":
		x.Total = value.Interface().(string)
	case "lightmos.restaking.ValidatorToken.retire":
		x.Retire = value.Interface().(string)
	case "lightmos.restaking.ValidatorToken.available":
		x.Available = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.ValidatorToken"))
		}
		panic(fmt.Errorf("message lightmos.restaking.ValidatorToken does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorToken) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "lightmos.restaking.ValidatorToken.address":
		panic(fmt.Errorf("field address of message lightmos.restaking.ValidatorToken is not mutable"))
	case "lightmos.restaking.ValidatorToken.total":
		panic(fmt.Errorf("field total of message lightmos.restaking.ValidatorToken is not mutable"))
	case "lightmos.restaking.ValidatorToken.retire":
		panic(fmt.Errorf("field retire of message lightmos.restaking.ValidatorToken is not mutable"))
	case "lightmos.restaking.ValidatorToken.available":
		panic(fmt.Errorf("field available of message lightmos.restaking.ValidatorToken is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.ValidatorToken"))
		}
		panic(fmt.Errorf("message lightmos.restaking.ValidatorToken does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorToken) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "lightmos.restaking.ValidatorToken.address":
		return protoreflect.ValueOfString("")
	case "lightmos.restaking.ValidatorToken.total":
		return protoreflect.ValueOfString("")
	case "lightmos.restaking.ValidatorToken.retire":
		return protoreflect.ValueOfString("")
	case "lightmos.restaking.ValidatorToken.available":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.ValidatorToken"))
		}
		panic(fmt.Errorf("message lightmos.restaking.ValidatorToken does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorToken) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in lightmos.restaking.ValidatorToken", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorToken) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorToken) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorToken) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorToken) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorToken)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Total)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Retire)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Available)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorToken)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Available) > 0 {
			i -= len(x.Available)
			copy(dAtA[i:], x.Available)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Available)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Retire) > 0 {
			i -= len(x.Retire)
			copy(dAtA[i:], x.Retire)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Retire)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Total) > 0 {
			i -= len(x.Total)
			copy(dAtA[i:], x.Total)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Total)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorToken)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorToken: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorToken: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Total = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Retire", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Retire = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Available", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Available = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: lightmos/restaking/validator_token.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidatorToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Total     string `protobuf:"bytes,2,opt,name=total,proto3" json:"total,omitempty"`
	Retire    string `protobuf:"bytes,3,opt,name=retire,proto3" json:"retire,omitempty"`
	Available string `protobuf:"bytes,4,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *ValidatorToken) Reset() {
	*x = ValidatorToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightmos_restaking_validator_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorToken) ProtoMessage() {}

// Deprecated: Use ValidatorToken.ProtoReflect.Descriptor instead.
func (*ValidatorToken) Descriptor() ([]byte, []int) {
	return file_lightmos_restaking_validator_token_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorToken) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ValidatorToken) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *ValidatorToken) GetRetire() string {
	if x != nil {
		return x.Retire
	}
	return ""
}

func (x *ValidatorToken) GetAvailable() string {
	if x != nil {
		return x.Available
	}
	return ""
}

var File_lightmos_restaking_validator_token_proto protoreflect.FileDescriptor

var file_lightmos_restaking_validator_token_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x19,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x89, 0x02, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xc8, 0xde, 0x1f,
	0x00, 0xda, 0xde, 0x1f, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x47, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2f, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x69, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x4d, 0x0a, 0x09,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2f, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0xbb, 0x01, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0xa2, 0x02, 0x03, 0x4c, 0x52, 0x58, 0xaa, 0x02, 0x12, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x6d, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x12,
	0x4c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0xe2, 0x02, 0x1e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x5c, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lightmos_restaking_validator_token_proto_rawDescOnce sync.Once
	file_lightmos_restaking_validator_token_proto_rawDescData = file_lightmos_restaking_validator_token_proto_rawDesc
)

func file_lightmos_restaking_validator_token_proto_rawDescGZIP() []byte {
	file_lightmos_restaking_validator_token_proto_rawDescOnce.Do(func() {
		file_lightmos_restaking_validator_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_lightmos_restaking_validator_token_proto_rawDescData)
	})
	return file_lightmos_restaking_validator_token_proto_rawDescData
}

var file_lightmos_restaking_validator_token_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lightmos_restaking_validator_token_proto_goTypes = []interface{}{
	(*ValidatorToken)(nil), // 0: lightmos.restaking.ValidatorToken
}
var file_lightmos_restaking_validator_token_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lightmos_restaking_validator_token_proto_init() }
func file_lightmos_restaking_validator_token_proto_init() {
	if File_lightmos_restaking_validator_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lightmos_restaking_validator_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lightmos_restaking_validator_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lightmos_restaking_validator_token_proto_goTypes,
		DependencyIndexes: file_lightmos_restaking_validator_token_proto_depIdxs,
		MessageInfos:      file_lightmos_restaking_validator_token_proto_msgTypes,
	}.Build()
	File_lightmos_restaking_validator_token_proto = out.File
	file_lightmos_restaking_validator_token_proto_rawDesc = nil
	file_lightmos_restaking_validator_token_proto_goTypes = nil
	file_lightmos_restaking_validator_token_proto_depIdxs = nil
}
