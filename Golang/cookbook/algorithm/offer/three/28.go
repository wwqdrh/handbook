package three

import "wwqdrh/handbook/cookbook/algorithm/structure"

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

input: [1,2,2,3,4,4,3]
output: true

input: [1,2,2,null,3,null,3]
output: false
*/

func Hand28(root *structure.TreeNode) bool {
	if root == nil {
		return true
	}
	var isSym func(left, right *structure.TreeNode) bool
	isSym = func(left, right *structure.TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return isSym(left.Left, right.Right) && isSym(left.Right, right.Left)
	}
	return isSym(root.Left, root.Right)
}
