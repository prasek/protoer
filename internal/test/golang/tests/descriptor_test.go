package tests

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/prasek/protoer/examples/desc"
	"github.com/prasek/protoer/internal/test/testutil"
)

func TestDefaultValues(t *testing.T) {
	fd, err := desc.LoadFileDescriptor("desc_test_defaults.proto")
	testutil.Ok(t, err)

	testCases := []struct {
		message, field string
		defaultVal     interface{}
	}{
		{"testprotos.PrimitiveDefaults", "fl32", float32(3.14159)},
		{"testprotos.PrimitiveDefaults", "fl64", 3.14159},
		{"testprotos.PrimitiveDefaults", "fl32d", float32(6.022140857e23)},
		{"testprotos.PrimitiveDefaults", "fl64d", 6.022140857e23},
		{"testprotos.PrimitiveDefaults", "fl32inf", float32(math.Inf(1))},
		{"testprotos.PrimitiveDefaults", "fl64inf", math.Inf(1)},
		{"testprotos.PrimitiveDefaults", "fl32negInf", float32(math.Inf(-1))},
		{"testprotos.PrimitiveDefaults", "fl64negInf", math.Inf(-1)},
		{"testprotos.PrimitiveDefaults", "fl32nan", float32(math.NaN())},
		{"testprotos.PrimitiveDefaults", "fl64nan", math.NaN()},
		{"testprotos.PrimitiveDefaults", "bl1", true},
		{"testprotos.PrimitiveDefaults", "bl2", false},
		{"testprotos.PrimitiveDefaults", "i32", int32(10101)},
		{"testprotos.PrimitiveDefaults", "i32n", int32(-10101)},
		{"testprotos.PrimitiveDefaults", "i32x", int32(0x20202)},
		{"testprotos.PrimitiveDefaults", "i32xn", int32(-0x20202)},
		{"testprotos.PrimitiveDefaults", "i64", int64(10101)},
		{"testprotos.PrimitiveDefaults", "i64n", int64(-10101)},
		{"testprotos.PrimitiveDefaults", "i64x", int64(0x20202)},
		{"testprotos.PrimitiveDefaults", "i64xn", int64(-0x20202)},
		{"testprotos.PrimitiveDefaults", "i32s", int32(10101)},
		{"testprotos.PrimitiveDefaults", "i32sn", int32(-10101)},
		{"testprotos.PrimitiveDefaults", "i32sx", int32(0x20202)},
		{"testprotos.PrimitiveDefaults", "i32sxn", int32(-0x20202)},
		{"testprotos.PrimitiveDefaults", "i64s", int64(10101)},
		{"testprotos.PrimitiveDefaults", "i64sn", int64(-10101)},
		{"testprotos.PrimitiveDefaults", "i64sx", int64(0x20202)},
		{"testprotos.PrimitiveDefaults", "i64sxn", int64(-0x20202)},
		{"testprotos.PrimitiveDefaults", "i32f", int32(10101)},
		{"testprotos.PrimitiveDefaults", "i32fn", int32(-10101)},
		{"testprotos.PrimitiveDefaults", "i32fx", int32(0x20202)},
		{"testprotos.PrimitiveDefaults", "i32fxn", int32(-0x20202)},
		{"testprotos.PrimitiveDefaults", "i64f", int64(10101)},
		{"testprotos.PrimitiveDefaults", "i64fn", int64(-10101)},
		{"testprotos.PrimitiveDefaults", "i64fx", int64(0x20202)},
		{"testprotos.PrimitiveDefaults", "i64fxn", int64(-0x20202)},
		{"testprotos.PrimitiveDefaults", "u32", uint32(10101)},
		{"testprotos.PrimitiveDefaults", "u32x", uint32(0x20202)},
		{"testprotos.PrimitiveDefaults", "u64", uint64(10101)},
		{"testprotos.PrimitiveDefaults", "u64x", uint64(0x20202)},
		{"testprotos.PrimitiveDefaults", "u32f", uint32(10101)},
		{"testprotos.PrimitiveDefaults", "u32fx", uint32(0x20202)},
		{"testprotos.PrimitiveDefaults", "u64f", uint64(10101)},
		{"testprotos.PrimitiveDefaults", "u64fx", uint64(0x20202)},

		{"testprotos.StringAndBytesDefaults", "dq", "this is a string with \"nested quotes\""},
		{"testprotos.StringAndBytesDefaults", "sq", "this is a string with \"nested quotes\""},
		{"testprotos.StringAndBytesDefaults", "escaped_bytes", []byte("\000\001\a\b\f\n\r\t\v\\'\"\xfe")},
		{"testprotos.StringAndBytesDefaults", "utf8_string", "\341\210\264"},
		{"testprotos.StringAndBytesDefaults", "string_with_zero", "hel\000lo"},
		{"testprotos.StringAndBytesDefaults", "bytes_with_zero", []byte("wor\000ld")},

		{"testprotos.EnumDefaults", "red", int32(0)},
		{"testprotos.EnumDefaults", "green", int32(1)},
		{"testprotos.EnumDefaults", "blue", int32(2)},
		{"testprotos.EnumDefaults", "zero", int32(0)},
		{"testprotos.EnumDefaults", "zed", int32(0)},
		{"testprotos.EnumDefaults", "one", int32(1)},
		{"testprotos.EnumDefaults", "dos", int32(2)},
	}
	for i, tc := range testCases {
		def := fd.FindMessage(tc.message).FindFieldByName(tc.field).GetDefaultValue()
		testutil.Eq(t, tc.defaultVal, def, "wrong default value for case %d: %s.%s", i, tc.message, tc.field)
	}
}

