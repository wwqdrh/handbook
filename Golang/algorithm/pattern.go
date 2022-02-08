package algorithm

import (
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
