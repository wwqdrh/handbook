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

// KMP 算法，它是用来解决字符串查找的问题，可以在一个字符串 (S)中查找一个子串(W)出现的位置。
// KMP 算法把字符匹配的时间复杂度缩小到 O(m+n) ,而空间复 杂度也只有O(m)。
// 因为“暴力搜索”的方法会反复回溯主串，导致效率低下，而KMP算法可以利用已经部 分匹配这个有效信息，
// 保持主串上的指针不回溯，通过修改子串的指针，让模式串尽量地移动到有效的 位置。
// 需要创建模式串的next，他表示当两个字符串进行模式匹配失败的时候，需要从模式串的哪一个位置重新开始匹配
// 我们利用这next数组，来减少字符比较的次数，当比较字符失败时，就根据next换到新的位置，这个位置就是next的值
// 每次遍历 失败的部分查看前后缀是否相同 相同的话下一个匹配位置就是相同的下一个
// [参考视频](https://www.bilibili.com/video/av3246487/?from=search&seid=17173603269940723925)
func StringKMP(mainStr, subStr string) int {
	if len(mainStr) < len(subStr) {
		return -1
	}
	// 构建next数组
	nextArr := make([]int, len(subStr))
	for i := 0; i < len(subStr); i++ {
		t := -1
		for j := 0; j < i-1; j++ {
			if subStr[:j+1] == subStr[i-(j+1):i] {
				t = j + 1 // 有相同的前后缀则将其移动到j后面一个
			}
		}
		nextArr[i] = t
	}

	// kmp
	i, y, index := 0, 0, -1
	// i+y: mainStr开始匹配的位置 y: subStr开始匹配的位置 index: 结果
	for i <= len(mainStr)-len(subStr) {
		for j := y; j <= len(subStr); j++ {
			// 匹配到一个
			if j == len(subStr) {
				index = i
				i++
				y = 0
				break
			}

			// 相等继续
			if mainStr[i+j] == subStr[j] {
				continue
			}

			// 这个位置匹配失败 根据nextArr寻找
			idx := j
			for ; idx > 1; idx-- {
				if nextArr[idx] >= 1 {
					i = i + (j - nextArr[j]) // 移动i, y
					y = nextArr[j]
					break
				}
			}
			if idx <= 1 {
				i++
				y = 0
			}
			break
		}
	}
	return index
}

// BM算法也是一种精确字符串匹配算法，它采用从右向左比的方法，
// 同时应用到了两种启发式规 则，即坏字符规则 和好后缀规则 ，来决定向右跳跃的距离。
// 基本思路就是从右往左进行字符匹 配，遇到不匹配的字符后从坏字符表和好后缀表找一个最大的右移值，将模式串右移继续匹配。
// BM算法 对后缀蛮力匹配算法的改进
func StringBM() {}

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
