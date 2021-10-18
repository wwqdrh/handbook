package algorithm_test

import (
	"reflect"
	"testing"
	"wwqdrh/handbook/algorithm/tree"
)

func TestBinarySearchTree(t *testing.T) {
	nums := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	bst := tree.NewBinarySearchTree(nums, func(i, j int) bool { return nums[i].(int) < nums[j].(int) })
	if !reflect.DeepEqual(bst.MidIter(), []interface{}{1, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error("二叉搜索树中序遍历失败")
	}

	// 测试前序遍历 9 15 7 20 3
	preorder := []interface{}{3, 9, 20, 15, 7}
	inorder := []interface{}{9, 3, 15, 20, 7}
	afterorder := []interface{}{9, 15, 7, 20, 3}
	bt := tree.NewBinaryTreeByMidBefore(preorder, inorder)
	if !reflect.DeepEqual(bt.BeforeIter(), preorder) {
		t.Error("二叉搜索树中序遍历失败")
	}
	if !reflect.DeepEqual(bt.MidIter(), inorder) {
		t.Error("二叉搜索树中序遍历失败")
	}
	if !reflect.DeepEqual(bt.AfterIter(), afterorder) {
		t.Error("二叉搜索树中序遍历失败")
	}
}

func TestAVLTreeHeightAuto(t *testing.T) {
	// 测试高度平衡问题
	res := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var avlTree *tree.AVLTree

	times := 1
	cur := times
	height := 1
	for _, item := range res {
		avlTree = avlTree.Insert(item, func(a, b interface{}) int { return a.(int) - b.(int) })
		if avlTree.GetHeight() != height {
			t.Error("AVL自平衡失败")
		}
		cur--
		if cur == 0 {
			times = times << 1
			height++
			cur = times
		}
	}
}

func TestAVLTreeOrderSequence(t *testing.T) {
	nums := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	avlTree := tree.NewAVLTree(func(i, j interface{}) int { return i.(int) - j.(int) }, nums...)
	if !reflect.DeepEqual(nums, avlTree.MidIter()) {
		t.Error("avl的二叉搜索树性质失效了")
	}
}

type set int

func (s set) Compare(se tree.Entryer) int {
	sh := se.(set)
	if s > sh {
		return -1
	} else if s < sh {
		return 1
	} else {
		return 0
	}
}

func (s set) GetValue() interface{} {
	return s
}
func TestAVLTreeRBTreeOrderSequence(t *testing.T) {
	rb := tree.RBTree{}
	ori := []set{}
	for i := 0; i < 8; i++ {
		s := set(i)
		ori = append(ori, s)
		rb.Insert(s)
	}

	target := []set{}
	for _, item := range rb.MidRec() {
		target = append(target, item.GetEntry().(set))
	}

	if !reflect.DeepEqual(ori, target) {
		t.Error("并非有序")
	}

	// rb.GetRoot()
	// rb.MidRec()
	// fmt.Println()
	// rb.LevelTraversal()
	// fmt.Println()
	// s := set(2)
	// rb.DeleteNode(s)
	// rb.GetRoot()
	// rb.MidRec()
	// fmt.Println()
	// rb.LevelTraversal()

}

func TestUnionFindSet(t *testing.T) {
	uf := tree.InitialUF(10)
	uf.Union(1, 5)
	if uf.GetMaxG() != 2 {
		t.Error("unionfindset error")
	}
}
