package four

import (
	"testing"
	"wwqdrh/handbook/algorithm/structure"
	"wwqdrh/handbook/algorithm/utils"
)

func TestHand32_2(t *testing.T) {
	// Hand32_2() // TODO 暂时无测试用例
}

func TestHand32_3(t *testing.T) {
	// Hand32_3() // TODO 暂时无测试用例
}

func TestHand33(t *testing.T) {
	if Hand33([]int{1, 6, 3, 2, 5}) {
		t.Error("发生错误")
	}
	if !Hand33([]int{1, 3, 2, 6, 5}) {
		t.Error("发生错误")
	}
}

func TestHand34(t *testing.T) {
	res := Hand34(structure.NewTreeNode([]string{"5", "4", "8", "11", "null", "13", "4", "7", "2", "null", "null", "5", "1"}), 22)
	if len(res) != 2 || !(utils.IntSliceCompare(res[0], []int{5, 4, 11, 2}) || utils.IntSliceCompare(res[1], []int{5, 4, 11, 2})) {
		t.Error("发生错误")
	}
}

func TestHand35(t *testing.T) {
	// Hand35() // TODO: 暂无测试方法
}

func TestHand36(t *testing.T) {
	res := Hand36(structure.NewTreeNode([]string{"4", "2", "5", "1", "3"}))
	if res.Val != 1 || res.Right.Val != 2 || res.Left.Val != 5 {
		t.Error("发生错误")
	}
}

func TestHand37(t *testing.T) {
	Hand37()
}

func TestHand38(t *testing.T) {
	if !utils.StringSliceCompare(Hand38("abc"), []string{"abc", "acb", "bac", "bca", "cab", "cba"}) {
		t.Error("发生错误")
	}
}

func TestHand39(t *testing.T) {
	if Hand39([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}) != 2 {
		t.Error("发生错误")
	}
}

func TestHand40(t *testing.T) {
	if !utils.IntSliceCompare(Hand40([]int{0, 1, 2, 1}, 1), []int{0}) {
		t.Error("发生错误")
	}
}