type descCase struct {
	name       string
	number     int32
	skipParent bool
	references map[string]childCases
}

type childCases struct {
	query interface{}
	cases []descCase
}

func refs(names ...string) []descCase {
	r := make([]descCase, len(names))
	for i, n := range names {
		r[i] = descCase{name: n, skipParent: true}
	}
	return r
}

func children(names ...string) []descCase {
	ch := make([]descCase, len(names))
	for i, n := range names {
		ch[i] = descCase{name: n}
	}
	return ch
}

type fld struct {
	name   string
	number int32
}

func fields(flds ...fld) []descCase {
	f := make([]descCase, len(flds))
	for i, field := range flds {
		f[i] = descCase{name: field.name, number: field.number, skipParent: true}
	}
	return f
}

func checkDescriptor(t *testing.T, caseName string, num int32, d desc.Descriptor, parent desc.Descriptor, fd *desc.FileDescriptor, c descCase) {
	// name and fully-qualified name
	testutil.Eq(t, c.name, d.GetFullyQualifiedName(), caseName)
	if _, ok := d.(*desc.FileDescriptor); ok {
		testutil.Eq(t, c.name, d.GetName(), caseName)
	} else {
		pos := strings.LastIndex(c.name, ".")
		n := c.name
		if pos >= 0 {
			n = c.name[pos+1:]
		}
		testutil.Eq(t, n, d.GetName(), caseName)
		// check that this object matches the canonical one returned by file descriptor
		testutil.Eq(t, d, d.GetFile().FindSymbol(d.GetFullyQualifiedName()), caseName)
	}

	// number
	switch d := d.(type) {
	case (*desc.FieldDescriptor):
		n := num + 1
		if c.number != 0 {
			n = c.number
		}
		testutil.Eq(t, n, d.GetNumber(), caseName)
	case (*desc.EnumValueDescriptor):
		n := num + 1
		if c.number != 0 {
			n = c.number
		}
		testutil.Eq(t, n, d.GetNumber(), caseName)
	default:
		if c.number != 0 {
			panic(fmt.Sprintf("%s: number should only be specified by fields and enum values! numnber = %d, desc = %v", caseName, c.number, d))
		}
	}

	// parent and file
	if !c.skipParent {
		testutil.Eq(t, parent, d.GetParent(), caseName)
		testutil.Eq(t, fd, d.GetFile(), caseName)
	}

	// comment
	/*
		if fd.GetName() == "desc_test1.proto" && d.GetName() != "desc_test1.proto" {
			expectedComment := "Comment for " + d.GetName()
			if msg, ok := d.(*MessageDescriptor); ok && msg.IsMapEntry() {
				// There are no comments on synthetic map-entry messages.
				expectedComment = ""
			} else if field, ok := d.(*FieldDescriptor); ok {
				if field.GetOwner().IsMapEntry() || field.GetType() == dpb.FieldDescriptorProto_TYPE_GROUP {
					// There are no comments for fields of synthetic map-entry messages either.
					// And comments for group fields end up on the synthetic message, not the field.
					expectedComment = ""
				}
			}
			testutil.Eq(t, expectedComment, strings.TrimSpace(d.GetSourceInfo().GetLeadingComments()), caseName)
		}
	*/

	// references
	for name, cases := range c.references {
		caseName := fmt.Sprintf("%s>%s", caseName, name)
		children := runQuery(d, cases.query)
		if testutil.Eq(t, len(cases.cases), len(children), caseName+" length") {
			for i, childCase := range cases.cases {
				caseName := fmt.Sprintf("%s[%d]", caseName, i)
				checkDescriptor(t, caseName, int32(i), children[i], d, fd, childCase)
			}
		}
	}
}

func runQuery(d desc.Descriptor, query interface{}) []desc.Descriptor {
	r := reflect.ValueOf(query).Call([]reflect.Value{reflect.ValueOf(d)})[0]
	if r.Kind() == reflect.Slice {
		ret := make([]desc.Descriptor, r.Len())
		for i := 0; i < r.Len(); i++ {
			ret[i] = r.Index(i).Interface().(desc.Descriptor)
		}
		return ret
	} else if r.IsNil() {
		return []desc.Descriptor{}
	} else {
		return []desc.Descriptor{r.Interface().(desc.Descriptor)}
	}
}
