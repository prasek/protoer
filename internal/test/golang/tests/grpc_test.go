package tests

import (
	"fmt"
	"reflect"
	"testing"

	"google.golang.org/grpc"

	gogo "github.com/gogo/protobuf/proto"
	golang "github.com/golang/protobuf/proto"
	"github.com/prasek/protoer/examples/desc"
	"github.com/prasek/protoer/internal/test/golang/testprotos"
	"github.com/prasek/protoer/internal/test/testutil"
)

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testprotos.TestService",
	HandlerType: (*testprotos.TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName: "DoSomethingElse",
		},
		{
			StreamName: "DoSomethingAgain",
		},
		{
			StreamName: "DoSomethingForever",
		},
	},
	Metadata: "desc_test_proto3.proto",
}

type testService struct {
	testprotos.TestServiceServer
}

func TestMethodExtension(t *testing.T) {
	/*
	   func (gio *GrpcInterceptorOptions) IsTx(info *grpc.UnaryServerInfo) bool {
	   	method := gio.methods[info.FullMethod]
	   	tx := proto.GetBoolExtension(method.GetMethodOptions(), stub.E_Tx, false)
	   	return tx
	   }
	*/

}

func TestLoadServiceDescriptors(t *testing.T) {
	s := grpc.NewServer()
	testprotos.RegisterTestServiceServer(s, testService{})
	sds, err := desc.LoadServiceDescriptors(s)
	testutil.Ok(t, err, err)
	testutil.Eq(t, 1, len(sds), "service descriptor len")
	sd := sds["testprotos.TestService"]

	cases := []struct {
		method, request, response string
		opt                       bool
		custom2                   *testprotos.CustomOption
		extlen                    int
	}{
		{"DoSomething", "testprotos.TestRequest", "jhump.protoreflect.desc.Bar",
			true, &testprotos.CustomOption{Name: "test123", Value: 55}, 5},
		{"DoSomethingElse", "testprotos.TestMessage", "testprotos.TestResponse",
			false, nil, 0},
		{"DoSomethingAgain", "jhump.protoreflect.desc.Bar", "testprotos.AnotherTestMessage",
			false, nil, 5},
		{"DoSomethingForever", "testprotos.TestRequest", "testprotos.TestResponse",
			false, nil, 0},
	}

	testutil.Eq(t, len(cases), len(sd.GetMethods()))

	for i, c := range cases {
		md := sd.GetMethods()[i]
		testutil.Eq(t, c.method, md.GetName())
		testutil.Eq(t, c.request, md.GetInputType().GetFullyQualifiedName())
		testutil.Eq(t, c.response, md.GetOutputType().GetFullyQualifiedName())
		testutil.Eq(t, c.opt, md.GetBoolExtension(testprotos.E_Custom, false))

		//default extensions
		v, err := md.RegisteredExtensions(nil)
		if v != nil {
			testutil.Ok(t, err)
		}
		var extDefault map[int32]*golang.ExtensionDesc = nil
		if v != nil && reflect.ValueOf(v).Len() > 0 {
			extDefault = v.(map[int32]*golang.ExtensionDesc)
		}

		//gogo extensiondesc
		desiredType := (map[int32]*gogo.ExtensionDesc)(nil)
		v, err = md.RegisteredExtensions(desiredType)
		if v != nil {
			testutil.Ok(t, err)
		}
		var extGogo map[int32]*gogo.ExtensionDesc = nil
		if v != nil && reflect.ValueOf(v).Len() > 0 {
			extGogo = v.(map[int32]*gogo.ExtensionDesc)
		}

		//golang extensiondesc
		desiredTypeGolang := (map[int32]*golang.ExtensionDesc)(nil)
		v, err = md.RegisteredExtensions(desiredTypeGolang)
		if v != nil {
			testutil.Ok(t, err)
		}
		var extGolang map[int32]*golang.ExtensionDesc = nil
		if v != nil && reflect.ValueOf(v).Len() > 0 {
			extGolang = v.(map[int32]*golang.ExtensionDesc)
		}

		if extDefault != nil {
			testutil.NotNil(t, extGogo, "gogo extensions are nil when they should be present")
			testutil.NotNil(t, extGolang, "golang extensions are nil when they should be present")
		} else {
			testutil.Nil(t, extGogo, "golang extensions are not nil when they should be")
			testutil.Nil(t, extGolang, "golang extensions are not nil when they should be")
		}

		//verify same
		if extGogo != nil && extGolang != nil && extDefault != nil {
			fmt.Printf("\nVerifying:\n")
			testutil.Eq(t, len(extGogo), c.extlen)
			other := []map[int32]*golang.ExtensionDesc{extDefault, extGolang}
			for _, check := range other {
				testutil.Eq(t, len(check), c.extlen)

				for k, v := range extGogo {
					//fmt.Printf(" - map[%d]\n", k)
					testutil.Eq(t, check[k].ExtendedType, v.ExtendedType)
					testutil.Eq(t, check[k].ExtensionType, v.ExtensionType)
					testutil.Eq(t, check[k].Field, v.Field)
					testutil.Eq(t, check[k].Name, v.Name)
					testutil.Eq(t, check[k].Tag, v.Tag)
					testutil.Eq(t, check[k].Filename, v.Filename)
				}
			}
		}

		v, err = md.GetExtension(testprotos.E_Custom2)
		if v != nil {
			testutil.Ok(t, err)
		}
		if c.custom2 == nil {
			testutil.Nil(t, v)
		} else {
			testutil.Eq(t, c.custom2, v)
		}
	}
}

func TestLoadServiceDescriptor(t *testing.T) {
	s := grpc.NewServer()
	testprotos.RegisterTestServiceServer(s, testService{})
	sd, err := desc.LoadServiceDescriptor(&_TestService_serviceDesc)
	testutil.Nil(t, err)
	testutil.NotNil(t, sd, "Service descriptor missing")

	cases := []struct{ method, request, response string }{
		{"DoSomething", "testprotos.TestRequest", "jhump.protoreflect.desc.Bar"},
		{"DoSomethingElse", "testprotos.TestMessage", "testprotos.TestResponse"},
		{"DoSomethingAgain", "jhump.protoreflect.desc.Bar", "testprotos.AnotherTestMessage"},
		{"DoSomethingForever", "testprotos.TestRequest", "testprotos.TestResponse"},
	}

	testutil.Eq(t, len(cases), len(sd.GetMethods()))

	for i, c := range cases {
		md := sd.GetMethods()[i]
		testutil.Eq(t, c.method, md.GetName())
		testutil.Eq(t, c.request, md.GetInputType().GetFullyQualifiedName())
		testutil.Eq(t, c.response, md.GetOutputType().GetFullyQualifiedName())
	}
}
