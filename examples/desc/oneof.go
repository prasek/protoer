package desc

var _ Descriptor = (*OneOfDescriptor)(nil)

func (od *OneOfDescriptor) GetName() string {
	return od.proto.GetName()
}

func (od *OneOfDescriptor) GetFullyQualifiedName() string {
	return od.fqn
}

func (od *OneOfDescriptor) GetParent() Descriptor {
	return od.parent
}

// GetOwner returns the message to which this one-of field set belongs.
func (od *OneOfDescriptor) GetOwner() *MessageDescriptor {
	return od.parent
}

func (od *OneOfDescriptor) GetFile() *FileDescriptor {
	return od.file
}

func (od *OneOfDescriptor) String() string {
	return od.proto.String()
}

// GetChoices returns the fields that are part of the one-of field set. At most one of
// these fields may be set for a given message.
func (od *OneOfDescriptor) GetChoices() []*FieldDescriptor {
	return od.choices
}
