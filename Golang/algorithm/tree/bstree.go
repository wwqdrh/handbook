package tree

import "sort"

type BinaryTree struct {
	val   interface{}
	left  *BinaryTree
	right *BinaryTree
}

// 二叉搜索树
// 左边比根节点小 右边比根节点大
type BinarySearchTree struct {
	*BinaryTree
}

// 构建二叉搜索树
// 有序序列是二叉树的中序遍历结果
// 可以根据中点来建立一个高度平衡的二叉搜索树
func NewBinarySearchTree(nums []interface{}, cb func(i, j int) bool) *BinarySearchTree {
	sort.Slice(nums, cb)

	var build func(i, j int) *BinaryTree
	build = func(i, j int) *BinaryTree {
		if i > j {
			return nil
		}
		if i == j {
			return &BinaryTree{val: nums[i]}
		}

		idx := (i + j) >> 1
		return &BinaryTree{
			val:   nums[idx],
			left:  build(i, idx-1),
			right: build(idx+1, j),
		}
	}

	return &BinarySearchTree{BinaryTree: build(0, len(nums)-1)}
}

// 构建二叉树
// 通过中序以及前序构建二叉树
func NewBinaryTreeByMidBefore(preorder []interface{}, inorder []interface{}) *BinaryTree {
	if len(preorder) == 0 || (len(inorder) != len(preorder)) {
		return nil
	}

	index := 0
	for ; index < len(inorder); index++ {
		if preorder[0] == inorder[index] {
			break
		}
	}

	cur := &BinaryTree{inorder[index], nil, nil}
	cur.left = NewBinaryTreeByMidBefore(preorder[1:index+1], inorder[:index])
	cur.right = NewBinaryTreeByMidBefore(preorder[index+1:], inorder[index+1:])

	return cur
}

////////////////////
// 遍历
////////////////////

// 前序遍历 迭代 将根打印出来 然后右子树加入栈 左子树加入栈
func (t *BinaryTree) BeforeIter() []interface{} {
	res := []interface{}{}

	for stack := []*BinaryTree{t}; len(stack) > 0; {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == nil {
			continue
		}

		res = append(res, cur.val)
		stack = append(stack, cur.right, cur.left)
	}

	return res
}

// 迭代 从某一个元素开始将左节点加入栈直到nil
// 弹出一个元素将值加入结果 并根据当前元素的右元素为起点将左节点都加入
func (t *BinaryTree) MidIter() []interface{} {
	res := []interface{}{}

	for stack, cur := []*BinaryTree{}, t; cur != nil || len(stack) > 0; {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, top.val)
		cur = top.right
	}

	return res
}

// 中 右 左 =逆序=> 左 右 后
func (t *BinaryTree) AfterIter() []interface{} {
	res := []interface{}{}

	for stack := []*BinaryTree{t}; len(stack) > 0; {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.val)

		if curr.left != nil {
			stack = append(stack, curr.left)
		}
		if curr.right != nil {
			stack = append(stack, curr.right)
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

////////////////////
// 二叉树修改
////////////////////
func (t *BinaryTree) Insert(val interface{}, cb func(a, b interface{}) bool) bool {

	return false // 插入失败
}

func (t *BinaryTree) Delete(val interface{}) bool {

	return false // 删除失败
}
