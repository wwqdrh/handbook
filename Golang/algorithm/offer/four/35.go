package four

import "wwqdrh/handbook/algorithm/list"

/**
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
还有一个 random 指针指向链表中的任意节点或者 null。
*/

func Hand35(head *list.ListRandNode) *list.ListRandNode {
	// 1、一种是原链表结构上新建以及插入，这样不用处理rand找元素时候不知新建没有，最后将结果拆分为两条链，一条就是原来的链，一条是复制的链
	// 2、通过map建立新旧对应关系
	maps := make(map[*list.ListRandNode]*list.ListRandNode)
	cur := head
	for cur != nil {
		temp := &list.ListRandNode{
			Val: cur.Val,
		}
		maps[cur] = temp
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		maps[cur].Next = maps[cur.Next]
		maps[cur].Random = maps[cur.Random]
		cur = cur.Next
	}
	return maps[head]
}
