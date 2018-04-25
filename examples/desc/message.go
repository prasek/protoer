package desc

import (
	"fmt"
)

var _ Descriptor = (*MessageDescriptor)(nil)

func (md *MessageDescriptor) GetName() string {
	return md.proto.GetName()
}

func (md *MessageDescriptor) GetFullyQualifiedName() string {
	return md.fqn
}

func (md *MessageDescriptor) GetParent() Descriptor {
	return md.parent
}

func (md *MessageDescriptor) GetFile() *FileDescriptor {
	return md.file
}

func (md *MessageDescriptor) String() string {
	return md.proto.String()
}

// IsMapEntry returns true if this is a synthetic message type that represents an entry
// in a map field.
func (md *MessageDescriptor) IsMapEntry() bool {
	return md.isMapEntry
}

// GetFields returns all of the fields for this message.
func (md *MessageDescriptor) GetFields() []*FieldDescriptor {
	return md.fields
}

// GetNestedMessageTypes returns all of the message types declared inside this message.
func (md *MessageDescriptor) GetNestedMessageTypes() []*MessageDescriptor {
	return md.nested
}

// GetNestedEnumTypes returns all of the enums declared inside this message.
func (md *MessageDescriptor) GetNestedEnumTypes() []*EnumDescriptor {
	return md.enums
}

// GetNestedExtensions returns all of the extensions declared inside this message.
func (md *MessageDescriptor) GetNestedExtensions() []*FieldDescriptor {
	return md.extensions
}

// GetOneOfs returns all of the one-of field sets declared inside this message.
func (md *MessageDescriptor) GetOneOfs() []*OneOfDescriptor {
	return md.oneOfs
}

// IsProto3 returns true if the file in which this message is defined declares a syntax of "proto3".
func (md *MessageDescriptor) IsProto3() bool {
	return md.isProto3
}

/*
// GetExtensionRanges returns the ranges of extension field numbers for this message.
func (md *MessageDescriptor) GetExtensionRanges() []proto.ExtensionRange {
	return md.extRanges
}

// IsExtendable returns true if this message has any extension ranges.
func (md *MessageDescriptor) IsExtendable() bool {
	return len(md.extRanges) > 0
}

// IsExtension returns true if the given tag number is within any of this message's
// extension ranges.
func (md *MessageDescriptor) IsExtension(tagNumber int32) bool {
	return md.extRanges.IsExtension(tagNumber)
}

type extRanges []proto.ExtensionRange

func (er extRanges) String() string {
	var buf bytes.Buffer
	first := true
	for _, r := range er {
		if first {
			first = false
		} else {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d..%d", r.Start, r.End)
	}
	return buf.String()
}

func (er extRanges) IsExtension(tagNumber int32) bool {
	i := sort.Search(len(er), func(i int) bool { return er[i].End >= tagNumber })
	return i < len(er) && tagNumber >= er[i].Start
}

func (er extRanges) Len() int {
	return len(er)
}

func (er extRanges) Less(i, j int) bool {
	return er[i].Start < er[j].Start
}

func (er extRanges) Swap(i, j int) {
	er[i], er[j] = er[j], er[i]
}
*/

// FindFieldByName finds the field with the given name. If no such field exists
// then nil is returned. Only regular fields are returned, not extensions.
func (md *MessageDescriptor) FindFieldByName(fieldName string) *FieldDescriptor {
	fqn := fmt.Sprintf("%s.%s", md.fqn, fieldName)
	if fd, ok := md.file.symbols[fqn].(*FieldDescriptor); ok && !fd.IsExtension() {
		return fd
	} else {
		return nil
	}
}

// FindFieldByNumber finds the field with the given tag number. If no such field
// exists then nil is returned. Only regular fields are returned, not extensions.
func (md *MessageDescriptor) FindFieldByNumber(tagNumber int32) *FieldDescriptor {
	if fd, ok := md.file.fieldIndex[md.fqn][tagNumber]; ok && !fd.IsExtension() {
		return fd
	} else {
		return nil
	}
}

func messageScope(md *MessageDescriptor) scope {
	return func(name string) Descriptor {
		n := merge(md.fqn, name)
		if d, ok := md.file.symbols[n]; ok {
			return d
		}
		return nil
	}
}
