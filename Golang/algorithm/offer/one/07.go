package one

import "wwqdrh/handbook/algorithm/tree"

/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

input: preorder = [3,9,20,15,7] inorder = [9,3,15,20,7]
output:
    3
   / \
  9  20
    /  \
   15   7
*/

func Hand7(preorder []int, inorder []int) *tree.TreeNode {
	if len(preorder) == 0 || (len(inorder) != len(preorder)) {
		return nil
	}

	index := 0
	for ; index < len(inorder); index++ {
		if preorder[0] == inorder[index] {
			break
		}
	}

	cur := &tree.TreeNode{Val: inorder[index]}
	cur.Left = Hand7(preorder[1:index+1], inorder[:index])
	cur.Right = Hand7(preorder[index+1:], inorder[index+1:])
	return cur
}
