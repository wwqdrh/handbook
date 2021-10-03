package four

import "wwqdrh/handbook/algorithm/tree"

/**
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]
*/

func Hand32_3(root *tree.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res, queue := [][]int{}, []*tree.TreeNode{root}
	level, flag := 0, 1 // 奇数层
	for len(queue) > 0 {
		res = append(res, []int{})
		if flag == 1 {
			for _, v := range queue {
				res[level] = append(res[level], v.Val)
			}
		} else {
			for i := len(queue) - 1; i >= 0; i-- {
				res[level] = append(res[level], queue[i].Val)
			}
		}
		tmp := []*tree.TreeNode{}
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		queue = tmp
		flag ^= 1
		level++
	}
	return res
}
