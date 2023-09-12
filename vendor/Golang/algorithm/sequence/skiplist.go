// 跳表

package sequence

import "math/rand"

const maxLevel = 32  // maxLevel 表示元素最多可以有多少个向前的指针，它决定了跳跃表的容量
const pFactor = 0.25 // 一个概率因子，它决定了随机生成 level 的大小分布。

type SkiplistNode struct {
	val     int
	forward []*SkiplistNode
}

type Skiplist struct {
	head  *SkiplistNode
	level int
}

func NewSkiplist() Skiplist {
	return Skiplist{&SkiplistNode{-1, make([]*SkiplistNode, maxLevel)}, 0}
}

func (Skiplist) randomLevel() int {
	// 水塘抽样?
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

func (s *Skiplist) Search(target int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 target 的元素
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	// 检测当前元素的值是否等于 target
	return curr != nil && curr.val == target
}

func (s *Skiplist) Add(num int) {
	// 由于新节点的高度可能挡住原来跳表关系，所以这里需要记录一下各个level的下一个节点是什么
	update := make([]*SkiplistNode, maxLevel)
	for i := range update {
		update[i] = s.head
	}
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	lv := s.randomLevel()
	s.level = max(s.level, lv)
	newNode := &SkiplistNode{num, make([]*SkiplistNode, lv)}
	for i, node := range update[:lv] {
		// 对第 i 层的状态进行更新，将当前元素的 forward 指向新的节点
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}

func (s *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	// 如果值不存在则返回 false
	if curr == nil || curr.val != num {
		return false
	}
	for i := 0; i < s.level && update[i].forward[i] == curr; i++ {
		// 对第 i 层的状态进行更新，将 forward 指向被删除节点的下一跳
		update[i].forward[i] = curr.forward[i]
	}
	// 更新当前的 level
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}
	return true
}

func (sl *Skiplist) Front() *SkiplistNode {
	return sl.head.forward[0]
}

// Next returns first element after e.
func (e *SkiplistNode) Next() *SkiplistNode {
	if e != nil {
		return e.forward[0]
	}
	return nil
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
