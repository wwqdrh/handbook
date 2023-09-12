package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewStruct(t *testing.T) {
	fields := []*Field{
		{
			Type: reflect.StructField{
				Name: "ID",
				Type: reflect.TypeOf(""),
			},
			Value: reflect.ValueOf("1"),
		},
		{
			Type: reflect.StructField{
				Name: "User",
				Type: reflect.TypeOf(""),
			},
			Value: reflect.ValueOf("user"),
		},
	}
	val := NewStruct(fields)
	fmt.Println(val)
}

func TestBytes2String(t *testing.T) {
	a := bytes2string([]byte("hello"))
	fmt.Println(a)
	fmt.Println(len(a))
}
