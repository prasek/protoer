package desc

func (vd *EnumValueDescriptor) GetName() string {
	return vd.proto.GetName()
}

// GetNumber returns the numeric value associated with this enum value.
func (vd *EnumValueDescriptor) GetNumber() int32 {
	return vd.proto.GetNumber()
}

func (vd *EnumValueDescriptor) GetFullyQualifiedName() string {
	return vd.fqn
}

func (vd *EnumValueDescriptor) GetParent() Descriptor {
	return vd.parent
}

// GetEnum returns the enum in which this enum value is defined.
func (vd *EnumValueDescriptor) GetEnum() *EnumDescriptor {
	return vd.parent
}

func (vd *EnumValueDescriptor) GetFile() *FileDescriptor {
	return vd.file
}

func (vd *EnumValueDescriptor) String() string {
	return vd.proto.String()
}
