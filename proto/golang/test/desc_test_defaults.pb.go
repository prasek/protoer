// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desc_test_defaults.proto

package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Color int32

const (
	Color_RED   Color = 0
	Color_GREEN Color = 1
	Color_BLUE  Color = 2
)

var Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}
func (x Color) String() string {
	return proto.EnumName(Color_name, int32(x))
}
func (x *Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Color_value, data, "Color")
	if err != nil {
		return err
	}
	*x = Color(value)
	return nil
}
func (Color) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Number int32

const (
	Number_ZERO Number = 0
	Number_ZED  Number = 0
	Number_NIL  Number = 0
	Number_NULL Number = 0
	Number_ONE  Number = 1
	Number_UNO  Number = 1
	Number_TWO  Number = 2
	Number_DOS  Number = 2
)

var Number_name = map[int32]string{
	0: "ZERO",
	// Duplicate value: 0: "ZED",
	// Duplicate value: 0: "NIL",
	// Duplicate value: 0: "NULL",
	1: "ONE",
	// Duplicate value: 1: "UNO",
	2: "TWO",
	// Duplicate value: 2: "DOS",
}
var Number_value = map[string]int32{
	"ZERO": 0,
	"ZED":  0,
	"NIL":  0,
	"NULL": 0,
	"ONE":  1,
	"UNO":  1,
	"TWO":  2,
	"DOS":  2,
}

