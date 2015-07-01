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
