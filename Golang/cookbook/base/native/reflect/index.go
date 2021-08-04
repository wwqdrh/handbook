package reflect

import (
	"reflect"
	"unsafe"
)

func String2Slice(words string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&words))))
}
