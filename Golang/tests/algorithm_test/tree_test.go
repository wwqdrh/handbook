package algorithm_test

import (
	"testing"
	"wwqdrh/handbook/algorithm/tree"
)

func TestUnionFindSet(t *testing.T) {
	uf := tree.InitialUF(10)
	uf.Union(1, 5)
	if uf.GetMaxG() != 2 {
		t.Error("unionfindset error")
	}
}
