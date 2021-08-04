package two

import (
	"testing"
	"wwqdrh/handbook/cookbook/algorithm/structure"
)

func TestHand13(t *testing.T) {
	if Hand13(2, 3, 1) != 3 {
		t.Error("发生错误")
	}
	if Hand13(3, 1, 0) != 1 {
		t.Error("发生错误")
	}
}

func TestHand14_1(t *testing.T) {
	if Hand14_1(2) != 1 {
		t.Error("发生错误")
	}
	if Hand14_1(10) != 36 {
		t.Error("发生错误")
	}
}

func TestHand14_2(t *testing.T) {
	if Hand14_2(2) != 1 {
		t.Error("发生错误")
	}
	if Hand14_2(10) != 36 {
		t.Error("发生错误")
	}
}

func TestHand15(t *testing.T) {
	if Hand15(11) != 3 {
		t.Error("发生错误")
	}
	if Hand15(128) != 1 {
		t.Error("发生错误")
	}
}

func TestHand16(t *testing.T) {
	if Hand16(float64(2), 10) != float64(1024) {
		t.Error("发生错误")
	}
	if int(Hand16(float64(2.1), 3)*1000) != 9261 {
		t.Error("发生错误")
	}
	if int(Hand16(float64(2), -2)*100) != 25 {
		t.Error("发生错误")
	}
}

func TestHand17(t *testing.T) {
	if Hand17(1)[8] != 9 {
		t.Error("出现错误")
	}
}

func TestHand18(t *testing.T) {
	if Hand18(structure.NewListNode([]int{4, 5, 1, 9}), 5).ToSlice()[2] != 9 {
		t.Error("发生错误")
	}
}

func TestHand19(t *testing.T) {
	if Hand19("aa", "a") {
		t.Error("发生错误")
	}
}

func TestHand20(t *testing.T) {
	if !Hand20("0") {
		t.Error("发生错误")
	}
	if Hand20("e") {
		t.Error("发生错误")
	}
}

func TestHand21(t *testing.T) {
	if Hand21([]int{1, 2, 3, 4})[1]&1 != 1 {
		t.Error("发生错误")
	}
}
