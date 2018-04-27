// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desc_test1.proto

/*
Package testprotos is a generated protocol buffer package.

It is generated from these files:
	desc_test1.proto
	desc_test_defaults.proto
	desc_test_proto3.proto

It has these top-level messages:
	TestMessage
	AnotherTestMessage
	PrimitiveDefaults
	StringAndBytesDefaults
	EnumDefaults
	TestRequest
	TestResponse
	CustomOption
*/
package testprotos

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Comment for NestedEnum
type TestMessage_NestedEnum int32

const (
	// Comment for VALUE1
	TestMessage_VALUE1 TestMessage_NestedEnum = 1
	// Comment for VALUE2
	TestMessage_VALUE2 TestMessage_NestedEnum = 2
)

var TestMessage_NestedEnum_name = map[int32]string{
	1: "VALUE1",
	2: "VALUE2",
}
var TestMessage_NestedEnum_value = map[string]int32{
	"VALUE1": 1,
	"VALUE2": 2,
}

func (x TestMessage_NestedEnum) Enum() *TestMessage_NestedEnum {
	p := new(TestMessage_NestedEnum)
	*p = x
	return p
}
func (x TestMessage_NestedEnum) String() string {
	return proto.EnumName(TestMessage_NestedEnum_name, int32(x))
}
func (x *TestMessage_NestedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_NestedEnum_value, data, "TestMessage_NestedEnum")
	if err != nil {
		return err
	}
	*x = TestMessage_NestedEnum(value)
	return nil
}
func (TestMessage_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorDescTest1, []int{0, 0}
}

// Comment for DeeplyNestedEnum
type TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum int32

const (
	// Comment for VALUE1
	TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1 TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum = 1
	// Comment for VALUE2
	TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE2 TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum = 2
)

var TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name = map[int32]string{
	1: "VALUE1",
	2: "VALUE2",
}
var TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value = map[string]int32{
	"VALUE1": 1,
	"VALUE2": 2,
}

func (x TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) Enum() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	p := new(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum)
	*p = x
	return p
}
func (x TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) String() string {
	return proto.EnumName(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name, int32(x))
}
func (x *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value, data, "TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum")
	if err != nil {
		return err
	}
	*x = TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum(value)
	return nil
}
func (TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorDescTest1, []int{0, 0, 0, 0, 0}
}

// Comment for TestMessage
type TestMessage struct {
	// Comment for nm
	Nm *TestMessage_NestedMessage `protobuf:"bytes,1,opt,name=nm" json:"nm,omitempty"`
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,2,opt,name=anm" json:"anm,omitempty"`
	// Comment for yanm
	Yanm *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,3,opt,name=yanm" json:"yanm,omitempty"`
	// Comment for ne
	Ne               []TestMessage_NestedEnum `protobuf:"varint,4,rep,name=ne,enum=testprotos.TestMessage_NestedEnum" json:"ne,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptorDescTest1, []int{0} }

func (m *TestMessage) GetNm() *TestMessage_NestedMessage {
	if m != nil {
		return m.Nm
	}
	return nil
}

func (m *TestMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage) GetYanm() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

func (m *TestMessage) GetNe() []TestMessage_NestedEnum {
	if m != nil {
		return m.Ne
	}
	return nil
}

// Comment for NestedMessage
type TestMessage_NestedMessage struct {
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,1,opt,name=anm" json:"anm,omitempty"`
	// Comment for yanm
	Yanm             *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,2,opt,name=yanm" json:"yanm,omitempty"`
	XXX_unrecognized []byte                                                                  `json:"-"`
}

func (m *TestMessage_NestedMessage) Reset()         { *m = TestMessage_NestedMessage{} }
func (m *TestMessage_NestedMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage_NestedMessage) ProtoMessage()    {}
func (*TestMessage_NestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptorDescTest1, []int{0, 0}
}

func (m *TestMessage_NestedMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage_NestedMessage) GetYanm() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

