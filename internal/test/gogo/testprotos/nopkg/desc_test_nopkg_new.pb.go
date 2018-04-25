// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nopkg/desc_test_nopkg_new.proto

/*
Package nopkg is a generated protocol buffer package.

It is generated from these files:
	nopkg/desc_test_nopkg_new.proto
	nopkg/desc_test_nopkg.proto

It has these top-level messages:
	TopLevel
*/
package nopkg

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

type TopLevel struct {
	I                            *int32   `protobuf:"varint,1,opt,name=i" json:"i,omitempty"`
	J                            *int64   `protobuf:"varint,2,opt,name=j" json:"j,omitempty"`
	K                            *int32   `protobuf:"zigzag32,3,opt,name=k" json:"k,omitempty"`
	L                            *int64   `protobuf:"zigzag64,4,opt,name=l" json:"l,omitempty"`
	M                            *uint32  `protobuf:"varint,5,opt,name=m" json:"m,omitempty"`
	N                            *uint64  `protobuf:"varint,6,opt,name=n" json:"n,omitempty"`
	O                            *uint32  `protobuf:"fixed32,7,opt,name=o" json:"o,omitempty"`
	P                            *uint64  `protobuf:"fixed64,8,opt,name=p" json:"p,omitempty"`
	Q                            *int32   `protobuf:"fixed32,9,opt,name=q" json:"q,omitempty"`
	R                            *int64   `protobuf:"fixed64,10,opt,name=r" json:"r,omitempty"`
	S                            *float32 `protobuf:"fixed32,11,opt,name=s" json:"s,omitempty"`
	T                            *float64 `protobuf:"fixed64,12,opt,name=t" json:"t,omitempty"`
	U                            []byte   `protobuf:"bytes,13,opt,name=u" json:"u,omitempty"`
	V                            *string  `protobuf:"bytes,14,opt,name=v" json:"v,omitempty"`
	W                            *bool    `protobuf:"varint,15,opt,name=w" json:"w,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *TopLevel) Reset()                    { *m = TopLevel{} }
func (m *TopLevel) String() string            { return proto.CompactTextString(m) }
func (*TopLevel) ProtoMessage()               {}
func (*TopLevel) Descriptor() ([]byte, []int) { return fileDescriptorDescTestNopkgNew, []int{0} }

var extRange_TopLevel = []proto.ExtensionRange{
	{Start: 100, End: 1000},
}

func (*TopLevel) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TopLevel
}

func (m *TopLevel) GetI() int32 {
	if m != nil && m.I != nil {
		return *m.I
	}
	return 0
}

func (m *TopLevel) GetJ() int64 {
	if m != nil && m.J != nil {
		return *m.J
	}
	return 0
}

func (m *TopLevel) GetK() int32 {
	if m != nil && m.K != nil {
		return *m.K
	}
	return 0
}

func (m *TopLevel) GetL() int64 {
	if m != nil && m.L != nil {
		return *m.L
	}
	return 0
}

func (m *TopLevel) GetM() uint32 {
	if m != nil && m.M != nil {
		return *m.M
	}
	return 0
}

func (m *TopLevel) GetN() uint64 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return 0
}

func (m *TopLevel) GetO() uint32 {
	if m != nil && m.O != nil {
		return *m.O
	}
	return 0
}

func (m *TopLevel) GetP() uint64 {
	if m != nil && m.P != nil {
		return *m.P
	}
	return 0
}

func (m *TopLevel) GetQ() int32 {
	if m != nil && m.Q != nil {
		return *m.Q
	}
	return 0
}

func (m *TopLevel) GetR() int64 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *TopLevel) GetS() float32 {
	if m != nil && m.S != nil {
		return *m.S
	}
	return 0
}

func (m *TopLevel) GetT() float64 {
	if m != nil && m.T != nil {
		return *m.T
	}
	return 0
}

func (m *TopLevel) GetU() []byte {
	if m != nil {
		return m.U
	}
	return nil
}

func (m *TopLevel) GetV() string {
	if m != nil && m.V != nil {
		return *m.V
	}
	return ""
}

func (m *TopLevel) GetW() bool {
	if m != nil && m.W != nil {
		return *m.W
	}
	return false
}

func init() {
	proto.RegisterType((*TopLevel)(nil), "TopLevel")
}

func init() { proto.RegisterFile("nopkg/desc_test_nopkg_new.proto", fileDescriptorDescTestNopkgNew) }

var fileDescriptorDescTestNopkgNew = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x1c, 0xd0, 0xbf, 0x6e, 0xea, 0x50,
	0x0c, 0xc7, 0x71, 0xfd, 0xf8, 0x1b, 0x72, 0xe1, 0x42, 0x99, 0xbc, 0xd5, 0xea, 0x64, 0x75, 0xe0,
	0x3c, 0x40, 0x37, 0xd4, 0xb1, 0x93, 0xd5, 0xa9, 0x0b, 0x6a, 0xe1, 0x34, 0x04, 0x92, 0x9c, 0x90,
	0x9c, 0x90, 0x07, 0xee, 0xd4, 0xb7, 0xa8, 0x9c, 0xc5, 0xd2, 0x47, 0x96, 0x07, 0x7f, 0xd3, 0xc7,
	0x2a, 0xd4, 0xd7, 0xcc, 0x9d, 0x7c, 0x7b, 0x3c, 0x44, 0xdf, 0xc6, 0xc3, 0xe0, 0x43, 0xe5, 0xfb,
	0x5d, 0xdd, 0x84, 0x18, 0x9e, 0x7e, 0x90, 0x26, 0xef, 0xa1, 0x7e, 0xf3, 0x77, 0x5f, 0x6c, 0x97,
	0x29, 0x72, 0x02, 0x43, 0xa6, 0x8a, 0xdc, 0x74, 0xa1, 0x11, 0x43, 0xc6, 0x8a, 0x8b, 0xe9, 0x4a,
	0x63, 0x86, 0x3c, 0x28, 0xae, 0xa6, 0x82, 0x26, 0x0c, 0xd9, 0x2a, 0x86, 0xbb, 0x92, 0xa6, 0x0c,
	0x59, 0x29, 0x4a, 0x53, 0x45, 0x33, 0x86, 0x4c, 0x14, 0x95, 0x29, 0xd0, 0x9c, 0x21, 0x73, 0x45,
	0x30, 0xd5, 0x94, 0x30, 0x64, 0xa6, 0xa8, 0x4d, 0x37, 0x5a, 0x30, 0x64, 0xad, 0xb8, 0x99, 0x1a,
	0x4a, 0x19, 0xb2, 0x51, 0x34, 0xa6, 0x96, 0xfe, 0x31, 0x64, 0xa4, 0x68, 0x4d, 0x91, 0x96, 0x0c,
	0x81, 0x22, 0x9a, 0x3a, 0x5a, 0x31, 0x64, 0xa9, 0xe8, 0x4c, 0x77, 0xfa, 0xcf, 0x90, 0x85, 0xe2,
	0x6e, 0xea, 0x69, 0xcd, 0x90, 0x44, 0xd1, 0x3f, 0x4f, 0x93, 0xd3, 0xe6, 0x77, 0xbe, 0x7f, 0xfd,
	0xd8, 0x67, 0x79, 0x3c, 0x77, 0x5f, 0xbb, 0x63, 0x28, 0xdd, 0xe5, 0xdc, 0x95, 0xb5, 0x1b, 0xfe,
	0x6f, 0xfc, 0x77, 0xe1, 0x8f, 0xd1, 0xe5, 0x55, 0xf4, 0x4d, 0xf5, 0x59, 0xb8, 0x2c, 0x64, 0xc1,
	0x59, 0xa9, 0x61, 0xdd, 0xba, 0xa1, 0xd7, 0xcb, 0x30, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x37,
	0xe0, 0xcf, 0xc7, 0x4d, 0x01, 0x00, 0x00,
}