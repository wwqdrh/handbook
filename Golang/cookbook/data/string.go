package data

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString shows a number of methods
// for searching a string
func SearchString() {
	s := "this is a test"

	// returns true because s contains
	// the word this
	fmt.Println(strings.Contains(s, "this"))

	// returns true because s contains the letter a
	// would also match if it contained b or c
	fmt.Println(strings.ContainsAny(s, "abc"))

	// returns true because s starts with this
	fmt.Println(strings.HasPrefix(s, "this"))

	// retursn true because s ends with this
	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString modifies a string in a number of ways
func ModifyString() {
	s := "simple string"

	// prints [simple string]
	fmt.Println(strings.Split(s, " "))

	// prints "Simple String"
	fmt.Println(strings.Title(s))

	// prints "simple string"; all trailing and
	// leading white space is removed
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader demonstrates how to create
// an io.Reader interface quickly with a string
func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)

	// prints s on Stdout
	io.Copy(os.Stdout, r)
}

/**
 * 驼峰转蛇形 snake string
 * @description XxYy to xx_yy , XxYY to xx_y_y
 **/
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

/**
 * 蛇形转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY
 **/
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
