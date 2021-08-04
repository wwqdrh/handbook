package four

import (
	"container/list"
	"wwqdrh/handbook/cookbook/algorithm/structure"
)

/**
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。

 我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

*/

func Hand36(root *structure.TreeNode) *structure.TreeNode {
	if root == nil {
		return nil
	}
	var pre, head, node *structure.TreeNode

	stack := list.New()
	for stack.Len() > 0 || root != nil {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		if elm := stack.Back(); elm != nil {
			node = elm.Value.(*structure.TreeNode)
			stack.Remove(elm)
		}
		if pre == nil {
			head = node
		} else {
			pre.Right = node
		}
		node.Left = pre
		pre, root = node, node.Right
	}
	head.Left, node.Right = node, head
	return head
}
