package proto

import (
	"reflect"
	"sync"
)

var (
	mu      sync.RWMutex
	protoer Protoer
)

func SetProtoer(up UntypedProtoer) {
	mu.Lock()
	defer mu.Unlock()
	protoer = ToProtoer(up)
}

func ToProtoer(up UntypedProtoer) Protoer {
	return &typedProtoer{up: up}
}

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

type Aliaser interface {
	Aliases() map[string]string
}

func Marshal(m Message) ([]byte, error) {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.Marshal(m)
}

func Unmarshal(b []byte, m Message) error {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.Unmarshal(b, m)
}

func Clone(m Message) Message {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.Clone(m)
}

func Equal(m1, m2 Message) bool {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.Equal(m1, m2)
}

func Merge(dst, src Message) {
	mu.RLock()
	defer mu.RUnlock()
	protoer.Merge(dst, src)
}

func Reset(m Message) {
	mu.RLock()
	defer mu.RUnlock()
	protoer.Reset(m)
}

func Size(m Message) int {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.Size(m)
}

func HasExtension(m Message, ext interface{}) bool {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.HasExtension(m, ext)
}

func ClearExtension(m Message, ext interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	protoer.ClearExtension(m, ext)
}

func SetExtension(m Message, ext interface{}, v interface{}) error {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.SetExtension(m, ext, v)
}

func GetExtension(m Message, ext interface{}) (extval interface{}, err error) {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.GetExtension(m, ext)
}

func RegisteredExtensions(m Message, desiredType interface{}) (interface{}, error) {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.RegisteredExtensions(m, desiredType)
}

func FileDescriptor(file string) []byte {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.FileDescriptor(file)
}

func MessageName(m Message) string {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.MessageName(m)
}

func MessageType(name string) reflect.Type {
	mu.RLock()
	defer mu.RUnlock()
	return protoer.MessageType(name)
}

func Aliases() map[string]string {
	mu.RLock()
	defer mu.RUnlock()
	if a, ok := protoer.(Aliaser); ok {
		return a.Aliases()
	}
	return map[string]string{}
}

type typedProtoer struct {
	up UntypedProtoer
}

func (tp *typedProtoer) Marshal(m Message) ([]byte, error) {
	return tp.up.Marshal(m)
}

func (tp *typedProtoer) Unmarshal(b []byte, m Message) error {
	return tp.up.Unmarshal(b, m)
}

func (tp *typedProtoer) Clone(m Message) Message {
	return tp.up.Clone(m).(Message)
}

func (tp *typedProtoer) Equal(m1, m2 Message) bool {
	return tp.up.Equal(m1, m2)
}

func (tp *typedProtoer) Merge(dst, src Message) {
	tp.up.Merge(dst, src)
}

func (tp *typedProtoer) Reset(m Message) {
	tp.up.Reset(m)
}

func (tp *typedProtoer) Size(m Message) int {
	return tp.up.Size(m)
}

func (tp *typedProtoer) HasExtension(m Message, ext interface{}) bool {
	return tp.up.HasExtension(m, ext)
}

func (tp *typedProtoer) ClearExtension(m Message, ext interface{}) {
	tp.up.ClearExtension(m, ext)
}

func (tp *typedProtoer) SetExtension(m Message, ext interface{}, v interface{}) error {
	return tp.up.SetExtension(m, ext, v)
}

func (tp *typedProtoer) GetExtension(m Message, ext interface{}) (extval interface{}, err error) {
	return tp.up.GetExtension(m, ext)
}

func (tp *typedProtoer) RegisteredExtensions(m Message, desiredType interface{}) (extensions interface{}, err error) {
	return tp.up.RegisteredExtensions(m, desiredType)
}

func (tp *typedProtoer) FileDescriptor(file string) []byte {
	return tp.up.FileDescriptor(file)
}

func (tp *typedProtoer) MessageName(m Message) string {
	return tp.up.MessageName(m)
}

func (tp *typedProtoer) MessageType(name string) reflect.Type {
	return tp.up.MessageType(name)
}

func (tp *typedProtoer) Aliases() map[string]string {
	if a, ok := tp.up.(Aliaser); ok {
		return a.Aliases()
	}
	return map[string]string{}
}

/*
 * Helper routines for simplifying the creation of optional fields of basic type.
 */

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	return &v
}

// Int32 is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it.
func Int32(v int32) *int32 {
	return &v
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
func Int(v int) *int32 {
	p := new(int32)
	*p = int32(v)
	return p
}

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 {
	return &v
}

// Float32 is a helper routine that allocates a new float32 value
// to store v and returns a pointer to it.
func Float32(v float32) *float32 {
	return &v
}

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 {
	return &v
}

// Uint32 is a helper routine that allocates a new uint32 value
// to store v and returns a pointer to it.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint64 is a helper routine that allocates a new uint64 value
// to store v and returns a pointer to it.
func Uint64(v uint64) *uint64 {
	return &v
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	return &v
}