// Comment for AnotherNestedMessage
type TestMessage_NestedMessage_AnotherNestedMessage struct {
	// Comment for yanm
	Yanm             []*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,1,rep,name=yanm" json:"yanm,omitempty"`
	XXX_unrecognized []byte                                                                    `json:"-"`
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage) Reset() {
	*m = TestMessage_NestedMessage_AnotherNestedMessage{}
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage) String() string {
	return proto.CompactTextString(m)
}
func (*TestMessage_NestedMessage_AnotherNestedMessage) ProtoMessage() {}
func (*TestMessage_NestedMessage_AnotherNestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptorDescTest1, []int{0, 0, 0}
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage) GetYanm() []*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

var E_TestMessage_NestedMessage_AnotherNestedMessage_Flags = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: ([]bool)(nil),
	Field:         200,
	Name:          "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.flags",
	Tag:           "varint,200,rep,packed,name=flags",
	Filename:      "desc_test1.proto",
}

// Comment for YetAnotherNestedMessage
type TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage struct {
	// Comment for foo
	Foo *string `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	// Comment for bar
	Bar *int32 `protobuf:"varint,2,opt,name=bar" json:"bar,omitempty"`
	// Comment for baz
	Baz []byte `protobuf:"bytes,3,opt,name=baz" json:"baz,omitempty"`
	// Comment for dne
	Dne *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum `protobuf:"varint,4,opt,name=dne,enum=testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum" json:"dne,omitempty"`
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,5,opt,name=anm" json:"anm,omitempty"`
	// Comment for nm
	Nm *TestMessage_NestedMessage `protobuf:"bytes,6,opt,name=nm" json:"nm,omitempty"`
	// Comment for tm
	Tm               *TestMessage `protobuf:"bytes,7,opt,name=tm" json:"tm,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) Reset() {
	*m = TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage{}
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) String() string {
	return proto.CompactTextString(m)
}
func (*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) ProtoMessage() {}
func (*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptorDescTest1, []int{0, 0, 0, 0}
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetFoo() string {
	if m != nil && m.Foo != nil {
		return *m.Foo
	}
	return ""
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetBar() int32 {
	if m != nil && m.Bar != nil {
		return *m.Bar
	}
	return 0
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetBaz() []byte {
	if m != nil {
		return m.Baz
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetDne() TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	if m != nil && m.Dne != nil {
		return *m.Dne
	}
	return TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetNm() *TestMessage_NestedMessage {
	if m != nil {
		return m.Nm
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetTm() *TestMessage {
	if m != nil {
		return m.Tm
	}
	return nil
}

// Comment for AnotherTestMessage
type AnotherTestMessage struct {
	// Comment for dne
	Dne *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum `protobuf:"varint,1,opt,name=dne,enum=testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum" json:"dne,omitempty"`
	// Comment for map_field1
	MapField1 map[int32]string `protobuf:"bytes,2,rep,name=map_field1,json=mapField1" json:"map_field1,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Comment for map_field2
	MapField2 map[int64]float32 `protobuf:"bytes,3,rep,name=map_field2,json=mapField2" json:"map_field2,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
	// Comment for map_field3
	MapField3 map[uint32]bool `protobuf:"bytes,4,rep,name=map_field3,json=mapField3" json:"map_field3,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Comment for map_field4
	MapField4 map[string]*AnotherTestMessage `protobuf:"bytes,5,rep,name=map_field4,json=mapField4" json:"map_field4,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Rocknroll *AnotherTestMessage_RockNRoll  `protobuf:"group,6,opt,name=RockNRoll,json=rocknroll" json:"rocknroll,omitempty"`
	// Comment for atmoo
	//
	// Types that are valid to be assigned to Atmoo:
	//	*AnotherTestMessage_Str
	//	*AnotherTestMessage_Int
	Atmoo                        isAnotherTestMessage_Atmoo `protobuf_oneof:"atmoo"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *AnotherTestMessage) Reset()                    { *m = AnotherTestMessage{} }
func (m *AnotherTestMessage) String() string            { return proto.CompactTextString(m) }
func (*AnotherTestMessage) ProtoMessage()               {}
func (*AnotherTestMessage) Descriptor() ([]byte, []int) { return fileDescriptorDescTest1, []int{1} }

