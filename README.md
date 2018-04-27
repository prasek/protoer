# Protoer - a pluggable proto package
[![Build Status](https://travis-ci.org/prasek/protoer.svg?branch=master)](https://travis-ci.org/prasek/protoer/branches)
[![codecov](https://codecov.io/gh/prasek/protoer/branch/master/graph/badge.svg)](https://codecov.io/gh/prasek/protoer)

## Reference library designed for embedding with zero dependencies
This is a reference library intended for embedding into libraries that want to 
use the proto package, but also want to allow users to pick gogo or golang as
the underlying implementation. 

The interfaces are defined to have zero additional dependencies and 
protoer.go can be easily embedded as an interop layer for the proto package.

## High-level API of the proto package
Support for the most commonly used proto package features are provided.

The UntypedProtoer interface has no dependencies and allows any compatible
type to be used as a protoer.

```go
type Protoer interface {
	Marshal(m Message) ([]byte, error)
	Unmarshal(b []byte, m Message) error

	Clone(m Message) Message
	Equal(m1, m2 Message) bool
	Merge(dst, src Message)
	Reset(m Message)
	Size(m Message) int

	HasExtension(m Message, ext interface{}) bool
	ClearExtension(m Message, ext interface{})
	GetExtension(m Message, ext interface{}) (interface{}, error)
	SetExtension(m Message, ext interface{}, v interface{}) error
	RegisteredExtensions(m Message, desiredType interface{}) (interface{}, error)

	FileDescriptor(file string) []byte
	MessageName(m Message) string
	MessageType(name string) reflect.Type
}

type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

type UntypedProtoer interface {
	Marshal(m interface{}) ([]byte, error)
	Unmarshal(b []byte, m interface{}) error

	Clone(m interface{}) interface{}
	Equal(m1, m2 interface{}) bool
	Merge(dst, src interface{})
	Reset(m interface{})
	Size(m interface{}) int

	HasExtension(m interface{}, ext interface{}) bool
	ClearExtension(m interface{}, ext interface{})
	GetExtension(m interface{}, ext interface{}) (interface{}, error)
	SetExtension(m interface{}, ext interface{}, v interface{}) error
	RegisteredExtensions(m interface{}, desiredType interface{}) (interface{}, error)

	FileDescriptor(file string) []byte
	MessageName(m interface{}) string
	MessageType(name string) reflect.Type
}
```

## Getting Started
To get started see examples/desc for an example library, based on
github.com/jhump/protoreflect/desc.

The desc package initializes with a default protoer.

```go
import(
  "github.com/prasek/protoer/proto"
  "github.com/prasek/protoer/proto/golang"
)

func init() {
	proto.SetProtoer(golang.NewProtoer(nil))
}
```

Then uses the protoer/proto package for commonly used proto features.

```go
func GetBoolExtension(pb proto.Message, ext interface{}, def bool) bool {
	v, err := proto.GetExtension(pb, ext)
	if err != nil {
		return def
	}
	if v == nil {
		return def
	}
	return *v.(*bool)
}
```

Desc clients can later choose to use gogo/protobuf/proto instead:
```go
import(
  "github.com/prasek/protoer/proto"
  "github.com/prasek/protoer/proto/gogo"
  "github.com/prasek/protoer/examples/desc"
)

func init() {
	proto.SetProtoer(gogo.NewProtoer(nil))
}

func main() {
  desc.GetBoolExtension(pb, E_Custom, false)
}
```

To get zero additional dependencies, the desc package could copy the 
relevant portion of the proto package into it's repo, for example:
* proto/protoer.go
* proto/golang/protoer.go
