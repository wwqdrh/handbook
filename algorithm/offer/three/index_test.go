package three

import (
	"testing"
	"wwqdrh/handbook/algorithm/structure"
	"wwqdrh/handbook/algorithm/utils"
)

func TestHand22(t *testing.T) {
	res := Hand22(structure.NewListNode([]int{1, 2, 3, 4, 5}), 2).ToSlice()
	if res[0] != 4 || res[1] != 5 {
		t.Error("发生错误")
	}
}

func TestHand24(t *testing.T) {
	res := Hand24(structure.NewListNode([]int{1, 2, 3})).ToSlice()
	if res[0] != 3 || res[1] != 2 || res[2] != 1 {
		t.Error("发生错误")
	}
}

func TestHand25(t *testing.T) {
	if !utils.IntSliceCompare(
		Hand25(structure.NewListNode([]int{1, 2, 4}), structure.NewListNode([]int{1, 3, 4})).ToSlice(),
		[]int{1, 1, 2, 3, 4, 4},
	) {
		t.Error("发生错误")
	}

}

func TestHand26(t *testing.T) {
	if Hand26(structure.NewTreeNode([]string{"1", "2", "3"}), structure.NewTreeNode([]string{"3", "1"})) {
		t.Error("发生错误")
	}
}

func TestHand27(t *testing.T) {
	if Hand27(structure.NewTreeNode([]string{"4", "2", "7", "1", "3", "6", "9"})).ToString() != "4,7,2,9,6,3,1" {
		t.Error("发生错误")
	}
}

func TestHand28(t *testing.T) {
	if !Hand28(structure.NewTreeNode([]string{"1", "2", "2", "3", "4", "4", "3"})) {
		t.Error("发生错误")
	}
	if Hand28(structure.NewTreeNode([]string{"1", "2", "2", "null", "3", "null", "3"})) {
		t.Error("发生错误")
	}
}

func TestHand29(t *testing.T) {
	if !utils.IntSliceCompare(Hand29([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5}) {
		t.Error("发生错误")
	}
}

func TestHand30(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	if stack.Top() != 2 {
		t.Error("发生错误")
	}
	if stack.Min() != 1 {
		t.Error("发生错误")
	}
	stack.Pop()
	if stack.Min() != 1 {
		t.Error("发生错误")
	}
	if stack.Top() != 1 {
		t.Error("发生错误")
	}
}

func TestHand31(t *testing.T) {
	if Hand31([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}) {
		t.Error("发生错误")
	}
}

func TestHand32_1(t *testing.T) {
	if !utils.IntSliceCompare(
		Hand32_1(structure.NewTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})),
		[]int{3, 9, 20, 15, 7},
	) {
		t.Error("发生错误")
	}
}
