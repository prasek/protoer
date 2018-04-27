package golang

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	otherproto "github.com/gogo/protobuf/proto"
	dpbother "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	thisproto "github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/prasek/protoer/proto"
	otherprotos "github.com/prasek/protoer/proto/gogo/test"
	"github.com/prasek/protoer/proto/golang/test"
	"github.com/stretchr/testify/require"

	_ "github.com/golang/protobuf/protoc-gen-go/plugin"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/protobuf/api"
	_ "google.golang.org/genproto/protobuf/field_mask"
)

func TestMain(m *testing.M) {
	proto.SetProtoer(NewProtoer(nil))

	code := m.Run()
	os.Exit(code)
}

type lookup struct {
	// key = "package.TypeName"
	message map[string]*dpb.DescriptorProto
	// key = "package.MessageType.FieldType
	field map[string]*dpb.FieldDescriptorProto
	// key = "package.EnumName"
	enum map[string]*dpb.EnumDescriptorProto
	// key = "package.ServiceName"
	service map[string]*dpb.ServiceDescriptorProto
	// key = "/package.ServiceName/MethodName"
	method map[string]*dpb.MethodDescriptorProto
}

func TestLoadFileDescriptorForWellKnownProtos(t *testing.T) {
	wellKnownProtos := map[string][]string{
		"google/protobuf/any.proto":             {"google.protobuf.Any"},
		"google/protobuf/api.proto":             {"google.protobuf.Api", "google.protobuf.Method", "google.protobuf.Mixin"},
		"google/protobuf/descriptor.proto":      {"google.protobuf.FileDescriptorSet", "google.protobuf.DescriptorProto"},
		"google/protobuf/duration.proto":        {"google.protobuf.Duration"},
		"google/protobuf/empty.proto":           {"google.protobuf.Empty"},
		"google/protobuf/field_mask.proto":      {"google.protobuf.FieldMask"},
		"google/protobuf/source_context.proto":  {"google.protobuf.SourceContext"},
		"google/protobuf/struct.proto":          {"google.protobuf.Struct", "google.protobuf.Value", "google.protobuf.NullValue"},
		"google/protobuf/timestamp.proto":       {"google.protobuf.Timestamp"},
		"google/protobuf/type.proto":            {"google.protobuf.Type", "google.protobuf.Field", "google.protobuf.Syntax"},
		"google/protobuf/wrappers.proto":        {"google.protobuf.DoubleValue", "google.protobuf.Int32Value", "google.protobuf.StringValue"},
		"google/protobuf/compiler/plugin.proto": {"google.protobuf.compiler.CodeGeneratorRequest"},
	}

	aliases := proto.Aliases()

	//for file, types := range wellKnownProtos {
	for file := range wellKnownProtos {
		fd, err := loadFileDescriptorProto(file)
		require.Nil(t, err)
		require.Equal(t, file, fd.GetName())

		// also try loading via alternate name
		if aliases == nil {
			continue
		}
		file = aliases[file]
		if file == "" {
			// not a file that has a known alternate, so nothing else to check...
			continue
		}
		fd, err = loadFileDescriptorProto(file)
		require.Nil(t, err)
		require.Equal(t, file, fd.GetName())
	}
}

