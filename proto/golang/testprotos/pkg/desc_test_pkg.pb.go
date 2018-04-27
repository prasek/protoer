// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/desc_test_pkg.proto

/*
Package pkg is a generated protocol buffer package.

It is generated from these files:
	pkg/desc_test_pkg.proto

It has these top-level messages:
	Bar
*/
package pkg

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Foo int32

const (
	Foo_ABC Foo = 0
	Foo_DEF Foo = 1
	Foo_GHI Foo = 2
	Foo_JKL Foo = 3
	Foo_MNO Foo = 4
	Foo_PQR Foo = 5
	Foo_STU Foo = 6
	Foo_VWX Foo = 7
	Foo_Y_Z Foo = 8
)

var Foo_name = map[int32]string{
	0: "ABC",
	1: "DEF",
	2: "GHI",
	3: "JKL",
	4: "MNO",
	5: "PQR",
	6: "STU",
	7: "VWX",
	8: "Y_Z",
}
var Foo_value = map[string]int32{
	"ABC": 0,
	"DEF": 1,
	"GHI": 2,
	"JKL": 3,
	"MNO": 4,
	"PQR": 5,
	"STU": 6,
	"VWX": 7,
	"Y_Z": 8,
}

func (x Foo) Enum() *Foo {
	p := new(Foo)
	*p = x
	return p
}
func (x Foo) String() string {
	return proto.EnumName(Foo_name, int32(x))
}
func (x *Foo) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Foo_value, data, "Foo")
	if err != nil {
		return err
	}
	*x = Foo(value)
	return nil
}
func (Foo) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Bar struct {
	Baz              []Foo  `protobuf:"varint,1,rep,name=baz,enum=jhump.protoreflect.desc.Foo" json:"baz,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Bar) Reset()                    { *m = Bar{} }
func (m *Bar) String() string            { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()               {}
func (*Bar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Bar) GetBaz() []Foo {
	if m != nil {
		return m.Baz
	}
	return nil
}

func init() {
	proto.RegisterType((*Bar)(nil), "jhump.protoreflect.desc.Bar")
	proto.RegisterEnum("jhump.protoreflect.desc.Foo", Foo_name, Foo_value)
}

func init() { proto.RegisterFile("pkg/desc_test_pkg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xc8, 0x4e, 0xd7,
	0x4f, 0x49, 0x2d, 0x4e, 0x8e, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x2f, 0xc8, 0x4e, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcf, 0xca, 0x28, 0xcd, 0x2d, 0x80, 0x70, 0x8a, 0x52, 0xd3, 0x72,
	0x52, 0x93, 0x4b, 0xf4, 0x40, 0xea, 0x94, 0x4c, 0xb9, 0x98, 0x9d, 0x12, 0x8b, 0x84, 0xf4, 0xb8,
	0x98, 0x93, 0x12, 0xab, 0x24, 0x18, 0x15, 0x98, 0x35, 0xf8, 0x8c, 0x64, 0xf4, 0x70, 0xa8, 0xd6,
	0x73, 0xcb, 0xcf, 0x0f, 0x02, 0x29, 0xd4, 0x0a, 0xe3, 0x62, 0x76, 0xcb, 0xcf, 0x17, 0x62, 0xe7,
	0x62, 0x76, 0x74, 0x72, 0x16, 0x60, 0x00, 0x31, 0x5c, 0x5c, 0xdd, 0x04, 0x18, 0x41, 0x0c, 0x77,
	0x0f, 0x4f, 0x01, 0x26, 0x10, 0xc3, 0xcb, 0xdb, 0x47, 0x80, 0x19, 0xc4, 0xf0, 0xf5, 0xf3, 0x17,
	0x60, 0x01, 0x31, 0x02, 0x02, 0x83, 0x04, 0x58, 0x41, 0x8c, 0xe0, 0x90, 0x50, 0x01, 0x36, 0x10,
	0x23, 0x2c, 0x3c, 0x42, 0x80, 0x1d, 0xc4, 0x88, 0x8c, 0x8f, 0x12, 0xe0, 0x70, 0xb2, 0x8e, 0xb2,
	0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x28, 0x4a, 0x2c, 0x4e,
	0xcd, 0xd6, 0x07, 0xbb, 0x23, 0xb5, 0x08, 0x42, 0xeb, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb,
	0x83, 0x3c, 0x08, 0x16, 0x28, 0xd6, 0x2f, 0xc8, 0x4e, 0xb7, 0x2e, 0xc8, 0x4e, 0x07, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x82, 0xd4, 0x04, 0x95, 0xfe, 0x00, 0x00, 0x00,
}