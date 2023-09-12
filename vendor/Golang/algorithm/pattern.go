package algorithm

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

////////////////////
// 各类匹配算法
////////////////////

// Rabin-Karp匹配
// 利用hash函数快速判断两个是否能组成成功
// 每次我们其实不用重新计算整个字符串的hash而是直接原hash值乘以R加上s[k-1]并且减去s[i]R^(k-1)
// 暴力匹配O(mn)是时间复杂度， Rabin-Karp 的时间复杂度在O(m+n)， 最坏的情况每次hash相同字符串不相同时间复杂度会变成O(mn)但是这种情况比较罕见

const primeRK = 16777619 // 32位FNV hash算法中的基础质数（相当于进制）

// hashStr 散列函数 返回字符串以及primeRK的k-1 ^ (len(sep)-1)
func hashStr(sep string) (uint32, uint32) {
	// 字符串hash
	var hash uint32 = 0
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i]) // 循环得到字符串hash
	}

	// 获取primeRK 的 len(sep)-1 次方 快速幂快速求平方
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func indexRabinKarp(s, substr string) int {
	hashss, pow := hashStr(substr) // 待匹配的hash值 以及primeRK^对应长度
	n := len(substr)
	var h uint32

	// 计算目标字符串前n位hash并与匹配字符串hash进行对比 hash相同就返回下标
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}

	// 如果前面不相同就开始不断匹配后面的
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}

// Rabin-Karp匹配
func RabinKarp(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	var k1 int = 1000000000 + 7
	var k2 int = 1337
	rand.Seed(time.Now().Unix())
	var kMod1 int64 = int64(rand.Intn(k1)) + int64(k1)
	var kMod2 int64 = int64(rand.Intn(k2)) + int64(k2)

	var hash_needle int64 = 0
	for i := 0; i < m; i++ {
		hash_needle = (hash_needle*kMod2 + int64(needle[i])) % kMod1
	}
	var hash_haystack int64 = 0
	var extra int64 = 1
	for i := 0; i < m-1; i++ {
		hash_haystack = (hash_haystack*kMod2 + int64(haystack[i%n])) % kMod1
		extra = (extra * kMod2) % kMod1
	}
	for i := m - 1; (i - m + 1) < n; i++ {
		hash_haystack = (hash_haystack*kMod2 + int64(haystack[i%n])) % kMod1
		if hash_haystack == hash_needle {
			return i - m + 1
		}
		hash_haystack = (hash_haystack - extra*int64(haystack[(i-m+1)%n])) % kMod1
		hash_haystack = (hash_haystack + kMod1) % kMod1
	}
	return -1
}

// sunday匹配
// 在匹配失败时关注文本串中参加匹配的最末位字符的下一位字符
// 暴力匹配匹配串在失败的时候每次只能移动一位
// sunday则是根据haystack参与的末尾下一位判断在不在匹配串中，在的话移动指定位置对齐然后进行判断
// 如果不在的话直接跳过移动整个匹配串的长度
func SundaySearch(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 && len(needle) != 0 {
		return -1
	}

	keys := make([]int, 256) // 记录ascii对应字符在字符串最右边的位置
	for i := 0; i < 256; i++ {
		keys[i] = -1
	}
	for i := 0; i < len(needle); i++ {
		keys[needle[i]] = i // 记录每个字符在字符串哪个位置，会被后面的覆盖 记录最右侧位置
	}

	//
	sIndex, pIndex, space := 0, 0, 0 // 双指针
	for sIndex < len(haystack) {
		if haystack[sIndex] == needle[pIndex] {
			sIndex++
			pIndex++
			if pIndex == len(needle) {
				return sIndex - len(needle)
			}
		} else {
			pIndex = 0
			if space+len(needle) < len(haystack) {
				space += len(needle) - keys[haystack[space+len(needle)]] // !寻找下个模式的匹配初始位置
			} else {
				return -1
			}
			sIndex = space
		}
	}
	return -1
}

func randInt(a, b int) int {
	return a + rand.Intn(b-a)
}

// 获取一个字符串中最长的重复子串 字符串hash+滑动窗口
// 模一般选取编码的信息量的平方的数量级。而取模则会带来哈希碰撞。本题中为了避免碰撞，我们使用双哈希，即用两套进制和模的组合，来对字符串进行编码。只有两种编码都相同时，我们才认为字符串相同。
// 但是力扣上的题还是无法通过 需要使用二分进行优化 就是说如果n长度是循环的那么1...n也是循环的
// l, r := 1, len(s)-1
//     length, start := 0, -1
//     for l <= r {
//         m := l + (r-l+1)/2
//         idx := check(arr, m, a1, a2, mod1, mod2)
//         if idx != -1 { // 有重复子串，移动左边界
//             l = m + 1
//             length = m
//             start = idx
//         } else { // 无重复子串，移动右边界
//             r = m - 1
//         }
//     }
func LongestSubStr(words string) string {
	length := len(words)
	if length <= 1 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	// 生成两个进制
	a1, a2 := randInt(26, 100), randInt(26, 100)
	// 生成两个模
	mod1, mod2 := randInt(1e9+7, math.MaxInt32), randInt(1e9+7, math.MaxInt32)
	// 先对所有字符进行编码
	arr := []byte(words)
	for i := range arr {
		arr[i] -= 'a'
	}

	l, r := 1, len(words)-1
	length, start := 0, -1
	for l <= r {
		m := l + (r-l+1)/2
		idx := check(arr, m, a1, a2, mod1, mod2)
		if idx != -1 { // 有重复子串，移动左边界
			l = m + 1
			length = m
			start = idx
		} else { // 无重复子串，移动右边界
			r = m - 1
		}
	}
	if start == -1 {
		return ""
	}
	return words[start : start+length]
}

