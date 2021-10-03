package five

/**
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
*/

func Hand48(s string) int {
	// 滑动窗口，当发现右边对象已经在vis里存在的时候，更新左边边界，遍历的过程不断维护最大值
	if len(s) == 0 {
		return 0
	}
	vis := make(map[rune]int)
	i, res := -1, 0
	for j, ch := range []rune(s) {
		if idx, ok := vis[ch]; ok {
			i = max(idx, i) // 更新左边界
		}
		vis[ch] = j
		res = max(res, j-i)
	}
	return res
}
