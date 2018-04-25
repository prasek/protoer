package desc

import (
	"fmt"
	"sort"
)

var _ Descriptor = (*EnumDescriptor)(nil)

type sortedValues []*EnumValueDescriptor

func (sv sortedValues) Len() int {
	return len(sv)
}

func (sv sortedValues) Less(i, j int) bool {
	return sv[i].GetNumber() < sv[j].GetNumber()
}

func (sv sortedValues) Swap(i, j int) {
	sv[i], sv[j] = sv[j], sv[i]
}

func (ed *EnumDescriptor) GetName() string {
	return ed.proto.GetName()
}

func (ed *EnumDescriptor) GetFullyQualifiedName() string {
	return ed.fqn
}

func (ed *EnumDescriptor) GetParent() Descriptor {
	return ed.parent
}

func (ed *EnumDescriptor) GetFile() *FileDescriptor {
	return ed.file
}

func (ed *EnumDescriptor) String() string {
	return ed.proto.String()
}

// GetValues returns all of the allowed values defined for this enum.
func (ed *EnumDescriptor) GetValues() []*EnumValueDescriptor {
	return ed.values
}

// FindValueByName finds the enum value with the given name. If no such value exists
// then nil is returned.
func (ed *EnumDescriptor) FindValueByName(name string) *EnumValueDescriptor {
	fqn := fmt.Sprintf("%s.%s", ed.fqn, name)
	if vd, ok := ed.file.symbols[fqn].(*EnumValueDescriptor); ok {
		return vd
	} else {
		return nil
	}
}

// FindValueByNumber finds the value with the given numeric value. If no such value
// exists then nil is returned. If aliases are allowed and multiple values have the
// given number, the first declared value is returned.
func (ed *EnumDescriptor) FindValueByNumber(num int32) *EnumValueDescriptor {
	index := sort.Search(len(ed.valuesByNum), func(i int) bool { return ed.valuesByNum[i].GetNumber() >= num })
	if index < len(ed.valuesByNum) {
		vd := ed.valuesByNum[index]
		if vd.GetNumber() == num {
			return vd
		}
	}
	return nil
}
