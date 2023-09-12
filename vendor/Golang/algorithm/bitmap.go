package algorithm

////////////////////
// 位图，状态压缩
// 1、判断元素是否存在
////////////////////

// 不超过3000
func isUniqueString(word string) bool {
	if len(word) == 0 || len(word) > 3000 {
		return false
	}

	// 构建位图
	var mark1, mark2, mark3, mark4 uint64
	var mark *uint64
	for _, r := range word {
		n := uint64(r)
		if n < 64 {
			mark = &mark1
		} else if n < 128 {
			mark = &mark2
			n -= 64
		} else if n < 192 {
			mark = &mark3
			n -= 128
		} else {
			mark = &mark4
			n -= 192
		}
		if (*mark)&(1<<n) != 0 {
			// 说明已经存在了
			return false
		}
		*mark = (*mark) | uint64(1<<n) // 设置状态
	}
	return true
}
