package tests

import (
	"os"
	"testing"

	"github.com/prasek/protoer/proto"
	"github.com/prasek/protoer/proto/golang"
)

func TestMain(m *testing.M) {
	proto.SetProtoer(golang.NewProtoer(nil))
	code := m.Run()
	os.Exit(code)
}