func TestProto3(t *testing.T) {

	file := "desc_test_proto3.proto"

	fd, err := loadFileDescriptorProto(file)
	require.Nil(t, err)
	require.Equal(t, 1, len(fd.Service), "service descriptor len")
	require.Equal(t, 3, len(fd.MessageType), "message descriptor len")
	require.Equal(t, 1, len(fd.EnumType), "enum descriptor len")

	l := makeLookup(fd)
	require.Equal(t, 4, len(l.method), "l.method descriptor len")
	require.Equal(t, 9, len(l.field), "l.field descriptor len")
	require.Equal(t, len(fd.Service), len(l.service), "l.service descriptor len")
	require.Equal(t, len(fd.MessageType), len(l.message), "l.message descriptor len")
	require.Equal(t, len(fd.EnumType), len(l.enum), "l.enum descriptor len")

	sd := l.service["test.TestService"]
	require.NotNil(t, sd)

	var sd2 = &dpb.ServiceDescriptorProto{}

	// marshal
	b, err := proto.Marshal(sd)
	require.Nil(t, err)

	// unmarshal
	err = proto.Unmarshal(b, sd2)
	require.Nil(t, err)
	ok := proto.Equal(sd, sd2)
	require.Equal(t, ok, true, "proto.Equal: sd != sd2")
	ok = reflect.DeepEqual(sd, sd2)
	require.Equal(t, ok, true, "reflect.DeepEqual: sd != sd2")

	// clone
	v := proto.Clone(sd2)
	sd3, ok := v.(*dpb.ServiceDescriptorProto)
	require.Equal(t, ok, true, "not *dpb.ServiceDescriptorProto")
	require.Equal(t, sd2 == sd3, false, "sd2 and sd3 pointer values are the same, should be different")

	// equal
	ok = proto.Equal(sd, sd3)
	require.Equal(t, ok, true, "proto.Equal: sd != sd3")
	ok = reflect.DeepEqual(sd, sd3)
	require.Equal(t, ok, true, "reflect.DeepEqual: sd != sd3")

	// reset
	proto.Reset(sd3)
	sd4 := &dpb.ServiceDescriptorProto{}
	ok = proto.Equal(sd3, sd4)
	require.Equal(t, ok, true, "proto.Equal: sd3 != sd4")
	ok = reflect.DeepEqual(sd3, sd4)
	require.Equal(t, ok, true, "reflect.DeepEqual: sd3 != sd4")

	//merge
	co1 := &test.CustomOption{Name: "foo"}
	co2 := &test.CustomOption{Value: 123}
	proto.Merge(co1, co2)
	require.Equal(t, "foo", co1.GetName(), "Merge name: co1")
	require.Equal(t, int32(123), co1.GetValue(), "Merge value: co1")
	require.Equal(t, "", co2.GetName(), "Merge name: co2")
	require.Equal(t, int32(123), co2.GetValue(), "Merge value: co2")

	size1 := proto.Size(co1)
	size2 := proto.Size(co2)
	require.Equal(t, 7, size1, "proto.Size1: not equal")
	require.Equal(t, 2, size2, "proto.Size2: not equal")

	// extensions
	method := l.method["/test.TestService/DoSomething"]
	require.NotNil(t, method, "Method DoSomething missing")

	// hasextension
	ok = proto.HasExtension(method.GetOptions(), test.E_Custom)
	require.Equal(t, true, ok, "HasExtension Custom")

	ok = proto.HasExtension(method.GetOptions(), test.E_Custom2)
	require.Equal(t, true, ok, "HasExtension Custom2")

	ok = proto.HasExtension(nil, test.E_Custom2)
	require.False(t, ok)

	emissing := *test.E_Custom
	emissing.Field = 50999
	ok = proto.HasExtension(method.GetOptions(), &emissing)
	require.Equal(t, false, ok, "HasExtension should not be found")

	// getextension - service
	m, err := proto.GetExtension(sd.GetOptions(), test.E_Svccustom)
	require.Nil(t, err)
	bptr, ok := m.(*bool)
	require.Equal(t, true, ok, "GetExtension Custom not *bool")
	require.Equal(t, true, *bptr, "GetExtension Custom not true")

	bval := proto.GetBoolExtension(sd.GetOptions(), test.E_Svccustom, false)
	require.Equal(t, true, bval, "GetExtension Custom not true")

	bval = proto.GetBoolExtension(sd.GetOptions(), test.E_Svccustom.Field, false)
	require.Equal(t, false, bval, "GetExtension Custom not true")

	// getextension - method
	m, err = proto.GetExtension(method.GetOptions(), test.E_Custom)
	require.Nil(t, err)
	bptr, ok = m.(*bool)
	require.Equal(t, true, ok, "GetExtension Custom not *bool")
	require.Equal(t, true, *bptr, "GetExtension Custom not true")

	bval = proto.GetBoolExtension(method.GetOptions(), test.E_Custom, false)
	require.Equal(t, true, bval, "GetExtension Custom not true")

	_, err = proto.GetExtension(nil, test.E_Custom)
	require.NotNil(t, err)

	_, err = proto.GetExtension(&test.TestResponse{}, test.E_Custom)
	require.NotNil(t, err)

	_, err = proto.GetExtension(method.GetOptions(), nil)
	require.NotNil(t, err)

	// clear extension
	proto.ClearExtension(method.GetOptions(), test.E_Custom)
	ok = proto.HasExtension(method.GetOptions(), test.E_Custom)
	require.Equal(t, false, ok, "ClearExtension Custom should not be found")

	m, err = proto.GetExtension(method.GetOptions(), test.E_Custom)
	require.Nil(t, m)
	require.NotNil(t, err)

	// set extension
	err = proto.SetExtension(method.GetOptions(), test.E_Custom, proto.Bool(false))
	require.Nil(t, err)
	ok = proto.HasExtension(method.GetOptions(), test.E_Custom)
	require.Equal(t, true, ok, "SetExtension Custom should be found")

	m, err = proto.GetExtension(method.GetOptions(), test.E_Custom)
	require.Nil(t, err)
	bptr, ok = m.(*bool)
	require.Equal(t, true, ok, "GetExtension Custom not *bool")
	require.Equal(t, false, *bptr, "GetExtension Custom not false after set")

	err = proto.SetExtension(nil, test.E_Custom, proto.Bool(false))
	require.NotNil(t, err)

	err = proto.SetExtension(method.GetOptions(), nil, proto.Bool(false))
	require.NotNil(t, err)

	err = proto.SetExtension(method.GetOptions(), test.E_Custom, "wrong type")
	require.NotNil(t, err)

	name := proto.MessageName(sd)
	require.Equal(t, "google.protobuf.ServiceDescriptorProto", name)

	mt := proto.MessageType(name)
	require.Equal(t, reflect.TypeOf((*dpb.ServiceDescriptorProto)(nil)), mt)

	// registered extensions
	e, err := proto.RegisteredExtensions(method.GetOptions(), (map[int32]*thisproto.ExtensionDesc)(nil))
	require.Nil(t, err)
	thisext, ok := e.(map[int32]*thisproto.ExtensionDesc)
	require.True(t, ok, "%T", e)
	require.True(t, ok)
	require.Equal(t, 2, len(thisext))
	require.Equal(t, "test.custom", thisext[50059].Name)
	require.Equal(t, "test.custom2", thisext[50060].Name)

	_, err = proto.RegisteredExtensions(method.GetOptions(), nil)
	require.Nil(t, err)
	thisext, ok = e.(map[int32]*thisproto.ExtensionDesc)
	require.True(t, ok, "%T", e)
	require.Equal(t, 2, len(thisext))
	require.Equal(t, "test.custom", thisext[50059].Name)
	require.Equal(t, "test.custom2", thisext[50060].Name)

	e, err = proto.RegisteredExtensions(method.GetOptions(), (map[int32]*otherproto.ExtensionDesc)(nil))
	require.Nil(t, err)
	otherext, ok := e.(map[int32]*otherproto.ExtensionDesc)
	require.True(t, ok)
	require.Equal(t, 2, len(otherext))
	require.Equal(t, "test.custom", otherext[50059].Name)
	require.Equal(t, "test.custom2", otherext[50060].Name)

	_, err = proto.RegisteredExtensions(method.GetOptions(), (*thisproto.ExtensionDesc)(nil))
	require.NotNil(t, err)

	_, err = proto.RegisteredExtensions(nil, (map[int32]*thisproto.ExtensionDesc)(nil))
	require.NotNil(t, err)

	_, err = proto.RegisteredExtensions(&test.TestResponse{}, (map[int32]*thisproto.ExtensionDesc)(nil))
	require.NotNil(t, err)

	var aliases = map[string]string{
		"google/protobuf/missing.proto": "github.com/golang/protobuf/ptypes/missing.proto",
	}
	p := NewProtoer(aliases)
	vv := p.FileDescriptor("google/protobuf/missing.proto")
	require.Nil(t, vv)

	vv = p.FileDescriptor("google/protobuf/missing2.proto")
	require.Nil(t, vv)
}

