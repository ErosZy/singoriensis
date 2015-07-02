package common

import (
	"reflect"
)

func CallObjMethod(objs interface{}, name string, params []interface{}) {
	in := make([]reflect.Value, 0)

	for _, v := range params {
		in = append(in, reflect.ValueOf(v))
	}

	tmp := reflect.ValueOf(objs)
	length := tmp.Len()

	for i := 0; i < length; i++ {
		item := tmp.Index(i)
		method := item.MethodByName(name)
		if method.IsValid() {
			method.Call(in)
		}
	}
}

func NewDjb2Hash(str string) int64 {
	var hash int64 = 0

	for _, v := range str {
		strInt := int64(v)
		hash = hash*33 + strInt

	}

	return hash
}
