package one

import (
	"testing"
	"wwqdrh/handbook/algorithm/list"
)

func TestHand3(t *testing.T) {
	res := Hand3([]int{2, 3, 1, 0, 2, 5, 3})
	if res != 2 && res != 3 {
		t.Error("发生错误")
	}
}

func TestHand4(t *testing.T) {
	inputs := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if Hand4(inputs, 5) != true {
		t.Error("发生错误")
	}
	if Hand4(inputs, 20) != false {
		t.Error("发生错误")
	}
}

func TestHand5(t *testing.T) {
	if Hand5("We are happy.") != "We%20are%20happy." {
		t.Error("发生错误")
	}
}

func TestHand6(t *testing.T) {
	head := list.NewListNode([]int{1, 3, 2})
	if res := Hand6(head); len(res) != 3 {
		t.Error("发生错误")
	} else {
		target := [3]int{2, 3, 1}
		for i, item := range res {
			if target[i] != item {
				t.Error("发生错误")
			}
		}
	}
}

func TestHand7(t *testing.T) {
	res := Hand7([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	// tree.NewTreeNode([]string{"3", "9", "20", "null", "null", "15", "7"})
	if res.ToString() != "3,9,20,null,null,15,7" {
		t.Error("发生错误")
	}
}

func TestHand9(t *testing.T) {
	queue := Constructor()
	queue.AppendTail(3)
	if queue.DeleteHead() != 3 {
		t.Error("发生错误")
	}
	if queue.DeleteHead() != -1 {
		t.Error("发生错误")
	}
	if queue.DeleteHead() != -1 {
		t.Error("发生错误")
	}
}

func TestHand10_1(t *testing.T) {
	if Hand10_1(2) != 1 {
		t.Error("发生错误")
	}
	if Hand10_1(5) != 5 {
		t.Error("发生错误")
	}
	if Hand10_1(45) != 134903163 {
		t.Error("发生错误")
	}
}

func TestHand10_2(t *testing.T) {
	if Hand10_2(0) != 1 {
		t.Error("发生错误")
	}
	if Hand10_2(7) != 21 {
		t.Error("发生错误")
	}
}

func TestHand11(t *testing.T) {
	if Hand11([]int{1, 3, 5}) != 1 {
		t.Error("发生错误")
	}
	if Hand11([]int{2, 2, 2, 0, 1}) != 0 {
		t.Error("发生错误")
	}
}

func TestHand12(t *testing.T) {
	if !Hand12([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED") {
		t.Error("发生错误")
	}
}
