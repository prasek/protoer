package desc

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/prasek/protoer/proto"
	"github.com/prasek/protoer/proto/golang"
)

// createFileDescriptor instantiates a new file descriptor for the given descriptor proto.
// The file's direct dependencies must be provided. If the given dependencies do not include
// all of the file's dependencies or if the contents of the descriptors are internally
// inconsistent (e.g. contain unresolvable symbols) then an error is returned.
func CreateFileDescriptor(fileDescriptorProto proto.Message, deps ...*FileDescriptor) (*FileDescriptor, error) {
	fileDescriptorProto, err := golang.ToNativeDescriptor(fileDescriptorProto)
	if err != nil {
		return nil, err
	}
	fd, ok := fileDescriptorProto.(*dpb.FileDescriptorProto)
	if !ok {
		return nil, fmt.Errorf("Not native FileDescriptorProto %T", fd)
	}
	ret := &FileDescriptor{proto: fd, symbols: map[string]Descriptor{}, fieldIndex: map[string]map[int32]*FieldDescriptor{}}
	pkg := fd.GetPackage()

	// populate references to file descriptor dependencies
	files := map[string]*FileDescriptor{}
	for _, f := range deps {
		files[f.proto.GetName()] = f
	}
	ret.deps = make([]*FileDescriptor, len(fd.GetDependency()))
	for i, d := range fd.GetDependency() {
		ret.deps[i] = files[d]
		if ret.deps[i] == nil {
			return nil, fmt.Errorf("Given dependencies did not include %q", d)
		}
	}
	ret.publicDeps = make([]*FileDescriptor, len(fd.GetPublicDependency()))
	for i, pd := range fd.GetPublicDependency() {
		ret.publicDeps[i] = ret.deps[pd]
	}
	ret.weakDeps = make([]*FileDescriptor, len(fd.GetWeakDependency()))
	for i, wd := range fd.GetWeakDependency() {
		ret.weakDeps[i] = ret.deps[wd]
	}
	ret.isProto3 = fd.GetSyntax() == "proto3"

	// populate all tables of child descriptors
	for _, m := range fd.GetMessageType() {
		md, n := createMessageDescriptor(ret, ret, pkg, m, ret.symbols)
		ret.symbols[n] = md
		ret.messages = append(ret.messages, md)
	}
	for _, e := range fd.GetEnumType() {
		ed, n := createEnumDescriptor(ret, ret, pkg, e, ret.symbols)
		ret.symbols[n] = ed
		ret.enums = append(ret.enums, ed)
	}
	for _, ex := range fd.GetExtension() {
		exd, n := createFieldDescriptor(ret, ret, pkg, ex)
		ret.symbols[n] = exd
		ret.extensions = append(ret.extensions, exd)
	}
	for _, s := range fd.GetService() {
		sd, n := createServiceDescriptor(ret, pkg, s, ret.symbols)
		ret.symbols[n] = sd
		ret.services = append(ret.services, sd)
	}
	sourceCodeInfo := createSourceInfoMap(fd)

	// now we can resolve all type references and source code info
	scopes := []scope{fileScope(ret)}
	path := make([]int32, 1, 8)
	path[0] = File_messagesTag
	for i, md := range ret.messages {
		if err := md.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return nil, err
		}
	}
	path[0] = File_enumsTag
	for i, ed := range ret.enums {
		ed.resolve(append(path, int32(i)), sourceCodeInfo)
	}
	path[0] = File_extensionsTag
	for i, exd := range ret.extensions {
		if err := exd.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return nil, err
		}
	}
	path[0] = File_servicesTag
	for i, sd := range ret.services {
		if err := sd.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// createFileDescriptors constructs a set of descriptors, one for each of the
// given descriptor protos. The given set of descriptor protos must include all
// transitive dependencies for every file.
func createFileDescriptors(fds []*dpb.FileDescriptorProto) (map[string]*FileDescriptor, error) {
	if len(fds) == 0 {
		return nil, nil
	}
	files := map[string]*dpb.FileDescriptorProto{}
	resolved := map[string]*FileDescriptor{}
	var name string
	for _, fd := range fds {
		name = fd.GetName()
		files[name] = fd
	}
	for _, fd := range fds {
		_, err := createFromSet(fd.GetName(), nil, files, resolved)
		if err != nil {
			return nil, err
		}
	}
	return resolved, nil
}

// createFileDescriptorFromSet creates a descriptor from the given file descriptor set. The
// set's *last* file will be the returned descriptor. The set's remaining files must comprise
// the full set of transitive dependencies of that last file. This is the same format and
// order used by protoc when emitting a FileDescriptorSet file with an invocation like so:
//    protoc --descriptor_set_out=./test.protoset --include_imports -I. test.proto
func CreateFileDescriptorFromSet(fileDescriptorSet proto.Message) (*FileDescriptor, error) {
	fileDescriptorSet, err := golang.ToNativeDescriptor(fileDescriptorSet)
	if err != nil {
		return nil, err
	}
	fds, ok := fileDescriptorSet.(*dpb.FileDescriptorSet)
	if !ok {
		return nil, fmt.Errorf("Not native FileDescriptorSet %T", fds)
	}
	files := fds.GetFile()
	if len(files) == 0 {
		return nil, errors.New("file descriptor set is empty")
	}
	resolved, err := createFileDescriptors(files)
	if err != nil {
		return nil, err
	}
	lastFilename := files[len(files)-1].GetName()
	return resolved[lastFilename], nil
}

// createFromSet creates a descriptor for the given filename. It recursively
// creates descriptors for the given file's dependencies.
func createFromSet(filename string, seen []string, files map[string]*dpb.FileDescriptorProto, resolved map[string]*FileDescriptor) (*FileDescriptor, error) {
	for _, s := range seen {
		if filename == s {
			return nil, fmt.Errorf("cycle in imports: %s", strings.Join(append(seen, filename), " -> "))
		}
	}
	seen = append(seen, filename)

	if d, ok := resolved[filename]; ok {
		return d, nil
	}
	fdp := files[filename]
	if fdp == nil {
		return nil, fmt.Errorf("file descriptor set missing a dependency: %s", filename)
	}
	deps := make([]*FileDescriptor, len(fdp.GetDependency()))
	for i, depName := range fdp.GetDependency() {
		if dep, err := createFromSet(depName, seen, files, resolved); err != nil {
			return nil, err
		} else {
			deps[i] = dep
		}
	}
	d, err := CreateFileDescriptor(fdp, deps...)
	if err != nil {
		return nil, err
	}
	resolved[filename] = d
	return d, nil
}

func createServiceDescriptor(fd *FileDescriptor, enclosing string, sd *dpb.ServiceDescriptorProto, symbols map[string]Descriptor) (*ServiceDescriptor, string) {
	serviceName := merge(enclosing, sd.GetName())
	ret := &ServiceDescriptor{proto: sd, file: fd, fqn: serviceName}
	for _, m := range sd.GetMethod() {
		md, n := createMethodDescriptor(fd, ret, serviceName, m)
		symbols[n] = md
		ret.methods = append(ret.methods, md)
	}
	return ret, serviceName
}

func createMethodDescriptor(fd *FileDescriptor, parent *ServiceDescriptor, enclosing string, md *dpb.MethodDescriptorProto) (*MethodDescriptor, string) {
	// request and response types get resolved later
	methodName := merge(enclosing, md.GetName())
	return &MethodDescriptor{proto: md, parent: parent, file: fd, fqn: methodName}, methodName
}

func createMessageDescriptor(fd *FileDescriptor, parent Descriptor, enclosing string, md *dpb.DescriptorProto, symbols map[string]Descriptor) (*MessageDescriptor, string) {
	msgName := merge(enclosing, md.GetName())
	ret := &MessageDescriptor{proto: md, parent: parent, file: fd, fqn: msgName}
	for _, f := range md.GetField() {
		fld, n := createFieldDescriptor(fd, ret, msgName, f)
		symbols[n] = fld
		ret.fields = append(ret.fields, fld)
	}
	for _, nm := range md.NestedType {
		nmd, n := createMessageDescriptor(fd, ret, msgName, nm, symbols)
		symbols[n] = nmd
		ret.nested = append(ret.nested, nmd)
	}
	for _, e := range md.EnumType {
		ed, n := createEnumDescriptor(fd, ret, msgName, e, symbols)
		symbols[n] = ed
		ret.enums = append(ret.enums, ed)
	}
	for _, ex := range md.GetExtension() {
		exd, n := createFieldDescriptor(fd, ret, msgName, ex)
		symbols[n] = exd
		ret.extensions = append(ret.extensions, exd)
	}
	for i, o := range md.GetOneofDecl() {
		od, n := createOneOfDescriptor(fd, ret, i, msgName, o)
		symbols[n] = od
		ret.oneOfs = append(ret.oneOfs, od)
	}
	/*
		for _, r := range md.GetExtensionRange() {
			// proto.ExtensionRange is inclusive (and that's how extension ranges are defined in code).
			// but protoc converts range to exclusive end in descriptor, so we must convert back
			end := r.GetEnd() - 1
			ret.extRanges = append(ret.extRanges, proto.ExtensionRange{
				Start: r.GetStart(),
				End:   end})
		}
		sort.Sort(ret.extRanges)
	*/
	ret.isProto3 = fd.isProto3
	ret.isMapEntry = md.GetOptions().GetMapEntry() &&
		len(ret.fields) == 2 &&
		ret.fields[0].GetNumber() == 1 &&
		ret.fields[1].GetNumber() == 2

	return ret, msgName
}
func createFieldDescriptor(fd *FileDescriptor, parent Descriptor, enclosing string, fld *dpb.FieldDescriptorProto) (*FieldDescriptor, string) {
	fldName := merge(enclosing, fld.GetName())
	ret := &FieldDescriptor{proto: fld, parent: parent, file: fd, fqn: fldName}
	if fld.GetExtendee() == "" {
		ret.owner = parent.(*MessageDescriptor)
	}
	// owner for extensions, field type (be it message or enum), and one-ofs get resolved later
	return ret, fldName
}

func createEnumDescriptor(fd *FileDescriptor, parent Descriptor, enclosing string, ed *dpb.EnumDescriptorProto, symbols map[string]Descriptor) (*EnumDescriptor, string) {
	enumName := merge(enclosing, ed.GetName())
	ret := &EnumDescriptor{proto: ed, parent: parent, file: fd, fqn: enumName}
	for _, ev := range ed.GetValue() {
		evd, n := createEnumValueDescriptor(fd, ret, enumName, ev)
		symbols[n] = evd
		ret.values = append(ret.values, evd)
	}
	if len(ret.values) > 0 {
		ret.valuesByNum = make(sortedValues, len(ret.values))
		copy(ret.valuesByNum, ret.values)
		sort.Stable(ret.valuesByNum)
	}
	return ret, enumName
}

func createEnumValueDescriptor(fd *FileDescriptor, parent *EnumDescriptor, enclosing string, evd *dpb.EnumValueDescriptorProto) (*EnumValueDescriptor, string) {
	valName := merge(enclosing, evd.GetName())
	return &EnumValueDescriptor{proto: evd, parent: parent, file: fd, fqn: valName}, valName
}

func createOneOfDescriptor(fd *FileDescriptor, parent *MessageDescriptor, index int, enclosing string, od *dpb.OneofDescriptorProto) (*OneOfDescriptor, string) {
	oneOfName := merge(enclosing, od.GetName())
	ret := &OneOfDescriptor{proto: od, parent: parent, file: fd, fqn: oneOfName}
	for _, f := range parent.fields {
		oi := f.proto.OneofIndex
		if oi != nil && *oi == int32(index) {
			f.oneOf = ret
			ret.choices = append(ret.choices, f)
		}
	}
	return ret, oneOfName
}
