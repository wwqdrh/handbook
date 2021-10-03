package three

import "wwqdrh/handbook/algorithm/list"

/**
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

input：1->2->4, 1->3->4
output：1->1->2->3->4->4
*/

func Hand25(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	p1, p2 := l1, l2
	dummy := &list.ListNode{}

	cur := dummy
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}
	if p1 != nil {
		cur.Next = p1
	} else if p2 != nil {
		cur.Next = p2
	}
	return dummy.Next
}
