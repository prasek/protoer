package tests

import (
	"os"
	"testing"

	"github.com/prasek/protoer/proto"
	"github.com/prasek/protoer/proto/gogo"
)

func TestMain(m *testing.M) {
	proto.SetProtoer(gogo.NewProtoer(nil))
	code := m.Run()
	os.Exit(code)
}
