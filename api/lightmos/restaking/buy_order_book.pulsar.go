// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package restaking

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_BuyOrderBook             protoreflect.MessageDescriptor
	fd_BuyOrderBook_index       protoreflect.FieldDescriptor
	fd_BuyOrderBook_amountDenom protoreflect.FieldDescriptor
	fd_BuyOrderBook_priceDenom  protoreflect.FieldDescriptor
	fd_BuyOrderBook_book        protoreflect.FieldDescriptor
)

func init() {
	file_lightmos_restaking_buy_order_book_proto_init()
	md_BuyOrderBook = File_lightmos_restaking_buy_order_book_proto.Messages().ByName("BuyOrderBook")
	fd_BuyOrderBook_index = md_BuyOrderBook.Fields().ByName("index")
	fd_BuyOrderBook_amountDenom = md_BuyOrderBook.Fields().ByName("amountDenom")
	fd_BuyOrderBook_priceDenom = md_BuyOrderBook.Fields().ByName("priceDenom")
	fd_BuyOrderBook_book = md_BuyOrderBook.Fields().ByName("book")
}

var _ protoreflect.Message = (*fastReflection_BuyOrderBook)(nil)

type fastReflection_BuyOrderBook BuyOrderBook

func (x *BuyOrderBook) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BuyOrderBook)(x)
}