func TestNativeDescriptor(t *testing.T) {
	cases := []struct {
		in  interface{}
		out interface{}
		err error
	}{
		{nil, nil, nil},
		{"wrongtype", nil, fmt.Errorf("wrong type")},

		{(*dpbother.FileDescriptorSet)(nil), (*dpb.FileDescriptorSet)(nil), nil},
		{(*dpbother.FileDescriptorProto)(nil), (*dpb.FileDescriptorProto)(nil), nil},
		{(*dpbother.DescriptorProto)(nil), (*dpb.DescriptorProto)(nil), nil},
		{(*dpbother.ExtensionRangeOptions)(nil), (*dpb.ExtensionRangeOptions)(nil), nil},
		{(*dpbother.FieldDescriptorProto)(nil), (*dpb.FieldDescriptorProto)(nil), nil},
		{(*dpbother.OneofDescriptorProto)(nil), (*dpb.OneofDescriptorProto)(nil), nil},
		{(*dpbother.EnumDescriptorProto)(nil), (*dpb.EnumDescriptorProto)(nil), nil},
		{(*dpbother.EnumValueDescriptorProto)(nil), (*dpb.EnumValueDescriptorProto)(nil), nil},
		{(*dpbother.ServiceDescriptorProto)(nil), (*dpb.ServiceDescriptorProto)(nil), nil},
		{(*dpbother.MethodDescriptorProto)(nil), (*dpb.MethodDescriptorProto)(nil), nil},
		{(*dpbother.FileOptions)(nil), (*dpb.FileOptions)(nil), nil},
		{(*dpbother.MessageOptions)(nil), (*dpb.MessageOptions)(nil), nil},
		{(*dpbother.FieldOptions)(nil), (*dpb.FieldOptions)(nil), nil},
		{(*dpbother.OneofOptions)(nil), (*dpb.OneofOptions)(nil), nil},
		{(*dpbother.EnumOptions)(nil), (*dpb.EnumOptions)(nil), nil},
		{(*dpbother.EnumValueOptions)(nil), (*dpb.EnumValueOptions)(nil), nil},
		{(*dpbother.ServiceOptions)(nil), (*dpb.ServiceOptions)(nil), nil},
		{(*dpbother.MethodOptions)(nil), (*dpb.MethodOptions)(nil), nil},
		{(*dpbother.UninterpretedOption)(nil), (*dpb.UninterpretedOption)(nil), nil},
		{(*dpbother.SourceCodeInfo)(nil), (*dpb.SourceCodeInfo)(nil), nil},
		{(*dpbother.GeneratedCodeInfo)(nil), (*dpb.GeneratedCodeInfo)(nil), nil},
		{&dpbother.FileDescriptorSet{}, &dpb.FileDescriptorSet{}, nil},

		{&dpbother.FileDescriptorSet{}, &dpb.FileDescriptorSet{}, nil},
		{&dpbother.FileDescriptorProto{}, &dpb.FileDescriptorProto{}, nil},
		{&dpbother.DescriptorProto{}, &dpb.DescriptorProto{}, nil},
		{&dpbother.ExtensionRangeOptions{}, &dpb.ExtensionRangeOptions{}, nil},
		{&dpbother.FieldDescriptorProto{}, &dpb.FieldDescriptorProto{}, nil},
		{&dpbother.OneofDescriptorProto{}, &dpb.OneofDescriptorProto{}, nil},
		{&dpbother.EnumDescriptorProto{}, &dpb.EnumDescriptorProto{}, nil},
		{&dpbother.EnumValueDescriptorProto{}, &dpb.EnumValueDescriptorProto{}, nil},
		{&dpbother.ServiceDescriptorProto{}, &dpb.ServiceDescriptorProto{}, nil},
		{&dpbother.MethodDescriptorProto{}, &dpb.MethodDescriptorProto{}, nil},
		{&dpbother.FileOptions{}, &dpb.FileOptions{}, nil},
		{&dpbother.MessageOptions{}, &dpb.MessageOptions{}, nil},
		{&dpbother.FieldOptions{}, &dpb.FieldOptions{}, nil},
		{&dpbother.OneofOptions{}, &dpb.OneofOptions{}, nil},
		{&dpbother.EnumOptions{}, &dpb.EnumOptions{}, nil},
		{&dpbother.EnumValueOptions{}, &dpb.EnumValueOptions{}, nil},
		{&dpbother.ServiceOptions{}, &dpb.ServiceOptions{}, nil},
		{&dpbother.MethodOptions{}, &dpb.MethodOptions{}, nil},
		{&dpbother.UninterpretedOption{}, &dpb.UninterpretedOption{}, nil},
		{&dpbother.SourceCodeInfo{}, &dpb.SourceCodeInfo{}, nil},
		{&dpbother.GeneratedCodeInfo{}, &dpb.GeneratedCodeInfo{}, nil},
	}

	for i, c := range cases {
		ctx := func() string { return fmt.Sprintf("test case %d", i) }
		out, err := ToNativeDescriptor(c.in)
		if c.err != nil {
			require.NotNil(t, err, ctx())
		} else {
			require.Nil(t, err, ctx())
		}
		require.Equal(t, c.out, out)
	}
}

