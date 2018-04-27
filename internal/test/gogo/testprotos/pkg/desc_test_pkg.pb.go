// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/desc_test_pkg.proto

/*
Package pkg is a generated protocol buffer package.

It is generated from these files:
	pkg/desc_test_pkg.proto

It has these top-level messages:
	Bar
*/
package pkg

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
func (Foo) EnumDescriptor() ([]byte, []int) { return fileDescriptorDescTestPkg, []int{0} }

type Bar struct {
	Baz              []Foo  `protobuf:"varint,1,rep,name=baz,enum=jhump.protoreflect.desc.Foo" json:"baz,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Bar) Reset()                    { *m = Bar{} }
func (m *Bar) String() string            { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()               {}
func (*Bar) Descriptor() ([]byte, []int) { return fileDescriptorDescTestPkg, []int{0} }

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

func init() { proto.RegisterFile("pkg/desc_test_pkg.proto", fileDescriptorDescTestPkg) }

var fileDescriptorDescTestPkg = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0x3b, 0x4f, 0xc4, 0x30,
	0x0c, 0xc0, 0x71, 0x8e, 0x00, 0x87, 0x32, 0x20, 0xab, 0xcb, 0x31, 0x30, 0x20, 0x26, 0xc4, 0x90,
	0x48, 0x48, 0x4c, 0x0c, 0xe8, 0x0a, 0x94, 0xf7, 0xeb, 0x80, 0x03, 0xba, 0x54, 0x69, 0x09, 0x69,
	0x49, 0x5b, 0x47, 0x49, 0xba, 0xf0, 0xe9, 0x91, 0xcb, 0xcc, 0xf6, 0xb3, 0xfc, 0x1f, 0x6c, 0x3e,
	0x73, 0xd6, 0xc8, 0x4f, 0x1d, 0xaa, 0x22, 0xea, 0x10, 0x0b, 0x67, 0x8d, 0x70, 0x1e, 0x23, 0x26,
	0xb3, 0xef, 0x7a, 0xe8, 0xdc, 0xdf, 0xe0, 0xf5, 0x57, 0xab, 0xab, 0x28, 0xa8, 0xdb, 0x3b, 0xe2,
	0x2c, 0x55, 0x3e, 0x11, 0x9c, 0x95, 0xea, 0x67, 0x7b, 0xb2, 0xcb, 0xf6, 0xb7, 0x0e, 0x77, 0xc4,
	0x3f, 0xb5, 0xc8, 0x10, 0x17, 0x14, 0x1e, 0x2c, 0x39, 0xcb, 0x10, 0x93, 0x29, 0x67, 0xf3, 0xf4,
	0x14, 0x56, 0x08, 0x67, 0xe7, 0x19, 0x4c, 0x08, 0x17, 0x97, 0x57, 0xb0, 0x4a, 0xb8, 0xbe, 0xb9,
	0x05, 0x46, 0xb8, 0xbb, 0x7f, 0x80, 0x35, 0xc2, 0xe3, 0xd3, 0x02, 0xd6, 0x09, 0xcf, 0x2f, 0xaf,
	0xb0, 0x41, 0x58, 0xbe, 0xbd, 0xc3, 0x94, 0xf0, 0x51, 0xe4, 0xb0, 0x99, 0xce, 0xf3, 0x13, 0xd3,
	0xc4, 0x7a, 0x28, 0x45, 0x85, 0x9d, 0x74, 0x5e, 0x05, 0x6d, 0xe5, 0x78, 0x87, 0xf6, 0xb2, 0xe9,
	0xa3, 0xf6, 0xbd, 0x6a, 0x25, 0x3d, 0x27, 0x0d, 0x1a, 0x1c, 0x35, 0xae, 0x83, 0x74, 0xd6, 0x1c,
	0x3b, 0x6b, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xac, 0x4b, 0x7a, 0xcc, 0x04, 0x01, 0x00, 0x00,
}
