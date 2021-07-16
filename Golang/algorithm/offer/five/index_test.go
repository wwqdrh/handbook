package five

import "testing"

func TestHand41(t *testing.T) {
	// TODO: 暂无测试用例
}

func TestHand42(t *testing.T) {
	if Hand42([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("发生错误")
	}
}

func TestHand43(t *testing.T) {
	if Hand43(12) != 5 {
		t.Error("发生错误")
	}
	if Hand43(13) != 6 {
		t.Error("发生错误")
	}
}

func TestHand44(t *testing.T) {
	if Hand44(3) != 3 {
		t.Error("发生错误")
	}
	if Hand44(11) != 0 {
		t.Error("发生错误")
	}
}

func TestHand45(t *testing.T) {
	if Hand45([]int{3, 30, 34, 5, 9}) != "3033459" {
		t.Error("发生错误")
	}
}

func TestHand46(t *testing.T) {
	if Hand46(12258) != 5 {
		t.Error("发生错误")
	}
}

func TestHand47(t *testing.T) {
	if Hand47([][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}) != 12 {
		t.Error("发生错误")
	}
}

func TestHand48(t *testing.T) {
	if Hand48("abcabcbb") != 3 {
		t.Error("发生错误")
	}
}

func TestHand49(t *testing.T) {
	if Hand49(10) != 12 {
		t.Error("发生错误")
	}
}

func TestHand50(t *testing.T) {
	if Hand50("abaccdeff") != 'b' {
		t.Error("发生错误")
	}
}