func TestNativeExtensionDesc(t *testing.T) {
	in := otherprotos.E_Custom
	out, err := ToNativeExtensionDesc(in)
	require.Nil(t, err)
	require.NotNil(t, out)
	require.Equal(t, in.ExtensionType, out.ExtensionType)
	require.Equal(t, in.Field, out.Field)
	require.Equal(t, in.Name, out.Name)
	require.Equal(t, in.Tag, out.Tag)
	require.Equal(t, in.Filename, out.Filename)

	//test cached
	out, err = ToNativeExtensionDesc(in)
	require.Nil(t, err)
	require.NotNil(t, out)
	require.Equal(t, in.ExtensionType, out.ExtensionType)
	require.Equal(t, in.Field, out.Field)
	require.Equal(t, in.Name, out.Name)
	require.Equal(t, in.Tag, out.Tag)
	require.Equal(t, in.Filename, out.Filename)
}

func TestDefaultValues(t *testing.T) {
	file := "desc_test_defaults.proto"
	fd, err := loadFileDescriptorProto(file)
	require.Nil(t, err)

	l := makeLookup(fd)

	testCases := []struct {
		message, field string
		defaultVal     interface{}
	}{
		{"test.PrimitiveDefaults", "fl32", float32(3.14159)},
		{"test.PrimitiveDefaults", "fl64", 3.14159},
		{"test.PrimitiveDefaults", "fl32d", float32(6.022140857e23)},
		{"test.PrimitiveDefaults", "fl64d", 6.022140857e23},
		{"test.PrimitiveDefaults", "bl1", true},
		{"test.PrimitiveDefaults", "bl2", false},
		{"test.PrimitiveDefaults", "i32", int32(10101)},
		{"test.PrimitiveDefaults", "i32n", int32(-10101)},
		{"test.PrimitiveDefaults", "i32x", int32(0x20202)},
		{"test.PrimitiveDefaults", "i32xn", int32(-0x20202)},
		{"test.PrimitiveDefaults", "i64", int64(10101)},
		{"test.PrimitiveDefaults", "i64n", int64(-10101)},
		{"test.PrimitiveDefaults", "i64x", int64(0x20202)},
		{"test.PrimitiveDefaults", "i64xn", int64(-0x20202)},
		{"test.PrimitiveDefaults", "i32s", int32(10101)},
		{"test.PrimitiveDefaults", "i32sn", int32(-10101)},
		{"test.PrimitiveDefaults", "i32sx", int32(0x20202)},
		{"test.PrimitiveDefaults", "i32sxn", int32(-0x20202)},
		{"test.PrimitiveDefaults", "i64s", int64(10101)},
		{"test.PrimitiveDefaults", "i64sn", int64(-10101)},
		{"test.PrimitiveDefaults", "i64sx", int64(0x20202)},
		{"test.PrimitiveDefaults", "i64sxn", int64(-0x20202)},
		{"test.PrimitiveDefaults", "i32f", int32(10101)},
		{"test.PrimitiveDefaults", "i32fn", int32(-10101)},
		{"test.PrimitiveDefaults", "i32fx", int32(0x20202)},
		{"test.PrimitiveDefaults", "i32fxn", int32(-0x20202)},
		{"test.PrimitiveDefaults", "i64f", int64(10101)},
		{"test.PrimitiveDefaults", "i64fn", int64(-10101)},
		{"test.PrimitiveDefaults", "i64fx", int64(0x20202)},
		{"test.PrimitiveDefaults", "i64fxn", int64(-0x20202)},
		{"test.PrimitiveDefaults", "u32", uint32(10101)},
		{"test.PrimitiveDefaults", "u32x", uint32(0x20202)},
		{"test.PrimitiveDefaults", "u64", uint64(10101)},
		{"test.PrimitiveDefaults", "u64x", uint64(0x20202)},
		{"test.PrimitiveDefaults", "u32f", uint32(10101)},
		{"test.PrimitiveDefaults", "u32fx", uint32(0x20202)},
		{"test.PrimitiveDefaults", "u64f", uint64(10101)},
		{"test.PrimitiveDefaults", "u64fx", uint64(0x20202)},
		{"test.StringAndBytesDefaults", "dq", "this is a string with \"nested quotes\""},
		{"test.StringAndBytesDefaults", "sq", "this is a string with \"nested quotes\""},
	}

	for i, tc := range testCases {
		fqn := fmt.Sprintf("%s.%s", tc.message, tc.field)
		def := l.field[fqn].GetDefaultValue()
		require.Equal(t, fmt.Sprintf("%v", tc.defaultVal), def, "wrong default value for case %d: %s.%s", i, tc.message, tc.field)
	}
}

