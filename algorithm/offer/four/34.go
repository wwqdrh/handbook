package four

import "wwqdrh/handbook/algorithm/structure"

/**
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。



示例:
给定如下二叉树，以及目标和 target = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1

[
   [5,4,11,2],
   [5,8,4,5]
]
*/

func Hand34(root *structure.TreeNode, target int) [][]int {
	res := make([][]int, 0)
	perm := make([]int, 0)
	var backtrack func(node *structure.TreeNode, n int) // n: 当前需要的和
	backtrack = func(node *structure.TreeNode, n int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if n == node.Val {
				t := []int{}
				t = append(t, perm...)
				t = append(t, node.Val)
				res = append(res, t)
			}
			return
		}
		perm = append(perm, node.Val)
		backtrack(node.Left, n-node.Val)
		backtrack(node.Right, n-node.Val)
		perm = perm[:len(perm)-1]
	}
	backtrack(root, target)
	return res
}
