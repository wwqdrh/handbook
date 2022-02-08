package tree

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
		if _, ok := cur.children.Load(w); !ok {
			node := &TrieTree{
				children: sync.Map{},
			}
			cur.children.Store(w, node)
		}
		val, _ := cur.children.Load(w)
		cur = val.(*TrieTree)
	}
	cur.isWord = true
}

// Search 查询
func (trie *TrieTree) Search(word string) bool {
	// 遍历字符串
	cur := trie
	for _, w := range word {
		// 如果发现有字符不存在，则说明不匹配
		if val, ok := cur.children.Load(w); !ok {
			return false
		} else {
			cur = val.(*TrieTree)
		}
	}
	// 遍历完整个字符后，且同时满足字符串已结束，说明匹配成功
	// 如果存在sleep，而搜索的sle,则为false
	return cur.isWord
}

// Delete 删除
func (trie *TrieTree) Delete(word string) {
	// 遍历字符串
	cur := trie
	for _, w := range word {
		// 如果发现有字符不存在，则说明不匹配
		if val, ok := cur.children.Load(w); !ok {
			return
		} else {
			cur = val.(*TrieTree)
		}
	}

	cur.isWord = false
}
