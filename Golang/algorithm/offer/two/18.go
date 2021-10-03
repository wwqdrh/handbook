package two

import "wwqdrh/handbook/algorithm/list"

/**
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

input: head=[4,5,1,9] val=5
output: [4, 1, 9]
*/

func Hand18(head *list.ListNode, val int) *list.ListNode {
	root := &list.ListNode{Next: head}
	for prev, cur := root, head; cur != nil; {
		if cur.Val == val {
			prev.Next = cur.Next
			cur.Next = nil
			cur = prev.Next
			continue
		}
		cur, prev = cur.Next, prev.Next
	}
	return root.Next
}
