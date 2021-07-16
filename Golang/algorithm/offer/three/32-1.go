package three

import (
	"container/list"
	"wwqdrh/handbook/algorithm/structure"
)

/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树: [3,9,20,null,null,15,7],

  3
   / \
  9  20
    /  \
   15   7

   返回[3,9,20,15,7]
*/

func Hand32_1(root *structure.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		cur := queue.Front()
		queue.Remove(cur)
		curVal := cur.Value.(*structure.TreeNode)
		res = append(res, curVal.Val)
		if curVal.Left != nil {
			queue.PushBack(curVal.Left)
		}
		if curVal.Right != nil {
			queue.PushBack(curVal.Right)
		}
	}
	return res
}
