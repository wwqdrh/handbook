package one

import "wwqdrh/handbook/algorithm/list"

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

input: [1,3,2] output: [2, 3, 1]
*/

func Hand6(head *list.ListNode) []int {
	if head == nil {
		return []int{}
	}

	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	res := make([]int, length)
	idx := length - 1
	for cur := head; cur != nil; cur = cur.Next {
		res[idx] = cur.Val
		idx--
	}

	return res
}