func (x Number) Enum() *Number {
	p := new(Number)
	*p = x
	return p
}
func (x Number) String() string {
	return proto.EnumName(Number_name, int32(x))
}
func (x *Number) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Number_value, data, "Number")
	if err != nil {
		return err
	}
	*x = Number(value)
	return nil
}
func (Number) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type PrimitiveDefaults struct {
	// simple default
	Fl32 *float32 `protobuf:"fixed32,1,opt,name=fl32,def=3.14159" json:"fl32,omitempty"`
	Fl64 *float64 `protobuf:"fixed64,2,opt,name=fl64,def=3.14159" json:"fl64,omitempty"`
	// exponent notation
	Fl32D *float32 `protobuf:"fixed32,3,opt,name=fl32d,def=6.022141e+23" json:"fl32d,omitempty"`
	Fl64D *float64 `protobuf:"fixed64,4,opt,name=fl64d,def=6.022140857e+23" json:"fl64d,omitempty"`
	// special values: inf, -inf, and nan
	Fl32Inf    *float32 `protobuf:"fixed32,5,opt,name=fl32inf,def=inf" json:"fl32inf,omitempty"`
	Fl64Inf    *float64 `protobuf:"fixed64,6,opt,name=fl64inf,def=inf" json:"fl64inf,omitempty"`
	Fl32NegInf *float32 `protobuf:"fixed32,7,opt,name=fl32negInf,def=-inf" json:"fl32negInf,omitempty"`
	Fl64NegInf *float64 `protobuf:"fixed64,8,opt,name=fl64negInf,def=-inf" json:"fl64negInf,omitempty"`
	Fl32Nan    *float32 `protobuf:"fixed32,9,opt,name=fl32nan,def=nan" json:"fl32nan,omitempty"`
	Fl64Nan    *float64 `protobuf:"fixed64,10,opt,name=fl64nan,def=nan" json:"fl64nan,omitempty"`
	Bl1        *bool    `protobuf:"varint,11,opt,name=bl1,def=1" json:"bl1,omitempty"`
	Bl2        *bool    `protobuf:"varint,12,opt,name=bl2,def=0" json:"bl2,omitempty"`
	// signed
	I32    *int32 `protobuf:"varint,13,opt,name=i32,def=10101" json:"i32,omitempty"`
	I32N   *int32 `protobuf:"varint,14,opt,name=i32n,def=-10101" json:"i32n,omitempty"`
	I32X   *int32 `protobuf:"varint,15,opt,name=i32x,def=131586" json:"i32x,omitempty"`
	I32Xn  *int32 `protobuf:"varint,16,opt,name=i32xn,def=-131586" json:"i32xn,omitempty"`
	I64    *int64 `protobuf:"varint,17,opt,name=i64,def=10101" json:"i64,omitempty"`
	I64N   *int64 `protobuf:"varint,18,opt,name=i64n,def=-10101" json:"i64n,omitempty"`
	I64X   *int64 `protobuf:"varint,19,opt,name=i64x,def=131586" json:"i64x,omitempty"`
	I64Xn  *int64 `protobuf:"varint,20,opt,name=i64xn,def=-131586" json:"i64xn,omitempty"`
	I32S   *int32 `protobuf:"zigzag32,21,opt,name=i32s,def=10101" json:"i32s,omitempty"`
	I32Sn  *int32 `protobuf:"zigzag32,22,opt,name=i32sn,def=-10101" json:"i32sn,omitempty"`
	I32Sx  *int32 `protobuf:"zigzag32,23,opt,name=i32sx,def=131586" json:"i32sx,omitempty"`
	I32Sxn *int32 `protobuf:"zigzag32,24,opt,name=i32sxn,def=-131586" json:"i32sxn,omitempty"`
	I64S   *int64 `protobuf:"zigzag64,25,opt,name=i64s,def=10101" json:"i64s,omitempty"`
	I64Sn  *int64 `protobuf:"zigzag64,26,opt,name=i64sn,def=-10101" json:"i64sn,omitempty"`
	I64Sx  *int64 `protobuf:"zigzag64,27,opt,name=i64sx,def=131586" json:"i64sx,omitempty"`
	I64Sxn *int64 `protobuf:"zigzag64,28,opt,name=i64sxn,def=-131586" json:"i64sxn,omitempty"`
	I32F   *int32 `protobuf:"fixed32,29,opt,name=i32f,def=10101" json:"i32f,omitempty"`
	I32Fn  *int32 `protobuf:"fixed32,30,opt,name=i32fn,def=-10101" json:"i32fn,omitempty"`
	I32Fx  *int32 `protobuf:"fixed32,31,opt,name=i32fx,def=131586" json:"i32fx,omitempty"`
	I32Fxn *int32 `protobuf:"fixed32,32,opt,name=i32fxn,def=-131586" json:"i32fxn,omitempty"`
	I64F   *int64 `protobuf:"fixed64,33,opt,name=i64f,def=10101" json:"i64f,omitempty"`
	I64Fn  *int64 `protobuf:"fixed64,34,opt,name=i64fn,def=-10101" json:"i64fn,omitempty"`
	I64Fx  *int64 `protobuf:"fixed64,35,opt,name=i64fx,def=131586" json:"i64fx,omitempty"`
	I64Fxn *int64 `protobuf:"fixed64,36,opt,name=i64fxn,def=-131586" json:"i64fxn,omitempty"`
	// unsigned
	U32              *uint32 `protobuf:"varint,37,opt,name=u32,def=10101" json:"u32,omitempty"`
	U32X             *uint32 `protobuf:"varint,38,opt,name=u32x,def=131586" json:"u32x,omitempty"`
	U64              *uint64 `protobuf:"varint,39,opt,name=u64,def=10101" json:"u64,omitempty"`
	U64X             *uint64 `protobuf:"varint,40,opt,name=u64x,def=131586" json:"u64x,omitempty"`
	U32F             *uint32 `protobuf:"fixed32,41,opt,name=u32f,def=10101" json:"u32f,omitempty"`
	U32Fx            *uint32 `protobuf:"fixed32,42,opt,name=u32fx,def=131586" json:"u32fx,omitempty"`
	U64F             *uint64 `protobuf:"fixed64,43,opt,name=u64f,def=10101" json:"u64f,omitempty"`
	U64Fx            *uint64 `protobuf:"fixed64,44,opt,name=u64fx,def=131586" json:"u64fx,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PrimitiveDefaults) Reset()                    { *m = PrimitiveDefaults{} }
func (m *PrimitiveDefaults) String() string            { return proto.CompactTextString(m) }
func (*PrimitiveDefaults) ProtoMessage()               {}
func (*PrimitiveDefaults) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

const Default_PrimitiveDefaults_Fl32 float32 = 3.14159
const Default_PrimitiveDefaults_Fl64 float64 = 3.14159
const Default_PrimitiveDefaults_Fl32D float32 = 6.022141e+23
const Default_PrimitiveDefaults_Fl64D float64 = 6.022140857e+23

var Default_PrimitiveDefaults_Fl32Inf float32 = float32(math.Inf(1))
var Default_PrimitiveDefaults_Fl64Inf float64 = math.Inf(1)
var Default_PrimitiveDefaults_Fl32NegInf float32 = float32(math.Inf(-1))
var Default_PrimitiveDefaults_Fl64NegInf float64 = math.Inf(-1)
var Default_PrimitiveDefaults_Fl32Nan float32 = float32(math.NaN())
var Default_PrimitiveDefaults_Fl64Nan float64 = math.NaN()

const Default_PrimitiveDefaults_Bl1 bool = true
const Default_PrimitiveDefaults_Bl2 bool = false
const Default_PrimitiveDefaults_I32 int32 = 10101
const Default_PrimitiveDefaults_I32N int32 = -10101
const Default_PrimitiveDefaults_I32X int32 = 131586
const Default_PrimitiveDefaults_I32Xn int32 = -131586
const Default_PrimitiveDefaults_I64 int64 = 10101
const Default_PrimitiveDefaults_I64N int64 = -10101
const Default_PrimitiveDefaults_I64X int64 = 131586
const Default_PrimitiveDefaults_I64Xn int64 = -131586
const Default_PrimitiveDefaults_I32S int32 = 10101
const Default_PrimitiveDefaults_I32Sn int32 = -10101
const Default_PrimitiveDefaults_I32Sx int32 = 131586
const Default_PrimitiveDefaults_I32Sxn int32 = -131586
const Default_PrimitiveDefaults_I64S int64 = 10101
const Default_PrimitiveDefaults_I64Sn int64 = -10101
const Default_PrimitiveDefaults_I64Sx int64 = 131586
const Default_PrimitiveDefaults_I64Sxn int64 = -131586
const Default_PrimitiveDefaults_I32F int32 = 10101
const Default_PrimitiveDefaults_I32Fn int32 = -10101
const Default_PrimitiveDefaults_I32Fx int32 = 131586
const Default_PrimitiveDefaults_I32Fxn int32 = -131586
const Default_PrimitiveDefaults_I64F int64 = 10101
const Default_PrimitiveDefaults_I64Fn int64 = -10101
const Default_PrimitiveDefaults_I64Fx int64 = 131586
const Default_PrimitiveDefaults_I64Fxn int64 = -131586
const Default_PrimitiveDefaults_U32 uint32 = 10101
const Default_PrimitiveDefaults_U32X uint32 = 131586
const Default_PrimitiveDefaults_U64 uint64 = 10101
const Default_PrimitiveDefaults_U64X uint64 = 131586
const Default_PrimitiveDefaults_U32F uint32 = 10101
const Default_PrimitiveDefaults_U32Fx uint32 = 131586
const Default_PrimitiveDefaults_U64F uint64 = 10101
const Default_PrimitiveDefaults_U64Fx uint64 = 131586

func (m *PrimitiveDefaults) GetFl32() float32 {
	if m != nil && m.Fl32 != nil {
		return *m.Fl32
	}
	return Default_PrimitiveDefaults_Fl32
}

func (m *PrimitiveDefaults) GetFl64() float64 {
	if m != nil && m.Fl64 != nil {
		return *m.Fl64
	}
	return Default_PrimitiveDefaults_Fl64
}

func (m *PrimitiveDefaults) GetFl32D() float32 {
	if m != nil && m.Fl32D != nil {
		return *m.Fl32D
	}
	return Default_PrimitiveDefaults_Fl32D
}

func (m *PrimitiveDefaults) GetFl64D() float64 {
	if m != nil && m.Fl64D != nil {
		return *m.Fl64D
	}
	return Default_PrimitiveDefaults_Fl64D
}

func (m *PrimitiveDefaults) GetFl32Inf() float32 {
	if m != nil && m.Fl32Inf != nil {
		return *m.Fl32Inf
	}
	return Default_PrimitiveDefaults_Fl32Inf
}

func (m *PrimitiveDefaults) GetFl64Inf() float64 {
	if m != nil && m.Fl64Inf != nil {
		return *m.Fl64Inf
	}
	return Default_PrimitiveDefaults_Fl64Inf
}

func (m *PrimitiveDefaults) GetFl32NegInf() float32 {
	if m != nil && m.Fl32NegInf != nil {
		return *m.Fl32NegInf
	}
	return Default_PrimitiveDefaults_Fl32NegInf
}

func (m *PrimitiveDefaults) GetFl64NegInf() float64 {
	if m != nil && m.Fl64NegInf != nil {
		return *m.Fl64NegInf
	}
	return Default_PrimitiveDefaults_Fl64NegInf
}

func (m *PrimitiveDefaults) GetFl32Nan() float32 {
	if m != nil && m.Fl32Nan != nil {
		return *m.Fl32Nan
	}
	return Default_PrimitiveDefaults_Fl32Nan
}

func (m *PrimitiveDefaults) GetFl64Nan() float64 {
	if m != nil && m.Fl64Nan != nil {
		return *m.Fl64Nan
	}
	return Default_PrimitiveDefaults_Fl64Nan
}

func (m *PrimitiveDefaults) GetBl1() bool {
	if m != nil && m.Bl1 != nil {
		return *m.Bl1
	}
	return Default_PrimitiveDefaults_Bl1
}

func (m *PrimitiveDefaults) GetBl2() bool {
	if m != nil && m.Bl2 != nil {
		return *m.Bl2
	}
	return Default_PrimitiveDefaults_Bl2
}

func (m *PrimitiveDefaults) GetI32() int32 {
	if m != nil && m.I32 != nil {
		return *m.I32
	}
	return Default_PrimitiveDefaults_I32
}

func (m *PrimitiveDefaults) GetI32N() int32 {
	if m != nil && m.I32N != nil {
		return *m.I32N
	}
	return Default_PrimitiveDefaults_I32N
}

func (m *PrimitiveDefaults) GetI32X() int32 {
	if m != nil && m.I32X != nil {
		return *m.I32X
	}
	return Default_PrimitiveDefaults_I32X
}

func (m *PrimitiveDefaults) GetI32Xn() int32 {
	if m != nil && m.I32Xn != nil {
		return *m.I32Xn
	}
	return Default_PrimitiveDefaults_I32Xn
}

func (m *PrimitiveDefaults) GetI64() int64 {
	if m != nil && m.I64 != nil {
		return *m.I64
	}
	return Default_PrimitiveDefaults_I64
}

func (m *PrimitiveDefaults) GetI64N() int64 {
	if m != nil && m.I64N != nil {
		return *m.I64N
	}
	return Default_PrimitiveDefaults_I64N
}

func (m *PrimitiveDefaults) GetI64X() int64 {
	if m != nil && m.I64X != nil {
		return *m.I64X
	}
	return Default_PrimitiveDefaults_I64X
}

func (m *PrimitiveDefaults) GetI64Xn() int64 {
	if m != nil && m.I64Xn != nil {
		return *m.I64Xn
	}
	return Default_PrimitiveDefaults_I64Xn
}

func (m *PrimitiveDefaults) GetI32S() int32 {
	if m != nil && m.I32S != nil {
		return *m.I32S
	}
	return Default_PrimitiveDefaults_I32S
}

func (m *PrimitiveDefaults) GetI32Sn() int32 {
	if m != nil && m.I32Sn != nil {
		return *m.I32Sn
	}
	return Default_PrimitiveDefaults_I32Sn
}

func (m *PrimitiveDefaults) GetI32Sx() int32 {
	if m != nil && m.I32Sx != nil {
		return *m.I32Sx
	}
	return Default_PrimitiveDefaults_I32Sx
}

func (m *PrimitiveDefaults) GetI32Sxn() int32 {
	if m != nil && m.I32Sxn != nil {
		return *m.I32Sxn
	}
	return Default_PrimitiveDefaults_I32Sxn
}

func (m *PrimitiveDefaults) GetI64S() int64 {
	if m != nil && m.I64S != nil {
		return *m.I64S
	}
	return Default_PrimitiveDefaults_I64S
}

func (m *PrimitiveDefaults) GetI64Sn() int64 {
	if m != nil && m.I64Sn != nil {
		return *m.I64Sn
	}
	return Default_PrimitiveDefaults_I64Sn
}

func (m *PrimitiveDefaults) GetI64Sx() int64 {
	if m != nil && m.I64Sx != nil {
		return *m.I64Sx
	}
	return Default_PrimitiveDefaults_I64Sx
}

func (m *PrimitiveDefaults) GetI64Sxn() int64 {
	if m != nil && m.I64Sxn != nil {
		return *m.I64Sxn
	}
	return Default_PrimitiveDefaults_I64Sxn
}

func (m *PrimitiveDefaults) GetI32F() int32 {
	if m != nil && m.I32F != nil {
		return *m.I32F
	}
	return Default_PrimitiveDefaults_I32F
}

func (m *PrimitiveDefaults) GetI32Fn() int32 {
	if m != nil && m.I32Fn != nil {
		return *m.I32Fn
	}
	return Default_PrimitiveDefaults_I32Fn
}

func (m *PrimitiveDefaults) GetI32Fx() int32 {
	if m != nil && m.I32Fx != nil {
		return *m.I32Fx
	}
	return Default_PrimitiveDefaults_I32Fx
}

func (m *PrimitiveDefaults) GetI32Fxn() int32 {
	if m != nil && m.I32Fxn != nil {
		return *m.I32Fxn
	}
	return Default_PrimitiveDefaults_I32Fxn
}

func (m *PrimitiveDefaults) GetI64F() int64 {
	if m != nil && m.I64F != nil {
		return *m.I64F
	}
	return Default_PrimitiveDefaults_I64F
}

func (m *PrimitiveDefaults) GetI64Fn() int64 {
	if m != nil && m.I64Fn != nil {
		return *m.I64Fn
	}
	return Default_PrimitiveDefaults_I64Fn
}

func (m *PrimitiveDefaults) GetI64Fx() int64 {
	if m != nil && m.I64Fx != nil {
		return *m.I64Fx
	}
	return Default_PrimitiveDefaults_I64Fx
}

func (m *PrimitiveDefaults) GetI64Fxn() int64 {
	if m != nil && m.I64Fxn != nil {
		return *m.I64Fxn
	}
	return Default_PrimitiveDefaults_I64Fxn
}

func (m *PrimitiveDefaults) GetU32() uint32 {
	if m != nil && m.U32 != nil {
		return *m.U32
	}
	return Default_PrimitiveDefaults_U32
}

func (m *PrimitiveDefaults) GetU32X() uint32 {
	if m != nil && m.U32X != nil {
		return *m.U32X
	}
	return Default_PrimitiveDefaults_U32X
}

func (m *PrimitiveDefaults) GetU64() uint64 {
	if m != nil && m.U64 != nil {
		return *m.U64
	}
	return Default_PrimitiveDefaults_U64
}

func (m *PrimitiveDefaults) GetU64X() uint64 {
	if m != nil && m.U64X != nil {
		return *m.U64X
	}
	return Default_PrimitiveDefaults_U64X
}

func (m *PrimitiveDefaults) GetU32F() uint32 {
	if m != nil && m.U32F != nil {
		return *m.U32F
	}
	return Default_PrimitiveDefaults_U32F
}

func (m *PrimitiveDefaults) GetU32Fx() uint32 {
	if m != nil && m.U32Fx != nil {
		return *m.U32Fx
	}
	return Default_PrimitiveDefaults_U32Fx
}

func (m *PrimitiveDefaults) GetU64F() uint64 {
	if m != nil && m.U64F != nil {
		return *m.U64F
	}
	return Default_PrimitiveDefaults_U64F
}

func (m *PrimitiveDefaults) GetU64Fx() uint64 {
	if m != nil && m.U64Fx != nil {
		return *m.U64Fx
	}
	return Default_PrimitiveDefaults_U64Fx
}

type StringAndBytesDefaults struct {
	Dq               *string `protobuf:"bytes,1,opt,name=dq,def=this is a string with \"nested quotes\"" json:"dq,omitempty"`
	Sq               *string `protobuf:"bytes,2,opt,name=sq,def=this is a string with \"nested quotes\"" json:"sq,omitempty"`
	EscapedBytes     []byte  `protobuf:"bytes,3,opt,name=escaped_bytes,json=escapedBytes,def=\\000\\001\\007\\010\\014\\n\\r\\t\\013\\\\\\'\\\"\\376" json:"escaped_bytes,omitempty"`
	Utf8String       *string `protobuf:"bytes,4,opt,name=utf8_string,json=utf8String,def=ሴ" json:"utf8_string,omitempty"`
	StringWithZero   *string `protobuf:"bytes,5,opt,name=string_with_zero,json=stringWithZero,def=hel\x00lo" json:"string_with_zero,omitempty"`
	BytesWithZero    []byte  `protobuf:"bytes,6,opt,name=bytes_with_zero,json=bytesWithZero,def=wor\\000ld" json:"bytes_with_zero,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StringAndBytesDefaults) Reset()                    { *m = StringAndBytesDefaults{} }
func (m *StringAndBytesDefaults) String() string            { return proto.CompactTextString(m) }
func (*StringAndBytesDefaults) ProtoMessage()               {}
func (*StringAndBytesDefaults) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

const Default_StringAndBytesDefaults_Dq string = "this is a string with \"nested quotes\""
const Default_StringAndBytesDefaults_Sq string = "this is a string with \"nested quotes\""

var Default_StringAndBytesDefaults_EscapedBytes []byte = []byte("\x00\x01\a\b\f\n\r\t\v\\'\"\xfe")

const Default_StringAndBytesDefaults_Utf8String string = "ሴ"
const Default_StringAndBytesDefaults_StringWithZero string = "hel\x00lo"

var Default_StringAndBytesDefaults_BytesWithZero []byte = []byte("wor\x00ld")

func (m *StringAndBytesDefaults) GetDq() string {
	if m != nil && m.Dq != nil {
		return *m.Dq
	}
	return Default_StringAndBytesDefaults_Dq
}

func (m *StringAndBytesDefaults) GetSq() string {
	if m != nil && m.Sq != nil {
		return *m.Sq
	}
	return Default_StringAndBytesDefaults_Sq
}

func (m *StringAndBytesDefaults) GetEscapedBytes() []byte {
	if m != nil && m.EscapedBytes != nil {
		return m.EscapedBytes
	}
	return append([]byte(nil), Default_StringAndBytesDefaults_EscapedBytes...)
}

func (m *StringAndBytesDefaults) GetUtf8String() string {
	if m != nil && m.Utf8String != nil {
		return *m.Utf8String
	}
	return Default_StringAndBytesDefaults_Utf8String
}

func (m *StringAndBytesDefaults) GetStringWithZero() string {
	if m != nil && m.StringWithZero != nil {
		return *m.StringWithZero
	}
	return Default_StringAndBytesDefaults_StringWithZero
}

func (m *StringAndBytesDefaults) GetBytesWithZero() []byte {
	if m != nil && m.BytesWithZero != nil {
		return m.BytesWithZero
	}
	return append([]byte(nil), Default_StringAndBytesDefaults_BytesWithZero...)
}

type EnumDefaults struct {
	Red              *Color  `protobuf:"varint,1,opt,name=red,enum=test.Color,def=0" json:"red,omitempty"`
	Green            *Color  `protobuf:"varint,2,opt,name=green,enum=test.Color,def=1" json:"green,omitempty"`
	Blue             *Color  `protobuf:"varint,3,opt,name=blue,enum=test.Color,def=2" json:"blue,omitempty"`
	Zero             *Number `protobuf:"varint,4,opt,name=zero,enum=test.Number,def=0" json:"zero,omitempty"`
	Zed              *Number `protobuf:"varint,5,opt,name=zed,enum=test.Number,def=0" json:"zed,omitempty"`
	One              *Number `protobuf:"varint,6,opt,name=one,enum=test.Number,def=1" json:"one,omitempty"`
	Dos              *Number `protobuf:"varint,7,opt,name=dos,enum=test.Number,def=2" json:"dos,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EnumDefaults) Reset()                    { *m = EnumDefaults{} }
func (m *EnumDefaults) String() string            { return proto.CompactTextString(m) }
func (*EnumDefaults) ProtoMessage()               {}
func (*EnumDefaults) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

const Default_EnumDefaults_Red Color = Color_RED
const Default_EnumDefaults_Green Color = Color_GREEN
const Default_EnumDefaults_Blue Color = Color_BLUE
const Default_EnumDefaults_Zero Number = Number_ZERO
const Default_EnumDefaults_Zed Number = Number_ZED
const Default_EnumDefaults_One Number = Number_ONE
const Default_EnumDefaults_Dos Number = Number_DOS

func (m *EnumDefaults) GetRed() Color {
	if m != nil && m.Red != nil {
		return *m.Red
	}
	return Default_EnumDefaults_Red
}

func (m *EnumDefaults) GetGreen() Color {
	if m != nil && m.Green != nil {
		return *m.Green
	}
	return Default_EnumDefaults_Green
}

func (m *EnumDefaults) GetBlue() Color {
	if m != nil && m.Blue != nil {
		return *m.Blue
	}
	return Default_EnumDefaults_Blue
}

func (m *EnumDefaults) GetZero() Number {
	if m != nil && m.Zero != nil {
		return *m.Zero
	}
	return Default_EnumDefaults_Zero
}

func (m *EnumDefaults) GetZed() Number {
	if m != nil && m.Zed != nil {
		return *m.Zed
	}
	return Default_EnumDefaults_Zed
}

func (m *EnumDefaults) GetOne() Number {
	if m != nil && m.One != nil {
		return *m.One
	}
	return Default_EnumDefaults_One
}

func (m *EnumDefaults) GetDos() Number {
	if m != nil && m.Dos != nil {
		return *m.Dos
	}
	return Default_EnumDefaults_Dos
}

func init() {
	proto.RegisterType((*PrimitiveDefaults)(nil), "test.PrimitiveDefaults")
	proto.RegisterType((*StringAndBytesDefaults)(nil), "test.StringAndBytesDefaults")
	proto.RegisterType((*EnumDefaults)(nil), "test.EnumDefaults")
	proto.RegisterEnum("test.Color", Color_name, Color_value)
	proto.RegisterEnum("test.Number", Number_name, Number_value)
}

func init() { proto.RegisterFile("desc_test_defaults.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1088 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdf, 0x4e, 0xe3, 0x46,
	0x14, 0xc6, 0x71, 0xe2, 0x24, 0x64, 0x08, 0x60, 0xa6, 0x2d, 0x3b, 0xb0, 0x4b, 0xd7, 0xcd, 0xc2,
	0xae, 0x0b, 0x25, 0xd8, 0x89, 0x31, 0x34, 0x77, 0xa5, 0x44, 0xd5, 0x4a, 0x34, 0xa9, 0x4c, 0xe9,
	0x4a, 0xc8, 0x52, 0x14, 0xf0, 0x4c, 0xb0, 0x1a, 0xc6, 0xc4, 0x7f, 0xba, 0x29, 0x4f, 0xd0, 0x67,
	0xe8, 0x3b, 0xf4, 0xae, 0x0f, 0xd4, 0xbe, 0x49, 0x35, 0x67, 0xec, 0xc4, 0xb0, 0x37, 0xbd, 0xca,
	0xe8, 0xfb, 0x7d, 0x67, 0xce, 0xf9, 0x46, 0xa3, 0x71, 0x10, 0xf1, 0x69, 0x7c, 0x3b, 0x4c, 0x68,
	0x9c, 0x0c, 0x7d, 0xca, 0x46, 0xe9, 0x24, 0x89, 0x5b, 0x0f, 0x51, 0x98, 0x84, 0x58, 0x15, 0x62,
	0xf3, 0x2f, 0x84, 0x36, 0x7e, 0x8a, 0x82, 0xfb, 0x20, 0x09, 0x7e, 0xa3, 0xe7, 0x99, 0x03, 0xbf,
	0x44, 0x2a, 0x9b, 0x74, 0xda, 0x44, 0xd1, 0x15, 0xa3, 0xd4, 0xad, 0x75, 0x5a, 0x96, 0x6d, 0x1d,
	0x7f, 0xeb, 0x82, 0x28, 0xa1, 0x63, 0x93, 0x92, 0xae, 0x18, 0xca, 0x13, 0xe8, 0xd8, 0xb8, 0x89,
	0x2a, 0xc2, 0xe4, 0x93, 0x32, 0x94, 0x36, 0x9c, 0x96, 0xd9, 0x6e, 0x5b, 0xb6, 0x45, 0x0f, 0xda,
	0x1d, 0x57, 0x22, 0xbc, 0x27, 0x3c, 0x8e, 0xed, 0x13, 0x15, 0x76, 0x58, 0xcf, 0x3c, 0xe6, 0xe9,
	0xf1, 0x49, 0x6e, 0x73, 0x6c, 0x1f, 0xef, 0xa0, 0x9a, 0xf0, 0x07, 0x9c, 0x91, 0x0a, 0x6c, 0x56,
	0x0e, 0x38, 0x73, 0x73, 0x4d, 0x62, 0xc7, 0x16, 0xb8, 0x0a, 0xfb, 0xe4, 0x18, 0x34, 0xbc, 0x8b,
	0x90, 0x70, 0x72, 0x3a, 0x7e, 0xcf, 0x19, 0xa9, 0xc1, 0x06, 0xea, 0xa1, 0xb0, 0x14, 0x74, 0xe9,
	0x72, 0xec, 0xcc, 0xb5, 0x0c, 0xfb, 0xcc, 0x5d, 0xb9, 0x9e, 0x4f, 0xc2, 0x47, 0x9c, 0xd4, 0xe5,
	0x24, 0x7c, 0xc4, 0xdd, 0x5c, 0xcb, 0x27, 0x11, 0x18, 0xc9, 0x49, 0x32, 0x0c, 0x1a, 0xde, 0x44,
	0xe5, 0x9b, 0x89, 0x45, 0x56, 0x74, 0xc5, 0x58, 0xee, 0xaa, 0x49, 0x94, 0x52, 0x57, 0x08, 0xf8,
	0x85, 0xd0, 0xdb, 0xa4, 0x01, 0x7a, 0x85, 0x8d, 0x26, 0x31, 0x80, 0xb6, 0x00, 0x41, 0xa7, 0x4d,
	0x56, 0x75, 0xc5, 0xa8, 0x74, 0x2b, 0x96, 0x69, 0x99, 0x96, 0x2b, 0x14, 0xbc, 0x8d, 0xd4, 0xa0,
	0xd3, 0xe6, 0x64, 0x0d, 0x48, 0xf5, 0x50, 0x22, 0xd0, 0x32, 0x36, 0x23, 0xeb, 0x92, 0x59, 0x1d,
	0xeb, 0xf8, 0xd4, 0x01, 0x36, 0xc3, 0x3b, 0xa8, 0x22, 0x7e, 0x39, 0xd1, 0x00, 0xd6, 0x0e, 0x33,
	0x2a, 0x55, 0xe8, 0xe7, 0xd8, 0x64, 0x43, 0x57, 0x8c, 0xf2, 0xa2, 0x9f, 0x63, 0xc3, 0x9e, 0x8e,
	0xcd, 0x09, 0x06, 0xb2, 0xe8, 0xe7, 0xd8, 0x3c, 0x63, 0x33, 0xf2, 0x99, 0x64, 0xf3, 0x7e, 0x8e,
	0x2d, 0xfb, 0x39, 0xf6, 0x8c, 0x93, 0xcf, 0x01, 0x16, 0xfa, 0x09, 0x15, 0x6f, 0xc1, 0xa8, 0x31,
	0xf9, 0x42, 0x57, 0x8c, 0x8d, 0xbc, 0x21, 0x48, 0xf8, 0x15, 0x4c, 0x1a, 0x73, 0xb2, 0x09, 0x2c,
	0x6f, 0x29, 0xc5, 0x9c, 0xce, 0xc8, 0x0b, 0x49, 0x0b, 0x31, 0xe2, 0x19, 0x7e, 0x8d, 0xaa, 0xb0,
	0xe0, 0x84, 0x00, 0x9e, 0xb7, 0xcd, 0x64, 0xe8, 0xeb, 0xd8, 0x31, 0xd9, 0xd2, 0x15, 0x03, 0x2f,
	0xfa, 0x3a, 0xb6, 0xec, 0xeb, 0xd8, 0x31, 0x27, 0xdb, 0xc0, 0x16, 0x7d, 0x85, 0x98, 0xd3, 0x19,
	0x79, 0x29, 0x69, 0x21, 0x4e, 0xd6, 0x57, 0x2c, 0x38, 0x79, 0x05, 0xb8, 0xd0, 0x17, 0xe4, 0x2c,
	0x2f, 0x23, 0x3b, 0xba, 0x62, 0xac, 0x17, 0xf3, 0xb2, 0x2c, 0x11, 0xe3, 0xe4, 0x4b, 0x60, 0xc5,
	0xbc, 0x2c, 0xcf, 0xcb, 0x66, 0xe4, 0xb5, 0xa4, 0x85, 0xbc, 0x2c, 0xcf, 0xcb, 0x66, 0x9c, 0xe8,
	0x80, 0x9f, 0xe4, 0x65, 0xf3, 0xbc, 0x8c, 0x7c, 0xa5, 0x2b, 0x86, 0x56, 0xcc, 0xcb, 0xb2, 0x44,
	0x8c, 0x93, 0x26, 0xb0, 0x62, 0x5e, 0x96, 0xe7, 0x65, 0x33, 0xf2, 0x46, 0xd2, 0x42, 0x5e, 0x96,
	0xe7, 0x15, 0x7d, 0x77, 0x01, 0x3f, 0xc9, 0xcb, 0xe4, 0x7d, 0x4a, 0x3b, 0x6d, 0xb2, 0xa7, 0x2b,
	0xc6, 0xea, 0xfc, 0x3e, 0xa5, 0xf2, 0xfe, 0xa6, 0xe2, 0x8e, 0xbe, 0x05, 0x32, 0xbf, 0x33, 0x42,
	0x83, 0x22, 0xc7, 0x26, 0xef, 0x74, 0xc5, 0x50, 0x17, 0x45, 0xf2, 0x12, 0xa6, 0xe2, 0xa2, 0x19,
	0x40, 0x16, 0x45, 0xe2, 0xa2, 0x6d, 0xc1, 0x86, 0x8c, 0x7c, 0xad, 0x2b, 0x46, 0x6d, 0x9e, 0x30,
	0xcd, 0x4e, 0x36, 0x85, 0xb3, 0xdb, 0x07, 0x36, 0xcf, 0x00, 0x22, 0x14, 0x8a, 0xa3, 0x39, 0xd0,
	0x15, 0xa3, 0xba, 0x28, 0xcc, 0x8e, 0x26, 0x85, 0xf0, 0xdf, 0x00, 0x5b, 0x14, 0x0a, 0xb1, 0xf9,
	0x4f, 0x09, 0x6d, 0x5e, 0x26, 0x51, 0xc0, 0xc7, 0xdf, 0x71, 0xff, 0xec, 0xf7, 0x84, 0xc6, 0xf3,
	0x47, 0xf3, 0x18, 0x95, 0xfc, 0x29, 0x3c, 0x99, 0xf5, 0xee, 0x5e, 0x72, 0x17, 0xc4, 0x7a, 0x10,
	0xeb, 0x23, 0x3d, 0x06, 0xb7, 0xfe, 0x31, 0x48, 0xee, 0xf4, 0x26, 0xa7, 0x71, 0x42, 0x7d, 0x7d,
	0x9a, 0x86, 0x09, 0x8d, 0x9b, 0x6e, 0xc9, 0x9f, 0x8a, 0xb2, 0x78, 0x0a, 0x8f, 0xe9, 0xff, 0x2f,
	0x8b, 0xa7, 0xf8, 0x47, 0xb4, 0x4a, 0xe3, 0xdb, 0xd1, 0x03, 0xf5, 0x87, 0x37, 0x62, 0x0c, 0x78,
	0x70, 0x1b, 0x5d, 0xc3, 0x33, 0x4d, 0xd3, 0x33, 0x4d, 0xcb, 0x33, 0xcd, 0x13, 0xcf, 0xb4, 0x4c,
	0xcf, 0xb4, 0x6c, 0x8f, 0x7b, 0x91, 0x97, 0x78, 0xa6, 0xd5, 0xf1, 0x3c, 0xef, 0x9d, 0xd7, 0xf4,
	0x3a, 0x27, 0x8e, 0xdb, 0xc8, 0xca, 0x21, 0x04, 0xde, 0x45, 0x2b, 0x69, 0xc2, 0x4e, 0x87, 0xb2,
	0x2d, 0xbc, 0xcc, 0xf5, 0x6e, 0xf9, 0xdf, 0x3f, 0xfe, 0x76, 0x91, 0xd0, 0x65, 0x64, 0x6c, 0x22,
	0x4d, 0x1a, 0x86, 0x62, 0xae, 0xe1, 0x23, 0x8d, 0x42, 0x78, 0x9b, 0xeb, 0xdd, 0xea, 0x1d, 0x9d,
	0x2c, 0x4d, 0x42, 0x77, 0x4d, 0xf2, 0x0f, 0x41, 0x72, 0x77, 0x4d, 0xa3, 0x10, 0x5b, 0x68, 0x1d,
	0xc6, 0x2b, 0x14, 0x54, 0x61, 0xd0, 0xfa, 0xc7, 0x30, 0x12, 0xb3, 0x4e, 0x7c, 0x77, 0x15, 0x1c,
	0x79, 0x49, 0xf3, 0xcf, 0x12, 0x6a, 0xf4, 0x78, 0x7a, 0x3f, 0x3f, 0xd8, 0x26, 0x2a, 0x47, 0xd4,
	0x87, 0x93, 0x5d, 0x6b, 0xaf, 0xb4, 0xc4, 0x77, 0xab, 0xf5, 0x7d, 0x38, 0x09, 0xa3, 0x6e, 0xd9,
	0xed, 0x9d, 0xbb, 0x02, 0x62, 0x03, 0x55, 0xc6, 0x11, 0xa5, 0x1c, 0x0e, 0xf2, 0x99, 0xab, 0xf2,
	0x83, 0xdb, 0xeb, 0xf5, 0x5d, 0x69, 0xc0, 0x7b, 0x48, 0xbd, 0x99, 0xa4, 0x14, 0xce, 0xeb, 0x99,
	0x51, 0x3d, 0xbb, 0xb8, 0xea, 0xb9, 0x80, 0xf1, 0x5b, 0xa4, 0xc2, 0xb4, 0x2a, 0xd8, 0x1a, 0xd2,
	0xd6, 0x4f, 0xef, 0x6f, 0x68, 0xd4, 0x55, 0xaf, 0x7b, 0xee, 0xc0, 0x05, 0x8e, 0xdf, 0xa0, 0xf2,
	0x23, 0xf5, 0xe1, 0x14, 0x9e, 0xdb, 0xca, 0xd7, 0x62, 0xba, 0x47, 0xea, 0x0b, 0x53, 0xc8, 0x29,
	0x24, 0xff, 0xc4, 0x34, 0xe8, 0xf7, 0x5c, 0x41, 0x85, 0xc9, 0x0f, 0x63, 0xf8, 0x54, 0x7d, 0x62,
	0x3a, 0x1f, 0x5c, 0xba, 0x82, 0xee, 0xef, 0xa1, 0x0a, 0xcc, 0x8a, 0x6b, 0x48, 0x84, 0xd7, 0x96,
	0x70, 0x1d, 0xc9, 0x7c, 0x9a, 0x82, 0x97, 0x11, 0x24, 0xd0, 0x4a, 0xfb, 0xbf, 0xa0, 0xaa, 0x2c,
	0x15, 0x9a, 0x98, 0x56, 0x5b, 0x12, 0x15, 0xd7, 0x50, 0x51, 0x43, 0xe5, 0xfe, 0xfb, 0x0b, 0x6d,
	0x49, 0xb0, 0xfe, 0xd5, 0xc5, 0x85, 0x94, 0x06, 0xfd, 0x9e, 0xa6, 0x88, 0xc5, 0x55, 0x7f, 0x20,
	0x17, 0x3f, 0x7f, 0x18, 0x68, 0x25, 0xb1, 0x38, 0x1f, 0x5c, 0x6a, 0xa5, 0xed, 0x92, 0xa6, 0x9c,
	0x1d, 0x5e, 0x1f, 0x8c, 0x83, 0xe4, 0x2e, 0xbd, 0x69, 0xdd, 0x86, 0xf7, 0x47, 0x0f, 0xd1, 0x28,
	0xa6, 0xbf, 0x1e, 0xc1, 0xff, 0x09, 0x1a, 0xc9, 0xdf, 0xa3, 0x71, 0x38, 0x19, 0xf1, 0xf1, 0x91,
	0x18, 0xff, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xa2, 0x18, 0x21, 0x7e, 0x08, 0x00, 0x00,
}