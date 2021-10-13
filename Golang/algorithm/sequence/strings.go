package sequence

import "strings"

////////////////////
// 字符串hash算法
////////////////////

func StringHash() {}

////////////////////
// 字符串匹配
////////////////////

func StringSunday() {}

func StringKMP() {}

////////////////////
// 回文串相关
////////////////////

func StringManacher() {}

////////////////////
// 与字符串相关
////////////////////

// 判断字符串中是否全都不同
// 实现一个算法 确定一个字符串的所有字符是否全都不同并且不能使用额外存储结构
// 如果能够使用额外结构的话那就是存储下是否访问过 或者将其转为set之类判断长度是否有变化
func IsUniqueString(s string) bool {
	result := true

	// 方法1
	for _, v := range s {
		if strings.Count(s, string(v)) > 1 {
			result = result && false
		}
	}
	result = result && true

	// 方法2
	for k, v := range s {
		if strings.Index(s, string(v)) != k {
			result = result && false
		}
	}
	result = result && true

	return result
}

// 将字符串翻转 不能使用额外的空间
// 考察string其实是字节切片，因此使用[]rune转换是不需要额外空间的
func ReverseString(words string) string {
	str := []rune(words)
	l := len(str)
	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str)
}
