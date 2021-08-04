package five

/**
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

input: "abaccdeff"
output: "b"
*/

func Hand50(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for _, ch := range s {
		if cnt[ch-'a'] == 1 {
			return byte(ch)
		}
	}
	return ' '
}
