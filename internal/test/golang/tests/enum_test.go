package tests

import (
	"reflect"
	"testing"

	"github.com/prasek/protoer/examples/desc"
	"github.com/prasek/protoer/internal/test/golang/testprotos"
	"github.com/prasek/protoer/internal/test/testutil"
)

func TestLoadEnumDescriptor(t *testing.T) {
	ed, err := desc.LoadEnumDescriptorForEnum(testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum(0))
	testutil.Ok(t, err)
	testutil.Eq(t, "DeeplyNestedEnum", ed.GetName())
	testutil.Eq(t, "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage.DeeplyNestedEnum", ed.GetFullyQualifiedName())
	fd := ed.GetFile()
	testutil.Eq(t, "desc_test1.proto", fd.GetName())
	ofd, err := desc.LoadFileDescriptor("desc_test1.proto")
	testutil.Ok(t, err)
	testutil.Eq(t, ofd, fd)

	ed2, err := desc.LoadEnumDescriptorForEnum((*testprotos.TestEnum)(nil)) // pointer type for interface
	testutil.Ok(t, err)
	testutil.Eq(t, "TestEnum", ed2.GetName())
	testutil.Eq(t, "testprotos.TestEnum", ed2.GetFullyQualifiedName())
	fd = ed2.GetFile()
	testutil.Eq(t, "desc_test_field_types.proto", fd.GetName())
	ofd, err = desc.LoadFileDescriptor("desc_test_field_types.proto")
	testutil.Ok(t, err)
	testutil.Eq(t, ofd, fd)
	testutil.Eq(t, fd, ed2.GetParent())

	// now use the APIs that take reflect.Type
	ed3, err := desc.LoadEnumDescriptorForType(reflect.TypeOf((*testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum)(nil)))
	testutil.Ok(t, err)
	testutil.Eq(t, ed, ed3)

	ed4, err := desc.LoadEnumDescriptorForType(reflect.TypeOf(testprotos.TestEnum_FIRST))
	testutil.Ok(t, err)
	testutil.Eq(t, ed2, ed4)
}

func TestEnumDescriptorFindValue(t *testing.T) {
	fd, err := desc.LoadFileDescriptor("desc_test_defaults.proto")
	testutil.Ok(t, err)
	ed, ok := fd.FindSymbol("testprotos.Number").(*desc.EnumDescriptor)
	testutil.Eq(t, true, ok)
	lastNumber := int32(-1)
	for _, vd := range ed.GetValues() {
		found := ed.FindValueByName(vd.GetName())
		testutil.Eq(t, vd, found)
		found = ed.FindValueByNumber(vd.GetNumber())
		if lastNumber == vd.GetNumber() {
			// found value will be the first one with the given number, not this one
			testutil.Eq(t, false, vd == found)
		} else {
			testutil.Eq(t, vd, found)
			lastNumber = vd.GetNumber()
		}
	}
	testutil.Eq(t, (*desc.EnumValueDescriptor)(nil), ed.FindValueByName("junk name"))
	testutil.Eq(t, (*desc.EnumValueDescriptor)(nil), ed.FindValueByNumber(99999))
}
