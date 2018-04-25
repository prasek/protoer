package desc

import (
	"fmt"

	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type sourceInfoMap map[string]*dpb.SourceCodeInfo_Location

func (m sourceInfoMap) Get(path []int32) *dpb.SourceCodeInfo_Location {
	return m[asMapKey(path)]
}

func (m sourceInfoMap) Put(path []int32, loc *dpb.SourceCodeInfo_Location) {
	m[asMapKey(path)] = loc
}

func (m sourceInfoMap) PutIfAbsent(path []int32, loc *dpb.SourceCodeInfo_Location) bool {
	k := asMapKey(path)
	if _, ok := m[k]; ok {
		return false
	}
	m[k] = loc
	return true
}

func asMapKey(slice []int32) string {
	// NB: arrays should be usable as map keys, but this does not
	// work due to a bug: https://github.com/golang/go/issues/22605
	//rv := reflect.ValueOf(slice)
	//arrayType := reflect.ArrayOf(rv.Len(), rv.Type().Elem())
	//array := reflect.New(arrayType).Elem()
	//reflect.Copy(array, rv)
	//return array.Interface()

	b := make([]byte, len(slice)*4)
	for i, s := range slice {
		j := i * 4
		b[j] = byte(s)
		b[j+1] = byte(s >> 8)
		b[j+2] = byte(s >> 16)
		b[j+3] = byte(s >> 24)
	}
	return string(b)
}

func createSourceInfoMap(fd *dpb.FileDescriptorProto) sourceInfoMap {
	res := sourceInfoMap{}
	for _, l := range fd.GetSourceCodeInfo().GetLocation() {
		res.Put(l.Path, l)
	}
	return res
}

func (fd *FileDescriptor) registerField(field *FieldDescriptor) {
	fields := fd.fieldIndex[field.owner.GetFullyQualifiedName()]
	if fields == nil {
		fields = map[int32]*FieldDescriptor{}
		fd.fieldIndex[field.owner.GetFullyQualifiedName()] = fields
	}
	fields[field.GetNumber()] = field
}

func (sd *ServiceDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap, scopes []scope) error {
	sd.sourceInfo = sourceCodeInfo.Get(path)
	path = append(path, Service_methodsTag)
	for i, md := range sd.methods {
		if err := md.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return err
		}
	}
	return nil
}

func (md *MethodDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap, scopes []scope) error {
	md.sourceInfo = sourceCodeInfo.Get(path)
	if desc, err := resolve(md.file, md.proto.GetInputType(), scopes); err != nil {
		return err
	} else {
		md.inType = desc.(*MessageDescriptor)
	}
	if desc, err := resolve(md.file, md.proto.GetOutputType(), scopes); err != nil {
		return err
	} else {
		md.outType = desc.(*MessageDescriptor)
	}
	return nil
}

func (md *MessageDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap, scopes []scope) error {
	md.sourceInfo = sourceCodeInfo.Get(path)
	path = append(path, Message_nestedMessagesTag)
	scopes = append(scopes, messageScope(md))
	for i, nmd := range md.nested {
		if err := nmd.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return err
		}
	}
	path[len(path)-1] = Message_enumsTag
	for i, ed := range md.enums {
		ed.resolve(append(path, int32(i)), sourceCodeInfo)
	}
	path[len(path)-1] = Message_fieldsTag
	for i, fld := range md.fields {
		if err := fld.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return err
		}
	}
	path[len(path)-1] = Message_extensionsTag
	for i, exd := range md.extensions {
		if err := exd.resolve(append(path, int32(i)), sourceCodeInfo, scopes); err != nil {
			return err
		}
	}
	path[len(path)-1] = Message_oneOfsTag
	for i, od := range md.oneOfs {
		od.resolve(append(path, int32(i)), sourceCodeInfo)
	}
	return nil
}

func (fd *FieldDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap, scopes []scope) error {
	if fd.proto.OneofIndex != nil && fd.oneOf == nil {
		return fmt.Errorf("Could not link field %s to one-of index %d", fd.fqn, *fd.proto.OneofIndex)
	}
	fd.sourceInfo = sourceCodeInfo.Get(path)
	if fd.proto.GetType() == dpb.FieldDescriptorProto_TYPE_ENUM {
		if desc, err := resolve(fd.file, fd.proto.GetTypeName(), scopes); err != nil {
			return err
		} else {
			fd.enumType = desc.(*EnumDescriptor)
		}
	}
	if fd.proto.GetType() == dpb.FieldDescriptorProto_TYPE_MESSAGE || fd.proto.GetType() == dpb.FieldDescriptorProto_TYPE_GROUP {
		if desc, err := resolve(fd.file, fd.proto.GetTypeName(), scopes); err != nil {
			return err
		} else {
			fd.msgType = desc.(*MessageDescriptor)
		}
	}
	if fd.proto.GetExtendee() != "" {
		if desc, err := resolve(fd.file, fd.proto.GetExtendee(), scopes); err != nil {
			return err
		} else {
			fd.owner = desc.(*MessageDescriptor)
		}
	}
	fd.file.registerField(fd)
	fd.isMap = fd.proto.GetLabel() == dpb.FieldDescriptorProto_LABEL_REPEATED &&
		fd.proto.GetType() == dpb.FieldDescriptorProto_TYPE_MESSAGE &&
		fd.GetMessageType().IsMapEntry()
	return nil
}

func (ed *EnumDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap) {
	ed.sourceInfo = sourceCodeInfo.Get(path)
	path = append(path, Enum_valuesTag)
	for i, evd := range ed.values {
		evd.resolve(append(path, int32(i)), sourceCodeInfo)
	}
}

func (vd *EnumValueDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap) {
	vd.sourceInfo = sourceCodeInfo.Get(path)
}
func (od *OneOfDescriptor) resolve(path []int32, sourceCodeInfo sourceInfoMap) {
	od.sourceInfo = sourceCodeInfo.Get(path)
}
