package three

import "container/list"

/**
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。

例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
output: true
*/

func Hand31(pushed []int, popped []int) bool {
	stack, i := list.New(), 0
	for _, num := range pushed {
		stack.PushBack(num)
		for stack.Len() > 0 {
			if elem := stack.Back(); elem != nil && elem.Value.(int) == popped[i] {
				stack.Remove(elem)
				i++
			} else {
				break
			}
		}
	}
	return stack.Len() == 0
}