// loadFileDescriptor loads a registered descriptor and decodes it. If the given
// name cannot be loaded but is a known standard name, an alias will be tried by the proto,
// so the standard files can be loaded even if linked against older "known bad"
// versions of packages.
func loadFileDescriptorProto(file string) (*dpb.FileDescriptorProto, error) {
	fdb := proto.FileDescriptor(file)
	if fdb == nil {
		return nil, fmt.Errorf("Missing file descriptor %s.", file)
	}

	fd, err := decodeFileDescriptorProto(file, fdb)
	if err != nil {
		return nil, err
	}

	// the file descriptor may have been laoded with an alias,
	// so we ensure the specified name to ensure it can be linked.
	fd.Name = proto.String(file)

	return fd, nil
}

// decodeFileDescriptorProto decodes the bytes of a registered file descriptor.
// Registered file descriptors are first "proto encoded" (e.g. binary format
// for the descriptor protos) and then gzipped. So this function gunzips and
// then unmarshals into a descriptor proto.
func decodeFileDescriptorProto(element string, fdb []byte) (*dpb.FileDescriptorProto, error) {
	raw, err := decompress(fdb)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress %q descriptor: %v", element, err)
	}
	fd := dpb.FileDescriptorProto{}
	if err := proto.Unmarshal(raw, &fd); err != nil {
		return nil, fmt.Errorf("bad descriptor for %q: %v", element, err)
	}
	return &fd, nil
}

