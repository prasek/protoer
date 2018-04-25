package desc

import "github.com/prasek/protoer/proto"

func GetExtension(pb proto.Message, field int32) (interface{}, error) {
	return proto.GetExtension(pb, field)
}

func GetBoolExtension(pb proto.Message, field int32, def bool) bool {
	v, err := proto.GetExtension(pb, field)
	if err != nil {
		return def
	}
	if v == nil {
		return def
	}
	return *v.(*bool)
}
