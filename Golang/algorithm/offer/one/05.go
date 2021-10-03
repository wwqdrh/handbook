package one

import "bytes"

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

input: "We are happy." output: "We%20are%20happy."
*/

func Hand5(s string) string {
	res := new(bytes.Buffer)
	for _, ch := range s {
		if ch == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteRune(ch)
		}
	}
	return res.String()
}
