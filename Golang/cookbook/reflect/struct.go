package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

// isinstance struct
func IsStruct(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Struct
}

// get struct name
func GetStructName(v interface{}) string {
	if !IsStruct(v) {
		return ""
	}
	return reflect.TypeOf(v).Name()
}

// iter struct field
func IterStructField(v interface{}) error {
	if !IsStruct(v) {
		return errors.New("不是结构体")
	}
	getType := reflect.TypeOf(v)
	getValue := reflect.ValueOf(v)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s %s: %v = %v\n", field.Name, field.Tag, field.Type, value)
	}
	return nil
}
