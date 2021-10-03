package three

import "wwqdrh/handbook/algorithm/tree"

/**
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

input: [4,2,7,1,3,6,9]
output: [4,7,2,9,6,3,1]
*/

func Hand27(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	left, right := root.Left, root.Right
	root.Left = Hand27(right)
	root.Right = Hand27(left)
	return root
}
