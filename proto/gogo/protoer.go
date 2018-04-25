package gogo

import (
	"fmt"
	"reflect"
	"strings"

	gogo "github.com/gogo/protobuf/proto"
	dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

//var _ proto.UntypedProtoer = (*protoer)(nil)

var extensionDescType = (map[int32]*gogo.ExtensionDesc)(nil)

var defaultAliases = map[string]string{
	// gogo/protobuf has short names for these but golang/proto has the long names
	"google/protobuf/descriptor.proto":      "descriptor.proto",
	"google/protobuf/compiler/plugin.proto": "plugin.proto",

	// gogo/protobuf v1.0.0 has the short names
	// and slightly after v1.0.0 has the long names
	// starting with 43a6153f8c1fc6068a7797e53f8a640222248e0e
	"google/protobuf/any.proto":            "any.proto",
	"google/protobuf/api.proto":            "api.proto",
	"google/protobuf/duration.proto":       "duration.proto",
	"google/protobuf/empty.proto":          "empty.proto",
	"google/protobuf/field_mask.proto":     "field_mask.proto",
	"google/protobuf/source_context.proto": "source_context.proto",
	"google/protobuf/struct.proto":         "struct.proto",
	"google/protobuf/timestamp.proto":      "timestamp.proto",
	"google/protobuf/type.proto":           "type.proto",
	"google/protobuf/wrappers.proto":       "wrappers.proto",
}

type UntypedProtoer interface {
	Marshal(m interface{}) ([]byte, error)
	Unmarshal(b []byte, m interface{}) error

	Clone(m interface{}) interface{}
	Equal(m1, m2 interface{}) bool
	Merge(dst, src interface{})
	Reset(m interface{})
	Size(m interface{}) int

	HasExtension(m interface{}, field int32) bool
	ClearExtension(m interface{}, field int32)
	GetExtension(m interface{}, field int32) (interface{}, error)
	SetExtension(m interface{}, field int32, v interface{}) error
	RegisteredExtensions(m interface{}, desiredType interface{}) (interface{}, error)

	FileDescriptor(file string) []byte
	MessageName(m interface{}) string
	MessageType(name string) reflect.Type
}

func NewProtoer(aliases map[string]string) UntypedProtoer {
	g := &protoer{
		aliases: make(map[string]string),
	}

	if aliases == nil {
		aliases = defaultAliases
	}

	// bidirectional mapping
	for k, v := range aliases {
		g.aliases[k] = v
		g.aliases[v] = k
	}
	return g
}

type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

type protoer struct {
	aliases map[string]string
}

func (p *protoer) Marshal(m interface{}) ([]byte, error) {
	return gogo.Marshal(m.(Message))
}

func (p *protoer) Unmarshal(b []byte, m interface{}) error {
	return gogo.Unmarshal(b, m.(Message))
}

func (p *protoer) Clone(m interface{}) interface{} {
	return gogo.Clone(m.(Message)).(Message)
}

func (p *protoer) Equal(m1, m2 interface{}) bool {
	return gogo.Equal(m1.(Message), m2.(Message))
}

func (p *protoer) Merge(dst, src interface{}) {
	gogo.Merge(dst.(Message), src.(Message))
}

func (p *protoer) Reset(m interface{}) {
	m.(Message).Reset()
}

func (p *protoer) Size(m interface{}) int {
	return gogo.Size(m.(Message))
}

func (p *protoer) getExt(m interface{}, field int32) (Message, *gogo.ExtensionDesc, error) {
	if m == nil || reflect.ValueOf(m).IsNil() {
		return nil, nil, fmt.Errorf("no m message")
	}

	m, err := ToNativeDescriptor(m)
	if err != nil {
		return nil, nil, err
	}

	extensions := gogo.RegisteredExtensions(m.(Message))
	if m == nil {
		return nil, nil, fmt.Errorf("no registered proto extensions")
	}
	ext, ok := extensions[field]
	if !ok {
		return nil, nil, fmt.Errorf("Field %d not found.", field)
	}

	return m.(Message), ext, nil
}

func (p *protoer) HasExtension(m interface{}, field int32) bool {
	msg, ext, err := p.getExt(m, field)
	if err != nil {
		return false
	}
	return gogo.HasExtension(msg, ext)
}

func (p *protoer) ClearExtension(m interface{}, field int32) {
	msg, ext, _ := p.getExt(m, field)
	gogo.ClearExtension(msg, ext)
}

func (p *protoer) SetExtension(m interface{}, field int32, v interface{}) error {
	msg, ext, err := p.getExt(m, field)
	if err != nil {
		return err
	}
	return gogo.SetExtension(msg, ext, v)
}

func (p *protoer) GetExtension(m interface{}, field int32) (interface{}, error) {
	msg, ext, err := p.getExt(m, field)
	if err != nil {
		return nil, err
	}
	return gogo.GetExtension(msg, ext)
}

// RegisteredExtensions returns a map of the registered extensions of a
// protocol buffer struct, indexed by the extension number.
// The argument m should be a nil pointer to the struct type.
func (p *protoer) RegisteredExtensions(m interface{}, desiredType interface{}) (interface{}, error) {
	if m == nil || reflect.ValueOf(m).IsNil() {
		return nil, fmt.Errorf("m is nil")
	}

	if desiredType == nil {
		desiredType = extensionDescType
	}
	dt := reflect.TypeOf(desiredType)
	if dt.Kind() != reflect.Map && dt.Key().Kind() != reflect.Int32 && dt.Elem().Kind() != reflect.Ptr {
		panic(fmt.Sprintf("desiredType is not map[int32]*XXX, got %T", desiredType))
	}

	m, err := ToNativeDescriptor(m)
	if err != nil {
		return nil, err
	}

	extensions := gogo.RegisteredExtensions(m.(Message))

	// desired type
	if reflect.TypeOf(m) == dt {
		return m, nil
	}

	// convert type
	out := reflect.MakeMap(dt)
	dte := dt.Elem().Elem()

	setField := func(e reflect.Value, name string, val interface{}) {
		f := e.Elem().FieldByName(name)
		f.Set(reflect.ValueOf(val))
	}

	for k, v := range extensions {
		e := reflect.New(dte)

		setField(e, "ExtendedType", v.ExtendedType)
		setField(e, "ExtensionType", v.ExtensionType)
		setField(e, "Field", v.Field)
		setField(e, "Name", v.Name)
		setField(e, "Tag", v.Tag)
		setField(e, "Filename", v.Filename)

		out.SetMapIndex(reflect.ValueOf(k), e)
	}
	return out.Interface(), nil
}

func (p *protoer) FileDescriptor(file string) []byte {
	fdb := gogo.FileDescriptor(file)
	if fdb == nil {
		var ok bool
		alias, ok := p.aliases[file]
		if ok {
			if fdb = gogo.FileDescriptor(alias); fdb == nil {
				return nil
			}
		} else {
			return nil
		}
	}
	return fdb
}

func (p *protoer) MessageName(m interface{}) string {
	return gogo.MessageName(m.(Message))
}

func (p *protoer) MessageType(name string) reflect.Type {
	return gogo.MessageType(name)
}

func (p *protoer) Aliases() map[string]string {
	return p.aliases
}

type extendableProto interface {
	gogo.Message
	ExtensionRangeArray() []gogo.ExtensionRange
}

func ToNativeDescriptor(m interface{}) (gogo.Message, error) {
	var pbm gogo.Message
	var ok bool

	if pbm, ok = m.(extendableProto); ok {
		return pbm, nil
	}

	//if using different fork of protobuf, then convert it
	pt := reflect.TypeOf(m)
	typeName := strings.Split(pt.String(), ".")[1]
	b, err := gogo.Marshal(m.(Message))
	if err != nil {
		return nil, err
	}
	switch typeName {
	case "FileDescriptorSet":
		pbm = &dpb.FileDescriptorSet{}
	case "FileDescriptorProto":
		pbm = &dpb.FileDescriptorProto{}
	case "DescriptorProto":
		pbm = &dpb.DescriptorProto{}
	case "ExtensionRangeOptions":
		pbm = &dpb.ExtensionRangeOptions{}
	case "FieldDescriptorProto":
		pbm = &dpb.FieldDescriptorProto{}
	case "OneofDescriptorProto":
		pbm = &dpb.OneofDescriptorProto{}
	case "EnumDescriptorProto":
		pbm = &dpb.EnumDescriptorProto{}
	case "EnumValueDescriptorProto":
		pbm = &dpb.EnumValueDescriptorProto{}
	case "ServiceDescriptorProto":
		pbm = &dpb.ServiceDescriptorProto{}
	case "MethodDescriptorProto":
		pbm = &dpb.MethodDescriptorProto{}
	case "FileOptions":
		pbm = &dpb.FileOptions{}
	case "MessageOptions":
		pbm = &dpb.MessageOptions{}
	case "FieldOptions":
		pbm = &dpb.FieldOptions{}
	case "OneofOptions":
		pbm = &dpb.OneofOptions{}
	case "EnumOptions":
		pbm = &dpb.EnumOptions{}
	case "EnumValueOptions":
		pbm = &dpb.EnumValueOptions{}
	case "ServiceOptions":
		pbm = &dpb.ServiceOptions{}
	case "MethodOptions":
		pbm = &dpb.MethodOptions{}
	case "UninterpretedOption":
		pbm = &dpb.UninterpretedOption{}
	case "SourceCodeInfo":
		pbm = &dpb.SourceCodeInfo{}
	case "GeneratedCodeInfo":
		pbm = &dpb.GeneratedCodeInfo{}
	default:
		return nil, fmt.Errorf("not proto extendableProto")
	}
	err = gogo.Unmarshal(b, pbm)
	if err != nil {
		return nil, err
	}
	return pbm, nil
}
