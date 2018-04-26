package desc

import "github.com/prasek/protoer/proto"

var _ Descriptor = (*MethodDescriptor)(nil)

func (md *MethodDescriptor) GetName() string {
	return md.proto.GetName()
}

func (md *MethodDescriptor) GetFullyQualifiedName() string {
	return md.fqn
}

func (md *MethodDescriptor) GetParent() Descriptor {
	return md.parent
}

// GetService returns the RPC service in which this method is declared.
func (md *MethodDescriptor) GetService() *ServiceDescriptor {
	return md.parent
}

func (md *MethodDescriptor) GetFile() *FileDescriptor {
	return md.file
}

func (md *MethodDescriptor) String() string {
	return md.proto.String()
}

// IsServerStreaming returns true if this is a server-streaming method.
func (md *MethodDescriptor) IsServerStreaming() bool {
	return md.proto.GetServerStreaming()
}

// IsClientStreaming returns true if this is a client-streaming method.
func (md *MethodDescriptor) IsClientStreaming() bool {
	return md.proto.GetClientStreaming()
}

// GetInputType returns the input type, or request type, of the RPC method.
func (md *MethodDescriptor) GetInputType() *MessageDescriptor {
	return md.inType
}

// GetOutputType returns the output type, or response type, of the RPC method.
func (md *MethodDescriptor) GetOutputType() *MessageDescriptor {
	return md.outType
}

func (md *MethodDescriptor) GetBoolExtension(ext interface{}, def bool) bool {
	return GetBoolExtension(md.proto.Options, ext, def)
}

func (md *MethodDescriptor) GetExtension(ext interface{}) (extval interface{}, err error) {
	return proto.GetExtension(md.proto.Options, ext)
}

func (md *MethodDescriptor) RegisteredExtensions(desiredType interface{}) (extensions interface{}, err error) {
	return proto.RegisteredExtensions(md.proto.Options, desiredType)
}