var extRange_AnotherTestMessage = []proto.ExtensionRange{
	{Start: 100, End: 200},
}

func (*AnotherTestMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_AnotherTestMessage
}

type isAnotherTestMessage_Atmoo interface {
	isAnotherTestMessage_Atmoo()
}

type AnotherTestMessage_Str struct {
	Str string `protobuf:"bytes,7,opt,name=str,oneof"`
}
type AnotherTestMessage_Int struct {
	Int int64 `protobuf:"varint,8,opt,name=int,oneof"`
}

func (*AnotherTestMessage_Str) isAnotherTestMessage_Atmoo() {}
func (*AnotherTestMessage_Int) isAnotherTestMessage_Atmoo() {}

func (m *AnotherTestMessage) GetAtmoo() isAnotherTestMessage_Atmoo {
	if m != nil {
		return m.Atmoo
	}
	return nil
}

func (m *AnotherTestMessage) GetDne() TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	if m != nil && m.Dne != nil {
		return *m.Dne
	}
	return TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1
}

func (m *AnotherTestMessage) GetMapField1() map[int32]string {
	if m != nil {
		return m.MapField1
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField2() map[int64]float32 {
	if m != nil {
		return m.MapField2
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField3() map[uint32]bool {
	if m != nil {
		return m.MapField3
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField4() map[string]*AnotherTestMessage {
	if m != nil {
		return m.MapField4
	}
	return nil
}

func (m *AnotherTestMessage) GetRocknroll() *AnotherTestMessage_RockNRoll {
	if m != nil {
		return m.Rocknroll
	}
	return nil
}

func (m *AnotherTestMessage) GetStr() string {
	if x, ok := m.GetAtmoo().(*AnotherTestMessage_Str); ok {
		return x.Str
	}
	return ""
}

func (m *AnotherTestMessage) GetInt() int64 {
	if x, ok := m.GetAtmoo().(*AnotherTestMessage_Int); ok {
		return x.Int
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AnotherTestMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AnotherTestMessage_OneofMarshaler, _AnotherTestMessage_OneofUnmarshaler, _AnotherTestMessage_OneofSizer, []interface{}{
		(*AnotherTestMessage_Str)(nil),
		(*AnotherTestMessage_Int)(nil),
	}
}

func _AnotherTestMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AnotherTestMessage)
	// atmoo
	switch x := m.Atmoo.(type) {
	case *AnotherTestMessage_Str:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Str)
	case *AnotherTestMessage_Int:
		_ = b.EncodeVarint(8<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Int))
	case nil:
	default:
		return fmt.Errorf("AnotherTestMessage.Atmoo has unexpected type %T", x)
	}
	return nil
}

func _AnotherTestMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AnotherTestMessage)
	switch tag {
	case 7: // atmoo.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Atmoo = &AnotherTestMessage_Str{x}
		return true, err
	case 8: // atmoo.int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Atmoo = &AnotherTestMessage_Int{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _AnotherTestMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AnotherTestMessage)
	// atmoo
	switch x := m.Atmoo.(type) {
	case *AnotherTestMessage_Str:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case *AnotherTestMessage_Int:
		n += proto.SizeVarint(8<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Comment for RockNRoll
type AnotherTestMessage_RockNRoll struct {
	// Comment for beatles
	Beatles *string `protobuf:"bytes,1,opt,name=beatles" json:"beatles,omitempty"`
	// Comment for stones
	Stones *string `protobuf:"bytes,2,opt,name=stones" json:"stones,omitempty"`
	// Comment for doors
	Doors            *string `protobuf:"bytes,3,opt,name=doors" json:"doors,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AnotherTestMessage_RockNRoll) Reset()         { *m = AnotherTestMessage_RockNRoll{} }
func (m *AnotherTestMessage_RockNRoll) String() string { return proto.CompactTextString(m) }
func (*AnotherTestMessage_RockNRoll) ProtoMessage()    {}
func (*AnotherTestMessage_RockNRoll) Descriptor() ([]byte, []int) {
	return fileDescriptorDescTest1, []int{1, 4}
}

func (m *AnotherTestMessage_RockNRoll) GetBeatles() string {
	if m != nil && m.Beatles != nil {
		return *m.Beatles
	}
	return ""
}

func (m *AnotherTestMessage_RockNRoll) GetStones() string {
	if m != nil && m.Stones != nil {
		return *m.Stones
	}
	return ""
}

func (m *AnotherTestMessage_RockNRoll) GetDoors() string {
	if m != nil && m.Doors != nil {
		return *m.Doors
	}
	return ""
}

var E_Xtm = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*TestMessage)(nil),
	Field:         100,
	Name:          "testprotos.xtm",
	Tag:           "bytes,100,opt,name=xtm",
	Filename:      "desc_test1.proto",
}

var E_Xs = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*string)(nil),
	Field:         101,
	Name:          "testprotos.xs",
	Tag:           "bytes,101,opt,name=xs",
	Filename:      "desc_test1.proto",
}

var E_Xi = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         102,
	Name:          "testprotos.xi",
	Tag:           "varint,102,opt,name=xi",
	Filename:      "desc_test1.proto",
}

var E_Xui = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         103,
	Name:          "testprotos.xui",
	Tag:           "varint,103,opt,name=xui",
	Filename:      "desc_test1.proto",
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "testprotos.TestMessage")
	proto.RegisterType((*TestMessage_NestedMessage)(nil), "testprotos.TestMessage.NestedMessage")
	proto.RegisterType((*TestMessage_NestedMessage_AnotherNestedMessage)(nil), "testprotos.TestMessage.NestedMessage.AnotherNestedMessage")
	proto.RegisterType((*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage)(nil), "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage")
	proto.RegisterType((*AnotherTestMessage)(nil), "testprotos.AnotherTestMessage")
	proto.RegisterType((*AnotherTestMessage_RockNRoll)(nil), "testprotos.AnotherTestMessage.RockNRoll")
	proto.RegisterEnum("testprotos.TestMessage_NestedEnum", TestMessage_NestedEnum_name, TestMessage_NestedEnum_value)
	proto.RegisterEnum("testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum", TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name, TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value)
	proto.RegisterExtension(E_TestMessage_NestedMessage_AnotherNestedMessage_Flags)
	proto.RegisterExtension(E_Xtm)
	proto.RegisterExtension(E_Xs)
	proto.RegisterExtension(E_Xi)
	proto.RegisterExtension(E_Xui)
}

func init() { proto.RegisterFile("desc_test1.proto", fileDescriptorDescTest1) }

var fileDescriptorDescTest1 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xc7, 0xdd, 0x99, 0x6c, 0x92, 0x3d, 0xb5, 0x25, 0x0c, 0xc5, 0x2e, 0xbd, 0x90, 0x10, 0x14,
	0x43, 0xc1, 0x8d, 0xd9, 0x44, 0xd0, 0xe0, 0x4d, 0x8b, 0x2d, 0x0a, 0x6d, 0x2f, 0xc6, 0x0f, 0x50,
	0x84, 0xb2, 0xcd, 0x4e, 0xd2, 0x90, 0xdd, 0x9d, 0xb0, 0x33, 0x91, 0xa4, 0x4f, 0xe3, 0x85, 0x0f,
	0xe0, 0x63, 0xe8, 0x3b, 0xf8, 0x26, 0xde, 0xc8, 0xcc, 0x26, 0xcd, 0xe6, 0xa3, 0xdd, 0x56, 0x5a,
	0xaf, 0x32, 0x67, 0x72, 0xfe, 0xbf, 0x39, 0x5f, 0x7b, 0xa0, 0xe4, 0x33, 0xd1, 0x3e, 0x91, 0x4c,
	0xc8, 0xba, 0x33, 0x88, 0xb9, 0xe4, 0x04, 0x94, 0xa1, 0x8f, 0xa2, 0xf2, 0xa7, 0x08, 0x6b, 0xef,
	0x99, 0x90, 0x47, 0x4c, 0x08, 0xaf, 0xcb, 0xc8, 0x73, 0x40, 0x51, 0x68, 0x1b, 0x65, 0xa3, 0xba,
	0xe6, 0x3e, 0x76, 0x66, 0x8e, 0x4e, 0xca, 0xc9, 0x39, 0x66, 0x42, 0x32, 0x7f, 0x62, 0x51, 0x14,
	0x85, 0xe4, 0x10, 0xb0, 0x17, 0x85, 0x36, 0xd2, 0xba, 0xd6, 0xb5, 0x74, 0xce, 0x6e, 0xc4, 0xe5,
	0x19, 0x8b, 0xe7, 0x61, 0x0a, 0x43, 0x3a, 0x90, 0x1b, 0x2b, 0x1c, 0xd6, 0x38, 0xfa, 0xef, 0x38,
	0xe7, 0x13, 0x93, 0x2b, 0x9f, 0xd1, 0x7c, 0xe2, 0x02, 0x8a, 0x98, 0x9d, 0x2b, 0xe3, 0xea, 0x86,
	0x5b, 0xb9, 0xfa, 0x95, 0xfd, 0x68, 0x18, 0x52, 0x14, 0xb1, 0xed, 0xef, 0x79, 0x58, 0x9f, 0x63,
	0x4d, 0x73, 0x37, 0x6e, 0x37, 0x77, 0x74, 0xb7, 0xb9, 0x6f, 0xff, 0xce, 0xc1, 0xe6, 0xaa, 0xbf,
	0x2f, 0x02, 0x30, 0xca, 0xf8, 0x4e, 0x03, 0xf8, 0x86, 0x61, 0xeb, 0x12, 0x0f, 0x52, 0x02, 0xdc,
	0xe1, 0x5c, 0x97, 0xd4, 0xa2, 0xea, 0xa8, 0x6e, 0x4e, 0xbd, 0x58, 0x57, 0xc5, 0xa4, 0xea, 0x98,
	0xdc, 0x9c, 0xeb, 0x19, 0xb9, 0xaf, 0x6e, 0xce, 0xc9, 0x10, 0xb0, 0xaf, 0xfb, 0x69, 0x54, 0x37,
	0xdc, 0xf6, 0xed, 0x07, 0xee, 0xbc, 0x66, 0x6c, 0x10, 0x8c, 0x53, 0x03, 0xa1, 0xde, 0x9b, 0xf6,
	0xdf, 0xbc, 0x9d, 0xfe, 0x27, 0x1f, 0x60, 0xfe, 0xa6, 0x1f, 0xe0, 0x13, 0x40, 0x32, 0xb4, 0x0b,
	0x5a, 0xb6, 0x75, 0x89, 0x8c, 0x22, 0x19, 0x56, 0x76, 0xa0, 0xb4, 0x98, 0x06, 0x01, 0xc8, 0x7f,
	0xdc, 0x3d, 0xfc, 0xb0, 0x5f, 0x2f, 0x19, 0x17, 0x67, 0xb7, 0x84, 0xdc, 0x97, 0x60, 0x76, 0x02,
	0xaf, 0x2b, 0xc8, 0xc3, 0x34, 0x71, 0x12, 0x7b, 0x0a, 0x6c, 0xff, 0x54, 0xc3, 0x52, 0xdc, 0x43,
	0x25, 0x83, 0x26, 0x8a, 0xca, 0x23, 0x80, 0xec, 0x07, 0x2a, 0x3f, 0x0a, 0x40, 0x96, 0x71, 0xd3,
	0x46, 0x1a, 0xff, 0xbd, 0x91, 0x10, 0x7a, 0x83, 0x93, 0x4e, 0x8f, 0x05, 0x7e, 0xdd, 0x46, 0x7a,
	0xfe, 0x9f, 0x5e, 0x9d, 0xb9, 0x73, 0xe4, 0x0d, 0x0e, 0xb4, 0xff, 0x7e, 0x24, 0xe3, 0x31, 0xb5,
	0xc2, 0xa9, 0x3d, 0x47, 0x73, 0x6d, 0x7c, 0x23, 0x9a, 0xbb, 0x40, 0x73, 0xe7, 0x68, 0x0d, 0xbd,
	0xb2, 0xae, 0x4f, 0x6b, 0x2c, 0xd0, 0x1a, 0x73, 0xb4, 0xa6, 0x6d, 0xde, 0x88, 0xd6, 0x5c, 0xa0,
	0x35, 0xc9, 0x01, 0x58, 0x31, 0x6f, 0xf7, 0xa3, 0x98, 0x07, 0x81, 0x9e, 0x5c, 0x70, 0xab, 0x19,
	0x30, 0xca, 0xdb, 0xfd, 0x63, 0xca, 0x83, 0x80, 0xce, 0xa4, 0x84, 0x00, 0x16, 0x32, 0xd6, 0x43,
	0x6c, 0xbd, 0xb9, 0x47, 0x95, 0xa1, 0xee, 0x7a, 0x91, 0xb4, 0x8b, 0x65, 0xa3, 0x8a, 0xd5, 0x5d,
	0x2f, 0x92, 0xdb, 0xaf, 0x60, 0x63, 0xbe, 0xec, 0x6a, 0x17, 0xf4, 0xd9, 0x58, 0x0f, 0x8c, 0x49,
	0xd5, 0x91, 0x6c, 0x82, 0xf9, 0xd5, 0x0b, 0x86, 0x4c, 0x6f, 0x0c, 0x8b, 0x26, 0x46, 0x0b, 0xbd,
	0x30, 0xd2, 0x6a, 0x77, 0x49, 0x8d, 0x57, 0xa8, 0xd1, 0x25, 0xea, 0xc6, 0x92, 0x7a, 0x7d, 0x85,
	0xba, 0x98, 0x56, 0x7f, 0x99, 0xa9, 0x9b, 0x4b, 0x6a, 0x2b, 0x51, 0x37, 0xd3, 0xea, 0x35, 0x37,
	0xe3, 0xd3, 0x4b, 0xd3, 0xdf, 0x81, 0x75, 0x51, 0x57, 0x62, 0x43, 0xe1, 0x94, 0x79, 0x32, 0x60,
	0x62, 0x02, 0x9f, 0x9a, 0xe4, 0x01, 0xe4, 0x85, 0xe4, 0x11, 0x13, 0x93, 0xda, 0x4c, 0x2c, 0x15,
	0xb6, 0xcf, 0x79, 0x2c, 0xf4, 0x4a, 0xb5, 0x68, 0x62, 0xec, 0x98, 0x45, 0xbf, 0xf4, 0xcb, 0xd8,
	0x2b, 0x80, 0xe9, 0xc9, 0x90, 0xf3, 0xd6, 0x5b, 0xc0, 0x23, 0x19, 0x66, 0x6e, 0x04, 0xff, 0xea,
	0x4d, 0xa4, 0x18, 0x2d, 0x07, 0xd0, 0x28, 0x7b, 0xb7, 0x30, 0x1d, 0x0d, 0x1a, 0x09, 0xed, 0xdf,
	0xcb, 0xf4, 0xef, 0xe8, 0x21, 0x40, 0xa3, 0x5e, 0xeb, 0x19, 0xe0, 0xd1, 0x30, 0x5b, 0xd0, 0x2d,
	0x1b, 0xd5, 0x1c, 0x55, 0xae, 0x7b, 0xf5, 0xcf, 0xb5, 0x6e, 0x4f, 0x9e, 0x0d, 0x4f, 0x9d, 0x36,
	0x0f, 0x6b, 0x83, 0xd8, 0x13, 0xac, 0x5f, 0xd3, 0x4a, 0x16, 0x27, 0xbf, 0xb5, 0x2e, 0xef, 0xf2,
	0xda, 0x0c, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xd6, 0xf0, 0xe9, 0x5f, 0x09, 0x00, 0x00,
}