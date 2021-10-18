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