func (x *BuyOrderBook) slowProtoReflect() protoreflect.Message {
	mi := &file_lightmos_restaking_buy_order_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BuyOrderBook_messageType fastReflection_BuyOrderBook_messageType
var _ protoreflect.MessageType = fastReflection_BuyOrderBook_messageType{}

type fastReflection_BuyOrderBook_messageType struct{}

func (x fastReflection_BuyOrderBook_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BuyOrderBook)(nil)
}
func (x fastReflection_BuyOrderBook_messageType) New() protoreflect.Message {
	return new(fastReflection_BuyOrderBook)
}
func (x fastReflection_BuyOrderBook_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrderBook
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BuyOrderBook) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrderBook
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BuyOrderBook) Type() protoreflect.MessageType {
	return _fastReflection_BuyOrderBook_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BuyOrderBook) New() protoreflect.Message {
	return new(fastReflection_BuyOrderBook)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BuyOrderBook) Interface() protoreflect.ProtoMessage {
	return (*BuyOrderBook)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BuyOrderBook) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Index != "" {
		value := protoreflect.ValueOfString(x.Index)
		if !f(fd_BuyOrderBook_index, value) {
			return
		}
	}
	if x.AmountDenom != "" {
		value := protoreflect.ValueOfString(x.AmountDenom)
		if !f(fd_BuyOrderBook_amountDenom, value) {
			return
		}
	}
	if x.PriceDenom != "" {
		value := protoreflect.ValueOfString(x.PriceDenom)
		if !f(fd_BuyOrderBook_priceDenom, value) {
			return
		}
	}
	if x.Book != nil {
		value := protoreflect.ValueOfMessage(x.Book.ProtoReflect())
		if !f(fd_BuyOrderBook_book, value) {
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
func (x *fastReflection_BuyOrderBook) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "lightmos.restaking.BuyOrderBook.index":
		return x.Index != ""
	case "lightmos.restaking.BuyOrderBook.amountDenom":
		return x.AmountDenom != ""
	case "lightmos.restaking.BuyOrderBook.priceDenom":
		return x.PriceDenom != ""
	case "lightmos.restaking.BuyOrderBook.book":
		return x.Book != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.BuyOrderBook"))
		}
		panic(fmt.Errorf("message lightmos.restaking.BuyOrderBook does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrderBook) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "lightmos.restaking.BuyOrderBook.index":
		x.Index = ""
	case "lightmos.restaking.BuyOrderBook.amountDenom":
		x.AmountDenom = ""
	case "lightmos.restaking.BuyOrderBook.priceDenom":
		x.PriceDenom = ""
	case "lightmos.restaking.BuyOrderBook.book":
		x.Book = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.BuyOrderBook"))
		}
		panic(fmt.Errorf("message lightmos.restaking.BuyOrderBook does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BuyOrderBook) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "lightmos.restaking.BuyOrderBook.index":
		value := x.Index
		return protoreflect.ValueOfString(value)
	case "lightmos.restaking.BuyOrderBook.amountDenom":
		value := x.AmountDenom
		return protoreflect.ValueOfString(value)
	case "lightmos.restaking.BuyOrderBook.priceDenom":
		value := x.PriceDenom
		return protoreflect.ValueOfString(value)
	case "lightmos.restaking.BuyOrderBook.book":
		value := x.Book
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.BuyOrderBook"))
		}
		panic(fmt.Errorf("message lightmos.restaking.BuyOrderBook does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BuyOrderBook) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "lightmos.restaking.BuyOrderBook.index":
		x.Index = value.Interface().(string)
	case "lightmos.restaking.BuyOrderBook.amountDenom":
		x.AmountDenom = value.Interface().(string)
	case "lightmos.restaking.BuyOrderBook.priceDenom":
		x.PriceDenom = value.Interface().(string)
	case "lightmos.restaking.BuyOrderBook.book":
		x.Book = value.Message().Interface().(*OrderBook)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.BuyOrderBook"))
		}
		panic(fmt.Errorf("message lightmos.restaking.BuyOrderBook does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BuyOrderBook) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "lightmos.restaking.BuyOrderBook.book":
		if x.Book == nil {
			x.Book = new(OrderBook)
		}
		return protoreflect.ValueOfMessage(x.Book.ProtoReflect())
	case "lightmos.restaking.BuyOrderBook.index":
		panic(fmt.Errorf("field index of message lightmos.restaking.BuyOrderBook is not mutable"))
	case "lightmos.restaking.BuyOrderBook.amountDenom":
		panic(fmt.Errorf("field amountDenom of message lightmos.restaking.BuyOrderBook is not mutable"))
	case "lightmos.restaking.BuyOrderBook.priceDenom":
		panic(fmt.Errorf("field priceDenom of message lightmos.restaking.BuyOrderBook is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.BuyOrderBook"))
		}
		panic(fmt.Errorf("message lightmos.restaking.BuyOrderBook does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BuyOrderBook) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "lightmos.restaking.BuyOrderBook.index":
		return protoreflect.ValueOfString("")
	case "lightmos.restaking.BuyOrderBook.amountDenom":
		return protoreflect.ValueOfString("")
	case "lightmos.restaking.BuyOrderBook.priceDenom":
		return protoreflect.ValueOfString("")
	case "lightmos.restaking.BuyOrderBook.book":
		m := new(OrderBook)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: lightmos.restaking.BuyOrderBook"))
		}
		panic(fmt.Errorf("message lightmos.restaking.BuyOrderBook does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BuyOrderBook) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in lightmos.restaking.BuyOrderBook", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BuyOrderBook) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrderBook) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BuyOrderBook) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BuyOrderBook) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BuyOrderBook)
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
		l = len(x.Index)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AmountDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.PriceDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Book != nil {
			l = options.Size(x.Book)
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
		x := input.Message.Interface().(*BuyOrderBook)
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
		if x.Book != nil {
			encoded, err := options.Marshal(x.Book)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.PriceDenom) > 0 {
			i -= len(x.PriceDenom)
			copy(dAtA[i:], x.PriceDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PriceDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.AmountDenom) > 0 {
			i -= len(x.AmountDenom)
			copy(dAtA[i:], x.AmountDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AmountDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Index) > 0 {
			i -= len(x.Index)
			copy(dAtA[i:], x.Index)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Index)))
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
		x := input.Message.Interface().(*BuyOrderBook)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrderBook: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
				x.Index = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
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
				x.AmountDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
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
				x.PriceDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Book", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Book == nil {
					x.Book = &OrderBook{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Book); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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
// source: lightmos/restaking/buy_order_book.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BuyOrderBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index       string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	AmountDenom string     `protobuf:"bytes,2,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	PriceDenom  string     `protobuf:"bytes,3,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	Book        *OrderBook `protobuf:"bytes,4,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *BuyOrderBook) Reset() {
	*x = BuyOrderBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightmos_restaking_buy_order_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrderBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrderBook) ProtoMessage() {}

// Deprecated: Use BuyOrderBook.ProtoReflect.Descriptor instead.
func (*BuyOrderBook) Descriptor() ([]byte, []int) {
	return file_lightmos_restaking_buy_order_book_proto_rawDescGZIP(), []int{0}
}

func (x *BuyOrderBook) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *BuyOrderBook) GetAmountDenom() string {
	if x != nil {
		return x.AmountDenom
	}
	return ""
}

func (x *BuyOrderBook) GetPriceDenom() string {
	if x != nil {
		return x.PriceDenom
	}
	return ""
}

func (x *BuyOrderBook) GetBook() *OrderBook {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_lightmos_restaking_buy_order_book_proto protoreflect.FileDescriptor

var file_lightmos_restaking_buy_order_book_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x6d, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x1e, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01,
	0x0a, 0x0c, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x44,
	0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x31, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x42, 0xca, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x42, 0x11, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x6d, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0xa2,
	0x02, 0x03, 0x4c, 0x52, 0x58, 0xaa, 0x02, 0x12, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x12, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x6d, 0x6f, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0xe2,
	0x02, 0x1e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x13, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lightmos_restaking_buy_order_book_proto_rawDescOnce sync.Once
	file_lightmos_restaking_buy_order_book_proto_rawDescData = file_lightmos_restaking_buy_order_book_proto_rawDesc
)

func file_lightmos_restaking_buy_order_book_proto_rawDescGZIP() []byte {
	file_lightmos_restaking_buy_order_book_proto_rawDescOnce.Do(func() {
		file_lightmos_restaking_buy_order_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_lightmos_restaking_buy_order_book_proto_rawDescData)
	})
	return file_lightmos_restaking_buy_order_book_proto_rawDescData
}

var file_lightmos_restaking_buy_order_book_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lightmos_restaking_buy_order_book_proto_goTypes = []interface{}{
	(*BuyOrderBook)(nil), // 0: lightmos.restaking.BuyOrderBook
	(*OrderBook)(nil),    // 1: lightmos.restaking.OrderBook
}
var file_lightmos_restaking_buy_order_book_proto_depIdxs = []int32{
	1, // 0: lightmos.restaking.BuyOrderBook.book:type_name -> lightmos.restaking.OrderBook
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lightmos_restaking_buy_order_book_proto_init() }
func file_lightmos_restaking_buy_order_book_proto_init() {
	if File_lightmos_restaking_buy_order_book_proto != nil {
		return
	}
	file_lightmos_restaking_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lightmos_restaking_buy_order_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrderBook); i {
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
			RawDescriptor: file_lightmos_restaking_buy_order_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lightmos_restaking_buy_order_book_proto_goTypes,
		DependencyIndexes: file_lightmos_restaking_buy_order_book_proto_depIdxs,
		MessageInfos:      file_lightmos_restaking_buy_order_book_proto_msgTypes,
	}.Build()
	File_lightmos_restaking_buy_order_book_proto = out.File
	file_lightmos_restaking_buy_order_book_proto_rawDesc = nil
	file_lightmos_restaking_buy_order_book_proto_goTypes = nil
	file_lightmos_restaking_buy_order_book_proto_depIdxs = nil
}
