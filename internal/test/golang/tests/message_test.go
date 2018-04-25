package tests

import (
	"reflect"
	"testing"

	"github.com/prasek/protoer/examples/desc"
	"github.com/prasek/protoer/internal/test/golang/testprotos"
	"github.com/prasek/protoer/internal/test/testutil"
)

func TestLoadMessageDescriptor(t *testing.T) {
	// loading enclosed messages should return the same descriptor
	// and have a reference to the same file descriptor
	md, err := desc.LoadMessageDescriptor("testprotos.TestMessage")
	testutil.Ok(t, err)
	testutil.Eq(t, "TestMessage", md.GetName())
	testutil.Eq(t, "testprotos.TestMessage", md.GetFullyQualifiedName())
	fd := md.GetFile()
	testutil.Eq(t, "desc_test1.proto", fd.GetName())
	testutil.Eq(t, fd, md.GetParent())

	md2, err := desc.LoadMessageDescriptorForMessage((*testprotos.TestMessage)(nil))
	testutil.Ok(t, err)
	testutil.Eq(t, md, md2)

	md3, err := desc.LoadMessageDescriptorForType(reflect.TypeOf((*testprotos.TestMessage)(nil)))
	testutil.Ok(t, err)
	testutil.Eq(t, md, md3)
}

func TestMessageDescriptorFindField(t *testing.T) {
	md, err := desc.LoadMessageDescriptor("testprotos.Frobnitz")
	testutil.Ok(t, err)
	for _, fd := range md.GetFields() {
		found := md.FindFieldByName(fd.GetName())
		testutil.Eq(t, fd, found)
		found = md.FindFieldByNumber(fd.GetNumber())
		testutil.Eq(t, fd, found)
	}
	testutil.Eq(t, (*desc.FieldDescriptor)(nil), md.FindFieldByName("junk name"))
	testutil.Eq(t, (*desc.FieldDescriptor)(nil), md.FindFieldByNumber(99999))
}
