package sequence

// 根据压入顺序判断 弹出顺序是否正确
// 模拟
// 使用一个辅助栈 如果不是出栈元素继续压栈
// 判断最后辅助栈是否为空
func VerifySeq(inStack []interface{}, outStack []interface{}) bool {
	helpStack := make([]interface{}, 0, len(inStack))

	outIdx := 0
	for inIdx := 0; inIdx < len(inStack); inIdx++ {
		helpStack = append(helpStack, inStack[inIdx])
		for len(helpStack) > 0 && helpStack[len(helpStack)-1] == outStack[outIdx] {
			helpStack = helpStack[:len(helpStack)-1]
			outIdx++
		}
	}

	return len(helpStack) == 0 && outIdx == len(outStack)
}
