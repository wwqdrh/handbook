package tests

import (
	"strverify/utils/datastruct"
	"sync"
	"testing"
)

func TestPlainTrieTree(t *testing.T) {
	tree := datastruct.NewTrieTree("")
	tree.Insert("abcdefghijk")
	tree.Insert("djaklsda")

	if tree.Search("abcdefgh") {
		t.Error("abcdefghijk---abcdefgh")
	}
	if !tree.Search("abcdefghijk") {
		t.Error("abcdefghijk---abcdefgh")
	}
	if !tree.Search("djaklsda") {
		t.Error("djaklsda---djaklsda")
	}
	if tree.Search("djaklsd") {
		t.Error("djaklsda---djaklsd")
	}
}

func TestConcurrentTrie(t *testing.T) {
	tree := datastruct.NewTrieTree("")
	wg := sync.WaitGroup{}
	wg.Add(10)
	go func() { tree.Insert("abcdefghijk"); wg.Done() }()
	go func() { tree.Insert("djaklsda"); wg.Done() }()
	go func() { tree.Insert("dasdada"); wg.Done() }()
	go func() { tree.Insert("djdaoias"); wg.Done() }()
	go func() { tree.Insert("asdias"); wg.Done() }()
	go func() { tree.Insert("aidoasefghijk"); wg.Done() }()
	go func() { tree.Insert("dasmmda"); wg.Done() }()
	go func() { tree.Insert("ucxzicjnada"); wg.Done() }()
	go func() { tree.Insert("nxzsndaas"); wg.Done() }()
	go func() { tree.Insert("asdpoiqas"); wg.Done() }()
	wg.Wait()

	if !tree.Search("abcdefghijk") {
		t.Error("字典树错误")
	}

	if !tree.Search("aidoasefghijk") {
		t.Error("字典树错误")
	}

	if !tree.Search("nxzsndaas") {
		t.Error("字典树错误")
	}

	if tree.Search("asdpoiqa") {
		t.Error("字典树错误")
	}
}
