package io

import "bytes"

func defangIPaddr(address string) string {
	buf := bytes.Buffer{}
	for _, ch := range address {
		if ch == '.' {
			buf.Write([]byte("[.]"))
		} else {
			buf.WriteByte(byte(ch))
		}
	}
	return buf.String()
}
