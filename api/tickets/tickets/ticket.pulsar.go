// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package tickets

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
	md_Ticket           protoreflect.MessageDescriptor
	fd_Ticket_name      protoreflect.FieldDescriptor
	fd_Ticket_event     protoreflect.FieldDescriptor
	fd_Ticket_id        protoreflect.FieldDescriptor
	fd_Ticket_createdAt protoreflect.FieldDescriptor
)

func init() {
	file_tickets_tickets_ticket_proto_init()
	md_Ticket = File_tickets_tickets_ticket_proto.Messages().ByName("Ticket")
	fd_Ticket_name = md_Ticket.Fields().ByName("name")
	fd_Ticket_event = md_Ticket.Fields().ByName("event")
	fd_Ticket_id = md_Ticket.Fields().ByName("id")
	fd_Ticket_createdAt = md_Ticket.Fields().ByName("createdAt")
}

var _ protoreflect.Message = (*fastReflection_Ticket)(nil)

type fastReflection_Ticket Ticket

func (x *Ticket) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Ticket)(x)
}

func (x *Ticket) slowProtoReflect() protoreflect.Message {
	mi := &file_tickets_tickets_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Ticket_messageType fastReflection_Ticket_messageType
var _ protoreflect.MessageType = fastReflection_Ticket_messageType{}

type fastReflection_Ticket_messageType struct{}

func (x fastReflection_Ticket_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Ticket)(nil)
}
func (x fastReflection_Ticket_messageType) New() protoreflect.Message {
	return new(fastReflection_Ticket)
}
func (x fastReflection_Ticket_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Ticket
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Ticket) Descriptor() protoreflect.MessageDescriptor {
	return md_Ticket
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Ticket) Type() protoreflect.MessageType {
	return _fastReflection_Ticket_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Ticket) New() protoreflect.Message {
	return new(fastReflection_Ticket)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Ticket) Interface() protoreflect.ProtoMessage {
	return (*Ticket)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Ticket) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_Ticket_name, value) {
			return
		}
	}
	if x.Event != "" {
		value := protoreflect.ValueOfString(x.Event)
		if !f(fd_Ticket_event, value) {
			return
		}
	}
	if x.Id != "" {
		value := protoreflect.ValueOfString(x.Id)
		if !f(fd_Ticket_id, value) {
			return
		}
	}
	if x.CreatedAt != int64(0) {
		value := protoreflect.ValueOfInt64(x.CreatedAt)
		if !f(fd_Ticket_createdAt, value) {
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
func (x *fastReflection_Ticket) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tickets.tickets.Ticket.name":
		return x.Name != ""
	case "tickets.tickets.Ticket.event":
		return x.Event != ""
	case "tickets.tickets.Ticket.id":
		return x.Id != ""
	case "tickets.tickets.Ticket.createdAt":
		return x.CreatedAt != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tickets.tickets.Ticket"))
		}
		panic(fmt.Errorf("message tickets.tickets.Ticket does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Ticket) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tickets.tickets.Ticket.name":
		x.Name = ""
	case "tickets.tickets.Ticket.event":
		x.Event = ""
	case "tickets.tickets.Ticket.id":
		x.Id = ""
	case "tickets.tickets.Ticket.createdAt":
		x.CreatedAt = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tickets.tickets.Ticket"))
		}
		panic(fmt.Errorf("message tickets.tickets.Ticket does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Ticket) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tickets.tickets.Ticket.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "tickets.tickets.Ticket.event":
		value := x.Event
		return protoreflect.ValueOfString(value)
	case "tickets.tickets.Ticket.id":
		value := x.Id
		return protoreflect.ValueOfString(value)
	case "tickets.tickets.Ticket.createdAt":
		value := x.CreatedAt
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tickets.tickets.Ticket"))
		}
		panic(fmt.Errorf("message tickets.tickets.Ticket does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Ticket) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tickets.tickets.Ticket.name":
		x.Name = value.Interface().(string)
	case "tickets.tickets.Ticket.event":
		x.Event = value.Interface().(string)
	case "tickets.tickets.Ticket.id":
		x.Id = value.Interface().(string)
	case "tickets.tickets.Ticket.createdAt":
		x.CreatedAt = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tickets.tickets.Ticket"))
		}
		panic(fmt.Errorf("message tickets.tickets.Ticket does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Ticket) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tickets.tickets.Ticket.name":
		panic(fmt.Errorf("field name of message tickets.tickets.Ticket is not mutable"))
	case "tickets.tickets.Ticket.event":
		panic(fmt.Errorf("field event of message tickets.tickets.Ticket is not mutable"))
	case "tickets.tickets.Ticket.id":
		panic(fmt.Errorf("field id of message tickets.tickets.Ticket is not mutable"))
	case "tickets.tickets.Ticket.createdAt":
		panic(fmt.Errorf("field createdAt of message tickets.tickets.Ticket is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tickets.tickets.Ticket"))
		}
		panic(fmt.Errorf("message tickets.tickets.Ticket does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Ticket) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tickets.tickets.Ticket.name":
		return protoreflect.ValueOfString("")
	case "tickets.tickets.Ticket.event":
		return protoreflect.ValueOfString("")
	case "tickets.tickets.Ticket.id":
		return protoreflect.ValueOfString("")
	case "tickets.tickets.Ticket.createdAt":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tickets.tickets.Ticket"))
		}
		panic(fmt.Errorf("message tickets.tickets.Ticket does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Ticket) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tickets.tickets.Ticket", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Ticket) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Ticket) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Ticket) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Ticket) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Ticket)
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
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Event)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Id)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.CreatedAt != 0 {
			n += 1 + runtime.Sov(uint64(x.CreatedAt))
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
		x := input.Message.Interface().(*Ticket)
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
		if x.CreatedAt != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.CreatedAt))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Id) > 0 {
			i -= len(x.Id)
			copy(dAtA[i:], x.Id)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Id)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Event) > 0 {
			i -= len(x.Event)
			copy(dAtA[i:], x.Event)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Event)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
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
		x := input.Message.Interface().(*Ticket)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Ticket: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Ticket: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
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
				x.Event = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
				x.Id = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
				}
				x.CreatedAt = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.CreatedAt |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: tickets/tickets/ticket.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Event     string `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Id        string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_tickets_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_tickets_tickets_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ticket) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_tickets_tickets_ticket_proto protoreflect.FileDescriptor

var file_tickets_tickets_ticket_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x22,
	0x60, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x42, 0x9c, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x0b, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0xa2, 0x02, 0x03, 0x54, 0x54, 0x58, 0xaa, 0x02, 0x0f, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0xca, 0x02, 0x0f,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0xe2,
	0x02, 0x1b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x3a, 0x3a, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tickets_tickets_ticket_proto_rawDescOnce sync.Once
	file_tickets_tickets_ticket_proto_rawDescData = file_tickets_tickets_ticket_proto_rawDesc
)

func file_tickets_tickets_ticket_proto_rawDescGZIP() []byte {
	file_tickets_tickets_ticket_proto_rawDescOnce.Do(func() {
		file_tickets_tickets_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_tickets_tickets_ticket_proto_rawDescData)
	})
	return file_tickets_tickets_ticket_proto_rawDescData
}

var file_tickets_tickets_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tickets_tickets_ticket_proto_goTypes = []interface{}{
	(*Ticket)(nil), // 0: tickets.tickets.Ticket
}
var file_tickets_tickets_ticket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tickets_tickets_ticket_proto_init() }
func file_tickets_tickets_ticket_proto_init() {
	if File_tickets_tickets_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tickets_tickets_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
			RawDescriptor: file_tickets_tickets_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tickets_tickets_ticket_proto_goTypes,
		DependencyIndexes: file_tickets_tickets_ticket_proto_depIdxs,
		MessageInfos:      file_tickets_tickets_ticket_proto_msgTypes,
	}.Build()
	File_tickets_tickets_ticket_proto = out.File
	file_tickets_tickets_ticket_proto_rawDesc = nil
	file_tickets_tickets_ticket_proto_goTypes = nil
	file_tickets_tickets_ticket_proto_depIdxs = nil
}
