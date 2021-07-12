package four

import "math"

/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

input: [1,6,3,2,5]
output: false
*/

func Hand33(postorder []int) bool {
	// 反向：中、右、左
	// 1、递归分治，中分割然后分别看左右是否满足
	// 2、单调递减
	stack, root := []int{}, math.MaxInt64 // root表示的是当前节点的父节点
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root { // 如果比父节点都大那么肯定就不对，因为一般这里比较都是左边的子树
			return false
		}
		for len(stack) > 0 && postorder[i] < stack[len(stack)-1] { // 找到合适的root节点作为它的左子树
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
