package desc

import (
	"fmt"
	"strconv"

	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

var _ Descriptor = (*FieldDescriptor)(nil)

func (fd *FieldDescriptor) GetName() string {
	return fd.proto.GetName()
}

// GetNumber returns the tag number of this field.
func (fd *FieldDescriptor) GetNumber() int32 {
	return fd.proto.GetNumber()
}

func (fd *FieldDescriptor) GetFullyQualifiedName() string {
	return fd.fqn
}

func (fd *FieldDescriptor) GetParent() Descriptor {
	return fd.parent
}

func (fd *FieldDescriptor) GetFile() *FileDescriptor {
	return fd.file
}

func (fd *FieldDescriptor) asFieldDescriptorProto() *dpb.FieldDescriptorProto {
	return fd.proto
}

func (fd *FieldDescriptor) String() string {
	return fd.proto.String()
}

func (fd *FieldDescriptor) GetJSONName() string {
	if jsonName := fd.proto.GetJsonName(); jsonName != "" {
		return jsonName
	}
	return fd.proto.GetName()
}

func (fd *FieldDescriptor) GetFullyQualifiedJSONName() string {
	return fmt.Sprintf("%s.%s", fd.GetParent().GetFullyQualifiedName(), fd.GetJSONName())
}

// GetOwner returns the message type that this field belongs to. If this is a normal
// field then this is the same as GetParent. But for extensions, this will be the
// extendee message whereas GetParent refers to where the extension was declared.
func (fd *FieldDescriptor) GetOwner() *MessageDescriptor {
	return fd.owner
}

// IsExtension returns true if this is an extension field.
func (fd *FieldDescriptor) IsExtension() bool {
	return fd.proto.GetExtendee() != ""
}

// GetOneOf returns the one-of field set to which this field belongs. If this field
// is not part of a one-of then this method returns nil.
func (fd *FieldDescriptor) GetOneOf() *OneOfDescriptor {
	return fd.oneOf
}

// GetType returns the type of this field. If the type indicates an enum, the
// enum type can be queried via GetEnumType. If the type indicates a message, the
// message type can be queried via GetMessageType.
func (fd *FieldDescriptor) getType() dpb.FieldDescriptorProto_Type {
	return fd.proto.GetType()
}

// IsRequired returns true if this field has the "required" label.
func (fd *FieldDescriptor) IsRequired() bool {
	return fd.proto.GetLabel() == dpb.FieldDescriptorProto_LABEL_REQUIRED
}

// IsRepeated returns true if this field has the "repeated" label.
func (fd *FieldDescriptor) IsRepeated() bool {
	return fd.proto.GetLabel() == dpb.FieldDescriptorProto_LABEL_REPEATED
}

// IsMap returns true if this is a map field. If so, it will have the "repeated"
// label its type will be a message that represents a map entry. The map entry
// message will have exactly two fields: tag #1 is the key and tag #2 is the value.
func (fd *FieldDescriptor) IsMap() bool {
	return fd.isMap
}

// GetMapKeyType returns the type of the key field if this is a map field. If it is
// not a map field, nil is returned.
func (fd *FieldDescriptor) GetMapKeyType() *FieldDescriptor {
	if fd.isMap {
		return fd.msgType.FindFieldByNumber(int32(1))
	}
	return nil
}

// GetMapValueType returns the type of the value field if this is a map field. If it
// is not a map field, nil is returned.
func (fd *FieldDescriptor) GetMapValueType() *FieldDescriptor {
	if fd.isMap {
		return fd.msgType.FindFieldByNumber(int32(2))
	}
	return nil
}

// GetMessageType returns the type of this field if it is a message type. If
// this field is not a message type, it returns nil.
func (fd *FieldDescriptor) GetMessageType() *MessageDescriptor {
	return fd.msgType
}

// GetEnumType returns the type of this field if it is an enum type. If this
// field is not an enum type, it returns nil.
func (fd *FieldDescriptor) GetEnumType() *EnumDescriptor {
	return fd.enumType
}

// GetDefaultValue returns the default value for this field.
//
// If this field represents a message type, this method always returns nil (even though
// for proto2 files, the default value should be a default instance of the message type).
// If the field represents an enum type, this method returns an int32 corresponding to the
// enum value. If this field is a map, it returns a nil map[interface{}]interface{}. If
// this field is repeated (and not a map), it returns a nil []interface{}.
//
// Otherwise, it returns the declared default value for the field or a zero value, if no
// default is declared or if the file is proto3. The type of said return value corresponds
// to the type of the field:
//   |-------------------------|
//   | Declared Type | Go Type |
//   |---------------+---------|
//   | int32         |         |
//   | sint32        | int32   |
//   | sfixed32      |         |
//   |---------------+---------|
//   | uint32        | uint32  |
//   | fixed32       |         |
//   |---------------+---------|
//   | int64         |         |
//   | sint64        | int64   |
//   | sfixed64      |         |
//   |---------------+---------|
//   | uint64        | uint64  |
//   | fixed64       |         |
//   |---------------+---------|
//   | float         | float32 |
//   |---------------+---------|
//   | double        | float64 |
//   |---------------+---------|
//   | bool          | bool    |
//   |---------------+---------|
//   | bytes         | []byte  |
//   |---------------+---------|
//   | string        | string  |
//   |-------------------------|
func (fd *FieldDescriptor) GetDefaultValue() interface{} {
	return fd.determineDefault()
}

func (fd *FieldDescriptor) determineDefault() interface{} {
	if fd.IsMap() {
		return map[interface{}]interface{}(nil)
	} else if fd.IsRepeated() {
		return []interface{}(nil)
	} else if fd.msgType != nil {
		return nil
	}

	proto3 := fd.file.isProto3
	if !proto3 {
		def := fd.asFieldDescriptorProto().GetDefaultValue()
		if def != "" {
			ret := parseDefaultValue(fd, def)
			if ret != nil {
				return ret
			}
			// if we can't parse default value, fall-through to return normal default...
		}
	}

	switch fd.getType() {
	case dpb.FieldDescriptorProto_TYPE_FIXED32,
		dpb.FieldDescriptorProto_TYPE_UINT32:
		return uint32(0)
	case dpb.FieldDescriptorProto_TYPE_SFIXED32,
		dpb.FieldDescriptorProto_TYPE_INT32,
		dpb.FieldDescriptorProto_TYPE_SINT32:
		return int32(0)
	case dpb.FieldDescriptorProto_TYPE_FIXED64,
		dpb.FieldDescriptorProto_TYPE_UINT64:
		return uint64(0)
	case dpb.FieldDescriptorProto_TYPE_SFIXED64,
		dpb.FieldDescriptorProto_TYPE_INT64,
		dpb.FieldDescriptorProto_TYPE_SINT64:
		return int64(0)
	case dpb.FieldDescriptorProto_TYPE_FLOAT:
		return float32(0.0)
	case dpb.FieldDescriptorProto_TYPE_DOUBLE:
		return float64(0.0)
	case dpb.FieldDescriptorProto_TYPE_BOOL:
		return false
	case dpb.FieldDescriptorProto_TYPE_BYTES:
		return []byte(nil)
	case dpb.FieldDescriptorProto_TYPE_STRING:
		return ""
	case dpb.FieldDescriptorProto_TYPE_ENUM:
		if proto3 {
			return int32(0)
		}
		enumVals := fd.GetEnumType().GetValues()
		if len(enumVals) > 0 {
			return enumVals[0].GetNumber()
		} else {
			return int32(0) // WTF?
		}
	default:
		panic(fmt.Sprintf("Unknown field type: %v", fd.getType()))
	}
}

func parseDefaultValue(fd *FieldDescriptor, val string) interface{} {
	switch fd.getType() {
	case dpb.FieldDescriptorProto_TYPE_ENUM:
		vd := fd.GetEnumType().FindValueByName(val)
		if vd != nil {
			return vd.GetNumber()
		}
		return nil
	case dpb.FieldDescriptorProto_TYPE_BOOL:
		if val == "true" {
			return true
		} else if val == "false" {
			return false
		}
		return nil
	case dpb.FieldDescriptorProto_TYPE_BYTES:
		return []byte(unescape(val))
	case dpb.FieldDescriptorProto_TYPE_STRING:
		return val
	case dpb.FieldDescriptorProto_TYPE_FLOAT:
		if f, err := strconv.ParseFloat(val, 32); err == nil {
			return float32(f)
		} else {
			return float32(0)
		}
	case dpb.FieldDescriptorProto_TYPE_DOUBLE:
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			return f
		} else {
			return float64(0)
		}
	case dpb.FieldDescriptorProto_TYPE_INT32,
		dpb.FieldDescriptorProto_TYPE_SINT32,
		dpb.FieldDescriptorProto_TYPE_SFIXED32:
		if i, err := strconv.ParseInt(val, 10, 32); err == nil {
			return int32(i)
		} else {
			return int32(0)
		}
	case dpb.FieldDescriptorProto_TYPE_UINT32,
		dpb.FieldDescriptorProto_TYPE_FIXED32:
		if i, err := strconv.ParseUint(val, 10, 32); err == nil {
			return uint32(i)
		} else {
			return uint32(0)
		}
	case dpb.FieldDescriptorProto_TYPE_INT64,
		dpb.FieldDescriptorProto_TYPE_SINT64,
		dpb.FieldDescriptorProto_TYPE_SFIXED64:
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			return i
		} else {
			return int64(0)
		}
	case dpb.FieldDescriptorProto_TYPE_UINT64,
		dpb.FieldDescriptorProto_TYPE_FIXED64:
		if i, err := strconv.ParseUint(val, 10, 64); err == nil {
			return i
		} else {
			return uint64(0)
		}
	default:
		return nil
	}
}
