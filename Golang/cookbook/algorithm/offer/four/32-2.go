package four

import "wwqdrh/handbook/cookbook/algorithm/structure"

/**
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

[
  [3],
  [9,20],
  [15,7]
]
*/

func Hand32_2(root *structure.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	level := 0
	queue := []*structure.TreeNode{root}
	for len(queue) > 0 {
		res = append(res, []int{})
		tmp := []*structure.TreeNode{}
		for _, i := range queue {
			res[level] = append(res[level], i.Val)

			if i.Left != nil {
				tmp = append(tmp, i.Left)
			}
			if i.Right != nil {
				tmp = append(tmp, i.Right)
			}
		}
		queue = tmp
		level++
	}
	return res
}