func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	return out, nil
}

func makeLookup(fd *dpb.FileDescriptorProto) *lookup {
	l := &lookup{
		message: make(map[string]*dpb.DescriptorProto),
		field:   make(map[string]*dpb.FieldDescriptorProto),
		enum:    make(map[string]*dpb.EnumDescriptorProto),
		method:  make(map[string]*dpb.MethodDescriptorProto),
		service: make(map[string]*dpb.ServiceDescriptorProto),
	}

	pkg := fd.GetPackage()

	merge := func(a, b string) string {
		if a == "" {
			return b
		} else {
			return a + "." + b
		}
	}

	for _, message := range fd.MessageType {
		fqn := merge(pkg, message.GetName())
		l.message[fqn] = message
		for _, field := range message.Field {
			fqnField := fmt.Sprintf("%s.%s", fqn, field.GetName())
			l.field[fqnField] = field
		}
	}
	for _, enum := range fd.EnumType {
		fqn := merge(pkg, enum.GetName())
		l.enum[fqn] = enum
	}
	for _, service := range fd.Service {
		fqn := merge(pkg, service.GetName())
		l.service[fqn] = service
		for _, method := range service.Method {
			fqnMethod := fmt.Sprintf("/%s/%s", fqn, method.GetName())
			l.method[fqnMethod] = method
		}
	}

	return l
}