// 返回起点或者-1
func check(arr []byte, curLen, a1, a2, mod1, mod2 int) int {
	length := len(arr)
	hashSet1, hashSet2 := map[int]bool{}, map[int]bool{} // 存储字符串hash值
	cur1, cur2 := 0, 0
	pow1, pow2 := 1, 1
	for i := 0; i < curLen; i++ {
		cur1 = (cur1*a1 + int(arr[i])) % mod1
		cur2 = (cur2*a2 + int(arr[i])) % mod2
		pow1 = (pow1 * a1) % mod1
		pow2 = (pow2 * a2) % mod2
	}
	hashSet1[cur1] = true
	hashSet2[cur2] = true
	for i := 1; i <= length-curLen; i++ {
		cur1 = (cur1*a1%mod1 - int(arr[i-1])*pow1%mod1) % mod1
		cur1 = (cur1 + int(arr[i+curLen-1])) % mod1
		cur2 = (cur2*a2%mod2 - int(arr[i-1])*pow2%mod2) % mod2
		cur2 = (cur2 + int(arr[i+curLen-1])) % mod2
		// 还可能变成负数 因为int64都可能超
		if cur1 < 0 {
			cur1 += mod1
		}
		if cur2 < 0 {
			cur2 += mod2
		}
		if hashSet1[cur1] && hashSet2[cur2] {
			return i
		}
		hashSet1[cur1] = true
		hashSet2[cur2] = true
	}
	return -1
}

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
					i = i + (idx - nextArr[idx]) // 移动i, y
					y = nextArr[idx]
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

func stringkmp2(pattern, word string) int {
	m, n := len(pattern), len(word)

	// 指定往前跳转
	pi := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		if j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j += 1
		}
		pi[i] = j
	}

	// kmp
	i, y := 0, 0
	// i+y: mainStr开始匹配的位置 y: subStr开始匹配的位置 index: 结果
	for i <= n-m {
		for j := y; j <= m; j++ {
			// 匹配到一个
			if j == m {
				return i
			}

			// 相等继续
			if word[i+j] == pattern[j] {
				continue
			}

			// 这个位置匹配失败 根据nextArr寻找
			idx := j
			for ; idx > 1; idx-- {
				if pi[idx] >= 1 {
					i = i + (idx - pi[idx] + 1) // 移动i, y
					y = pi[idx]
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
	return -1
}

// BM算法也是一种精确字符串匹配算法，它采用从右向左比的方法，
// 同时应用到了两种启发式规 则，即坏字符规则 和好后缀规则 ，来决定向右跳跃的距离。
// 基本思路就是从右往左进行字符匹 配，遇到不匹配的字符后从坏字符表和好后缀表找一个最大的右移值，将模式串右移继续匹配。
// BM算法 对后缀蛮力匹配算法的改进
func StringBM() {}

////////////////////
// 回文
////////////////////

//用于存储最后长度最长的子串中间字符的信息
type str struct {
	i   int //位置
	num int //数值
}

func Manacher(s string) string {
	length := len(s)
	btyestr := []byte(s)

	if length < 2 {
		return s
	}

	bytestr2 := make([]byte, 2*length+1)
	p := make([]int, 2*length+1)

	for m := 0; m < length; m++ {

		bytestr2[2*m] = 1
		bytestr2[2*m+1] = btyestr[m]

	}
	bytestr2[2*length] = 1

	fmt.Println(bytestr2)

	//在每个字符间插入相同的字符
	for i := 0; i < len(bytestr2); i++ {

		p[i] = 1
		if i != 0 && i != 2*length {
			for j := 1; ; j++ {
				if bytestr2[i-j] == bytestr2[i+j] {
					p[i] = p[i] + 1
				} else {
					break
				}
				if i-j == 0 || i+j == 2*length {
					break
				}
			}
		}
	}

	maxdata := str{0, 1}

	for j := 0; j < 2*length+1; j++ {

		if p[j] > maxdata.num {

			maxdata.i = j
			maxdata.num = p[j]
		}
	}

	bytestr3 := make([]byte, maxdata.num-1)
	//找出最长子串，跳过之前我自己添加的字符
	for h := 0; h < maxdata.num-1; h++ {
		if (maxdata.num-1)%2 == 0 {
			//跳过之前我自己添加的字符，所以+2*h
			bytestr3[h] = bytestr2[maxdata.i+2-maxdata.num+2*h]
		} else {
			bytestr3[h] = bytestr2[maxdata.i+2-maxdata.num+2*h]
		}

	}

	ss := string(bytestr3)
	return ss
}
