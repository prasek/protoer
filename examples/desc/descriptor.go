package desc

import dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Descriptor is the common interface implemented by all descriptor objects.
type Descriptor interface {
	// GetName returns the name of the object described by the descriptor. This will
	// be a base name that does not include enclosing message names or the package name.
	// For file descriptors, this indicates the path and name to the described file.
	GetName() string
	// GetFullyQualifiedName returns the fully-qualified name of the object described by
	// the descriptor. This will include the package name and any enclosing message names.
	// For file descriptors, this returns the path and name to the described file (same as
	// GetName).
	GetFullyQualifiedName() string
	// GetParent returns the enclosing element in a proto source file. If the described
	// object is a top-level object, this returns the file descriptor. Otherwise, it returns
	// the element in which the described object was declared. File descriptors have no
	// parent and return nil.
	GetParent() Descriptor
	// GetFile returns the file descriptor in which this element was declared. File
	// descriptors return themselves.
	GetFile() *FileDescriptor

	/*
		// GetOptions returns the options proto containing options for the described element.
		GetOptions() Message

		// AsProto() returns the underlying proto
		AsProto() Message

			// GetSourceInfo returns any source code information that was present in the file
			// descriptor. Source code info is optional. If no source code info is available for
			// the element (including if there is none at all in the file descriptor) then this
			// returns nil
			GetSourceInfo() *dpb.SourceCodeInfo_Location
	*/
}

// scope represents a lexical scope in a proto file in which messages and enums
// can be declared.
type scope func(string) Descriptor

// FileDescriptor describes a proto source file.
type FileDescriptor struct {
	proto      *dpb.FileDescriptorProto
	symbols    map[string]Descriptor
	deps       []*FileDescriptor
	publicDeps []*FileDescriptor
	weakDeps   []*FileDescriptor
	messages   []*MessageDescriptor
	enums      []*EnumDescriptor
	extensions []*FieldDescriptor
	services   []*ServiceDescriptor
	fieldIndex map[string]map[int32]*FieldDescriptor
	isProto3   bool
}

// ServiceDescriptor describes an RPC service declared in a proto file.
type ServiceDescriptor struct {
	proto      *dpb.ServiceDescriptorProto
	file       *FileDescriptor
	methods    []*MethodDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}

// MethodDescriptor describes an RPC method declared in a proto file.
type MethodDescriptor struct {
	proto      *dpb.MethodDescriptorProto
	parent     *ServiceDescriptor
	file       *FileDescriptor
	inType     *MessageDescriptor
	outType    *MessageDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}

// MessageDescriptor describes a protocol buffer message.
type MessageDescriptor struct {
	proto      *dpb.DescriptorProto
	parent     Descriptor
	file       *FileDescriptor
	fields     []*FieldDescriptor
	nested     []*MessageDescriptor
	enums      []*EnumDescriptor
	extensions []*FieldDescriptor
	oneOfs     []*OneOfDescriptor
	//extRanges  extRanges
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
	isProto3   bool
	isMapEntry bool
}

// FieldDescriptor describes a field of a protocol buffer message.
type FieldDescriptor struct {
	proto      *dpb.FieldDescriptorProto
	parent     Descriptor
	owner      *MessageDescriptor
	file       *FileDescriptor
	oneOf      *OneOfDescriptor
	msgType    *MessageDescriptor
	enumType   *EnumDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
	isMap      bool
}

// EnumDescriptor describes an enum declared in a proto file.
type EnumDescriptor struct {
	proto       *dpb.EnumDescriptorProto
	parent      Descriptor
	file        *FileDescriptor
	values      []*EnumValueDescriptor
	valuesByNum sortedValues
	fqn         string
	sourceInfo  *dpb.SourceCodeInfo_Location
}

// EnumValueDescriptor describes an allowed value of an enum declared in a proto file.
type EnumValueDescriptor struct {
	proto      *dpb.EnumValueDescriptorProto
	parent     *EnumDescriptor
	file       *FileDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}

// OneOfDescriptor describes a one-of field set declared in a protocol buffer message.
type OneOfDescriptor struct {
	proto      *dpb.OneofDescriptorProto
	parent     *MessageDescriptor
	file       *FileDescriptor
	choices    []*FieldDescriptor
	fqn        string
	sourceInfo *dpb.SourceCodeInfo_Location
}
