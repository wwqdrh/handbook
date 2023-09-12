package sequence

import (
	"sort"
	"strings"
)

////////////////////
// 字符串hash算法
////////////////////

func StringHash() {}

////////////////////
// 字符串匹配
////////////////////

func StringSunday() {}

// 字符串最长公共前缀
// 排序 然后将第一个和最后一个进行比较
func StringCommonPrefix(words ...string) int {
	if len(words) <= 0 {
		return -1
	}
	sort.Slice(words, func(i, j int) bool { return words[i] < words[j] })

	// ！不能比较byte
	str1, str2 := []rune(words[0]), []rune(words[len(words)-1])
	idx := 0

	for i, j := 0, 0; i < len(str1) && j < len(str2); i, j = i+1, j+1 {
		if str1[i] != str2[j] {
			break
		}
		idx++
	}
	return idx
}

////////////////////
// 回文串相关
////////////////////

func StringManacher() {}

// 查找字符串中最长的回文子串
// 以某个元素为中心 分别计算奇偶数长度的最大长度回文长度
func StringMaxSubStrPalindrome(words []rune) []rune {
	if len(words) < 2 {
		return words
	}

	var index, length int // 存储

	// 计算最大的长度并记录index
	palindrome := func(l, r int) {
		for (l >= 0 && r < len(words)) && words[l] == words[r] {
			l--
			r++
		}
		if length < r-l-1 {
			index = l + 1
			length = r - l - 1
		}
	}

	for i := 0; i < len(words)-1; i++ {
		palindrome(i, i)
		palindrome(i, i+1)
	}

	return words[index : index+length]
}

// 查找字符串中最长的回文子序列(不一定是连续的)
func StringMaxSubSeqPalindrome() {}

// 构造最长回文串
// 1、字符出现次数为双数的组合
// 2、字符出现次数为双数的组合 + 只出现一次的字符
// 遍历统计 如果遇到双数就将count+1 剩下落单的一个也可以加上
func StringMaxPalindromeBuild(words []rune) int {
	hashSet := make(map[rune]bool)

	count := 0
	for _, word := range words {
		if _, ok := hashSet[word]; ok {
			delete(hashSet, word)
			count++
		} else {
			hashSet[word] = true
		}
	}
	if len(hashSet) > 0 {
		return count*2 + 1
	} else {
		return count * 2
	}
}

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
