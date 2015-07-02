package common

import (
	"math/big"
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

func NewDjb2Hash(str string) *big.Int {
	hash := big.NewInt(5381)
	mulNum := big.NewInt(33)

	for _, v := range str {
		strInt := big.NewInt(int64(v))
		hash = hash.Mul(hash, mulNum)
		hash = hash.Add(hash, strInt)
	}

	return hash
}
