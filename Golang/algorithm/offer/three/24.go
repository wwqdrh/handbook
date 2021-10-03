package three

import "wwqdrh/handbook/algorithm/list"

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

input: 1->2->3->4->5->NULL
output: 5->4->3->2->1->NULL
*/

func Hand24(head *list.ListNode) *list.ListNode {
	var (
		prev *list.ListNode = nil
		cur  *list.ListNode = head
	)
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		cur, prev = next, cur
	}
	return prev
}
