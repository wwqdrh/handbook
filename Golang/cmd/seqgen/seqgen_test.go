package seqgen

import "testing"

func TestCheckList(t *testing.T) {
	for _, one := range []struct {
		List []int
		Ok   bool
		Next int
	}{
		{[]int{3, 5, 7, 9, 11}, true, 13},
		{[]int{2, 4, 8, 16, 32}, true, 64},
		{[]int{2, 15, 41, 80}, true, 132},
		{[]int{1, 2, 6, 15, 31}, true, 56},
		{[]int{1, 1, 3, 15, 105, 945}, true, 10395},
		{[]int{2, 14, 64, 202, 502, 1062, 2004}, true, 3474},
		{[]int{1, 3}, false, 0},
		{[]int{1, 3, 5}, true, 7},
		{[]int{11, 9, 7, 5, 3}, true, 1},
		{[]int{2, -4, 8, -16, 32}, true, -64},
		{[]int{1, 4, 9, 16, 25}, true, 36},
		{[]int{1, 1, 2, 6}, true, 15},
		{[]int{1, 3, 5, 7, 10}, false, 0},
		{[]int{2, 2, 2, 2, 2}, true, 2},
		{[]int{1, 3, 6, 10, 15, 20}, false, 0},
	} {
		next, ok := checkList(one.List)
		if ok == one.Ok && next == one.Next {
			t.Log("correct", "input", one.List, "get", next, ok, "want", one.Next, one.Ok)
		} else {
			t.Error("wrong", "input", one.List, "get", next, ok, "want", one.Next, one.Ok)
		}
	}
}
