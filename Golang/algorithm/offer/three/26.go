package three

import (
	"container/list"
	"wwqdrh/handbook/algorithm/tree"
)

/**
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

*/

func Hand26(A *tree.TreeNode, B *tree.TreeNode) bool {
	var isPart func(a *tree.TreeNode, b *tree.TreeNode) bool
	isPart = func(a *tree.TreeNode, b *tree.TreeNode) bool {
		if b == nil {
			return true
		}
		if a == nil || a.Val != b.Val {
			return false
		}
		return isPart(a.Left, b.Left) && isPart(a.Right, b.Right)
	}

	if A == nil || B == nil {
		return false
	}
	queue := list.New()
	queue.PushBack(A)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		frontVal := front.Value.(*tree.TreeNode)
		if isPart(frontVal, B) {
			return true
		}
		if frontVal.Left != nil {
			queue.PushBack(frontVal.Left)
		}
		if frontVal.Right != nil {
			queue.PushBack(frontVal.Right)
		}
	}
	return false
}
