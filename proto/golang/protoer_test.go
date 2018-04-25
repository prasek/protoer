package golang

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/prasek/protoer/internal/test/testutil"
	"github.com/prasek/protoer/proto"

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
		testutil.Ok(t, err)
		testutil.Eq(t, file, fd.GetName())

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
		testutil.Ok(t, err)
		testutil.Eq(t, file, fd.GetName())
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
