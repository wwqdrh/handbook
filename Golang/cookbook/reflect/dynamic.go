package reflect

import (
	"reflect"
	"unsafe"
)

type Field struct {
	Type  reflect.StructField
	Value reflect.Value
}

func NewStruct(fields []*Field) interface{} {
	typFields := []reflect.StructField{}
	for _, item := range fields {
		typFields = append(typFields, item.Type)
	}

	typ := reflect.StructOf(typFields)
	val := reflect.New(typ)
	elm := val.Elem()
	for i, item := range fields {
		elm.Field(i).Set(item.Value)
	}
	return elm.Addr().Interface()
}

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
