package desc

import "github.com/prasek/protoer/proto"

func GetExtension(pb proto.Message, ext interface{}) (interface{}, error) {
	return proto.GetExtension(pb, ext)
}

func GetBoolExtension(pb proto.Message, ext interface{}, def bool) bool {
	v, err := proto.GetExtension(pb, ext)
	if err != nil {
		return def
	}
	if v == nil {
		return def
	}
	return *v.(*bool)
}
