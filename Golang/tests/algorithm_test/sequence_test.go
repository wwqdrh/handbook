package algorithm_test

import (
	"reflect"
	"testing"

	"wwqdrh/handbook/algorithm/sequence"
)

func TestLinkedList(t *testing.T) {
	list := sequence.NewLinkedList([]interface{}{1, 2, 3, 4, 5})
	if !reflect.DeepEqual([]interface{}{1, 2, 3, 4, 5}, list.ToSlice()) {
		t.Error("linkedlist失败")
	}
}

func TestLinkedListMerge(t *testing.T) {
	list := sequence.NewLinkedList([]interface{}{1, 2, 3, 4, 5})
	list2 := sequence.NewLinkedList([]interface{}{6, 7, 8})
	list.Merge(list2)
	if !reflect.DeepEqual([]interface{}{1, 2, 3, 4, 5, 6, 7, 8}, list.ToSlice()) {
		t.Error("linkedlist失败")
	}
}

func TestLinkedListLastK(t *testing.T) {
	list := sequence.NewLinkedList([]interface{}{1, 2, 3, 4, 5})
	if list.LastK(2).(int) != 4 {
		t.Error("lask失败")
	}
	if list.LastK(5).(int) != 1 {
		t.Error("lask失败")
	}
	if list.LastK(6) != nil {
		t.Error("lask失败")
	}
}

func TestLinkedListReversed(t *testing.T) {
	list := sequence.NewLinkedList([]interface{}{1, 2, 3, 4, 5})
	list.Reversed()
	if !reflect.DeepEqual([]interface{}{5, 4, 3, 2, 1}, list.ToSlice()) {
		t.Error("linkedlist反序失败")
	}
}

func TestKMPString(t *testing.T) {
	if sequence.StringKMP("abcdecdfg", "cdfg") != 5 {
		t.Error("KMP算法出错")
	}
	if sequence.StringKMP("abcdecdeffg", "cdef") != 5 {
		t.Error("KMP算法出错")
	}
}

func TestStringCommonPrefix(t *testing.T) {
	if sequence.StringCommonPrefix("你好世界", "你好", "你好") != 2 {
		t.Error("公共前缀算法出错")
	}
	if sequence.StringCommonPrefix("abcdecdeffg", "abc", "ab") != 2 {
		t.Error("公共前缀算法出错")
	}
}

func TestStringPalidorm(t *testing.T) {
	if string(sequence.StringMaxSubStrPalindrome([]rune("abcddcba"))) != "abcddcba" {
		t.Error("最大回文串算法出错")
	}
}

func TestStringPalidormBuild(t *testing.T) {
	if sequence.StringMaxPalindromeBuild([]rune("abccccdd")) != 7 {
		t.Error("最大回文串构造算法出错")
	}
}

func TestStackQueue(t *testing.T) {
	queue := sequence.NewStackQueue()
	queue.Push(1)
	if val, _ := queue.Pop(); val.(int) != 1 {
		t.Error("stackqueue失败")
	}
	queue.Push(1)
	queue.Push(2)
	if val, _ := queue.Pop(); val.(int) != 1 {
		t.Error("stackqueue失败")
	}
	if val, _ := queue.Pop(); val.(int) != 2 {
		t.Error("stackqueue失败")
	}
	if _, err := queue.Pop(); err == nil {
		t.Error("stackqueue失败")
	}
}

func TestVerfySeq(t *testing.T) {
	if !sequence.VerifySeq([]interface{}{1, 2, 3, 4, 5}, []interface{}{3, 2, 1, 5, 4}) {
		t.Error("verify seq 失败")
	}
	if sequence.VerifySeq([]interface{}{1, 2, 3, 4, 5}, []interface{}{4, 3, 1, 2, 5}) {
		t.Error("verify seq 失败")
	}
}
