package desc

import (
	"fmt"
	"strings"
)

var _ Descriptor = (*FileDescriptor)(nil)

func (fd *FileDescriptor) GetName() string {
	return fd.proto.GetName()
}

func (fd *FileDescriptor) GetFullyQualifiedName() string {
	return fd.proto.GetName()
}

func (fd *FileDescriptor) GetPackage() string {
	return fd.proto.GetPackage()
}

func (fd *FileDescriptor) GetParent() Descriptor {
	return nil
}

func (fd *FileDescriptor) GetFile() *FileDescriptor {
	return fd
}

func (fd *FileDescriptor) String() string {
	return fd.proto.String()
}

// IsProto3 returns true if the file declares a syntax of "proto3".
func (fd *FileDescriptor) IsProto3() bool {
	return fd.isProto3
}

// GetDependencies returns all of this file's dependencies. These correspond to
// import statements in the file.
func (fd *FileDescriptor) GetDependencies() []*FileDescriptor {
	return fd.deps
}

// GetPublicDependencies returns all of this file's public dependencies. These
// correspond to public import statements in the file.
func (fd *FileDescriptor) GetPublicDependencies() []*FileDescriptor {
	return fd.publicDeps
}

// GetWeakDependencies returns all of this file's weak dependencies. These
// correspond to weak import statements in the file.
func (fd *FileDescriptor) GetWeakDependencies() []*FileDescriptor {
	return fd.weakDeps
}

// GetMessageTypes returns all top-level messages declared in this file.
func (fd *FileDescriptor) GetMessageTypes() []*MessageDescriptor {
	return fd.messages
}

// GetEnumTypes returns all top-level enums declared in this file.
func (fd *FileDescriptor) GetEnumTypes() []*EnumDescriptor {
	return fd.enums
}

// GetExtensions returns all top-level extensions declared in this file.
func (fd *FileDescriptor) GetExtensions() []*FieldDescriptor {
	return fd.extensions
}

// GetServices returns all services declared in this file.
func (fd *FileDescriptor) GetServices() []*ServiceDescriptor {
	return fd.services
}

// FindSymbol returns the descriptor contained within this file for the
// element with the given fully-qualified symbol name. If no such element
// exists then this method returns nil.
func (fd *FileDescriptor) FindSymbol(symbol string) Descriptor {
	return fd.symbols[symbol]
}

// FindMessage finds the message with the given fully-qualified name. If no
// such element exists in this file then nil is returned.
func (fd *FileDescriptor) FindMessage(msgName string) *MessageDescriptor {
	if md, ok := fd.symbols[msgName].(*MessageDescriptor); ok {
		return md
	} else {
		return nil
	}
}

// FindEnum finds the enum with the given fully-qualified name. If no such
// element exists in this file then nil is returned.
func (fd *FileDescriptor) FindEnum(enumName string) *EnumDescriptor {
	if ed, ok := fd.symbols[enumName].(*EnumDescriptor); ok {
		return ed
	} else {
		return nil
	}
}

// FindService finds the service with the given fully-qualified name. If no
// such element exists in this file then nil is returned.
func (fd *FileDescriptor) FindService(serviceName string) *ServiceDescriptor {
	if sd, ok := fd.symbols[serviceName].(*ServiceDescriptor); ok {
		return sd
	} else {
		return nil
	}
}

// FindExtension finds the extension field for the given extended type name and
// tag number. If no such element exists in this file then nil is returned.
func (fd *FileDescriptor) FindExtension(extendeeName string, tagNumber int32) *FieldDescriptor {
	if exd, ok := fd.fieldIndex[extendeeName][tagNumber]; ok && exd.IsExtension() {
		return exd
	} else {
		return nil
	}
}

// FindExtensionByName finds the extension field with the given fully-qualified
// name. If no such element exists in this file then nil is returned.
func (fd *FileDescriptor) FindExtensionByName(extName string) *FieldDescriptor {
	if exd, ok := fd.symbols[extName].(*FieldDescriptor); ok && exd.IsExtension() {
		return exd
	} else {
		return nil
	}
}

func fileScope(fd *FileDescriptor) scope {
	// we search symbols in this file, but also symbols in other files that have
	// the same package as this file or a "parent" package (in protobuf,
	// packages are a hierarchy like C++ namespaces)
	prefixes := createPrefixList(fd.proto.GetPackage())
	return func(name string) Descriptor {
		for _, prefix := range prefixes {
			n := merge(prefix, name)
			d := findSymbol(fd, n, false)
			if d != nil {
				return d
			}
		}
		return nil
	}
}

func resolve(fd *FileDescriptor, name string, scopes []scope) (Descriptor, error) {
	if strings.HasPrefix(name, ".") {
		// already fully-qualified
		d := findSymbol(fd, name[1:], false)
		if d != nil {
			return d, nil
		}
	} else {
		// unqualified, so we look in the enclosing (last) scope first and move
		// towards outermost (first) scope, trying to resolve the symbol
		for i := len(scopes) - 1; i >= 0; i-- {
			d := scopes[i](name)
			if d != nil {
				return d, nil
			}
		}
	}
	return nil, fmt.Errorf("File %q included an unresolvable reference to %q", fd.proto.GetName(), name)
}

func findSymbol(fd *FileDescriptor, name string, public bool) Descriptor {
	d := fd.symbols[name]
	if d != nil {
		return d
	}

	// When public = false, we are searching only directly imported symbols. But we
	// also need to search transitive public imports due to semantics of public imports.
	var deps []*FileDescriptor
	if public {
		deps = fd.publicDeps
	} else {
		deps = fd.deps
	}
	for _, dep := range deps {
		d = findSymbol(dep, name, true)
		if d != nil {
			return d
		}
	}

	return nil
}
