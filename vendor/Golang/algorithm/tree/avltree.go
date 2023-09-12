package tree

// 自平衡二叉树 带有平衡条件的二叉查找树
// TODO: 目前版本只能是不允许重复元素
// 某结点的左子树与右子树的高度(深度)差即为该结点的平衡因子（Balance Factor）。平衡二叉树上所有结点的平衡因子只可能是 -1，0 或 1。
// 由于插入操作可能会破坏AVL树的平衡特性，故在插入完成之前通过对树进行简单修正来恢复平衡

type compareFunc func(a, b interface{}) int

type AVLTree struct {
	value  interface{}
	height int
	left   *AVLTree
	right  *AVLTree
}

// 新的avl
func NewAVLTree(cb compareFunc, nodes ...interface{}) *AVLTree {
	var tree *AVLTree
	for _, node := range nodes {
		tree = tree.Insert(node, cb)
	}
	return tree
}

func (t *AVLTree) leftRotate() *AVLTree {
	headNode := t.right
	t.right = headNode.left
	headNode.left = t
	//更新结点高度
	t.height = max(t.left.GetHeight(), t.right.GetHeight()) + 1
	headNode.height = max(headNode.left.GetHeight(), headNode.right.GetHeight()) + 1
	return headNode
}

func (t *AVLTree) rightRotate() *AVLTree {
	headNode := t.left
	t.left = headNode.right
	headNode.right = t
	//更新结点高度
	t.height = max(t.left.GetHeight(), t.right.GetHeight()) + 1
	headNode.height = max(headNode.left.GetHeight(), headNode.right.GetHeight()) + 1
	return headNode
}

func (t *AVLTree) leftRightRotate() *AVLTree {
	//以失衡点左结点先左旋转
	sonHeadNode := t.left.leftRotate()
	t.left = sonHeadNode
	//再以失衡点左旋转
	return t.rightRotate()
}

func (t *AVLTree) rightLeftRotate() *AVLTree {
	//以失衡点右结点先右旋转
	sonHeadNode := t.right.rightRotate()
	t.right = sonHeadNode
	//再以失衡点左旋转
	return t.leftRotate()
}

func (t *AVLTree) adjust() *AVLTree {
	if t.left.GetHeight()-t.right.GetHeight() == 2 {
		if t.left.left.GetHeight() > t.left.right.GetHeight() {
			t = t.rightRotate()
		} else {
			t = t.rightLeftRotate()
		}
	} else if t.right.GetHeight()-t.left.GetHeight() == 2 {
		if t.right.right.GetHeight() > t.right.left.GetHeight() {
			t = t.leftRotate()
		} else {
			t = t.leftRightRotate()
		}
	}

	return t
}

func (t *AVLTree) GetHeight() int {
	if t == nil {
		return 0
	}
	return t.height
}

//查找子树最小值
func (t *AVLTree) getMin() interface{} {
	if t == nil {
		return -1
	}
	if t.left == nil {
		return t.value
	} else {
		return t.left.getMin()
	}
}

// 查找元素
func (t *AVLTree) Search(value interface{}, cb compareFunc) bool {
	if t == nil {
		return false
	}
	if status := cb(value, t.value); status < 0 {
		return t.left.Search(value, cb)
	} else if status > 0 {
		return t.right.Search(value, cb)
	} else {
		return true
	}
}

// 添加元素
func (t *AVLTree) Insert(value interface{}, cb compareFunc) *AVLTree {
	if t == nil {
		return &AVLTree{value: value, height: 1}
	}
	if status := cb(value, t.value); status < 0 {
		t.left = t.left.Insert(value, cb)
		t = t.adjust()
	} else if status > 0 {
		t.right = t.right.Insert(value, cb)
		t = t.adjust()
	} else {
		println("the node exists")
	}

	if left, right := t.left.GetHeight(), t.right.GetHeight(); left >= right {
		t.height = left + 1
	} else {
		t.height = right + 1
	}

	return t
}

// 删除元素
func (t *AVLTree) Delete(value interface{}, cb compareFunc) *AVLTree {
	if t == nil {
		return t
	}
	if status := cb(value, t.value); status < 0 {
		t.left = t.left.Delete(value, cb)
	} else if status > 0 {
		t.right = t.right.Delete(value, cb)
	} else {
		if t.left != nil && t.right != nil {
			t.value = t.right.getMin()
			t.right = t.right.Delete(t.value, cb)
		} else if t.left != nil {
			t = t.left
		} else {
			t = t.right
		}
	}
	if t != nil {
		t.height = max(t.left.GetHeight(), t.right.GetHeight()) + 1
		t = t.adjust()
	}
	return t
}

// 遍历 TODO: 最好将二叉树这些公共方法抽离出来

// 迭代 从某一个元素开始将左节点加入栈直到nil
// 弹出一个元素将值加入结果 并根据当前元素的右元素为起点将左节点都加入
func (t *AVLTree) MidIter() []interface{} {
	res := []interface{}{}

	for stack, cur := []*AVLTree{}, t; cur != nil || len(stack) > 0; {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, top.value)
		cur = top.right
	}

	return res
}
