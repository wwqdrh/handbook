package reflect

import "unsafe"

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
