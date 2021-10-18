package algorithm_test

import (
	"testing"
	"wwqdrh/handbook/algorithm/sequence"
)

func TestIsUniqueString(t *testing.T) {
	if !sequence.IsUniqueString("abcdefg") {
		t.Error("失败")
	}

	if sequence.IsUniqueString("abcdeafg") {
		t.Error("失败")
	}
}

func TestReverseString(t *testing.T) {
	if sequence.ReverseString("abcd") != "dcba" {
		t.Error("失败")
	}
}
