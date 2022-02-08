package datastruct

import "sync"

// TrieTree 并发安全的字典树结构
type TrieTree struct {
	// 子节点
	children sync.Map
	// 节点类型
	isWord bool
}

// NewTrieTree 字典树构造函数
func NewTrieTree(word string) *TrieTree {
	tree := &TrieTree{
		children: sync.Map{},
		isWord:   false,
	}
	tree.Insert(word)
	return tree
}

// Insert 插入元素
func (trie *TrieTree) Insert(word string) {
	cur := trie
	for _, w := range word {
		val, _ := cur.children.LoadOrStore(w, &TrieTree{
			children: sync.Map{},
		})
		cur = val.(*TrieTree)
	}
	cur.isWord = true
}

// Search 查询
func (trie *TrieTree) Search(word string) bool {
	// 遍历字符串
	cur := trie
	for _, w := range word {
		if val, ok := cur.children.Load(w); !ok {
			return false
		} else {
			// 继续遍历
			cur = val.(*TrieTree)
		}
	}
	return cur.isWord
}

// Delete 删除
func (trie *TrieTree) Delete(word string) {
	// 遍历字符串
	cur := trie
	for _, w := range word {
		if val, ok := cur.children.Load(w); !ok {
			return
		} else {
			cur = val.(*TrieTree)
		}
	}

	cur.isWord = false
}
