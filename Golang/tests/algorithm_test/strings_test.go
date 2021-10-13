package algorithm_test

import (
	"testing"
	"wwqdrh/handbook/algorithm"
)

func TestIsUniqueString(t *testing.T) {
	if !algorithm.IsUniqueString("abcdefg") {
		t.Error("失败")
	}

	if algorithm.IsUniqueString("abcdeafg") {
		t.Error("失败")
	}
}

func TestReverseString(t *testing.T) {
	if algorithm.ReverseString("abcd") != "dcba" {
		t.Error("失败")
	}
}
