package tree

import (
	"container/heap"
	"errors"
)

// 霍夫曼树
// wpl值最小的二叉树

type HuffmanTree struct {
	key          string
	value        float64
	ltree, rtree *HuffmanTree
}

type minHeap []*HuffmanTree

func (h *minHeap) Len() int             { return len((*h)) }
func (h *minHeap) Less(i, j int) bool   { return (*h)[i].value < (*h)[j].value }
func (h *minHeap) Swap(i, j int)        { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(val interface{}) { *h = append(*h, val.(*HuffmanTree)) }
func (h *minHeap) Pop() interface{} {
	t := *h
	res := t[len(t)-1]
	*h = t[:len(t)-1]
	return res
}

func NewHuffmanTree(keys []string, value []float64) (tree *HuffmanTree, err error) {
	if len(keys) != len(value) {
		return nil, errors.New("keys与对应的value数量不一致")
	}

	n := len(keys)
	forsets := new(minHeap)
	for i := 0; i < n; i++ {
		heap.Push(forsets, &HuffmanTree{
			key:   keys[i],
			value: value[i],
		})
	}

	// 迭代n-1次构造huffman树
	var root *HuffmanTree
	for i := 1; i < n; i++ {
		// 获取堆中的最小以及第二小节点 使用堆(插入与删除O(logn))比直接遍历找最小的两个(O(n^2))时间复杂度低
		minn := heap.Pop(forsets).(*HuffmanTree)
		minnSub := heap.Pop(forsets).(*HuffmanTree)
		root = &HuffmanTree{
			key:   minn.key + minnSub.key,
			value: minn.value + minnSub.value,
			ltree: minn,
			rtree: minnSub,
		}
		heap.Push(forsets, root)
	}
	return root, nil
}

func (t *HuffmanTree) WPLValue(idx int) float64 {
	if t == nil {
		return 0
	}

	if t.ltree == nil && t.rtree == nil {
		return t.value * float64(idx)
	} else {
		left := t.ltree.WPLValue(idx + 1)
		right := t.rtree.WPLValue(idx + 1)
		return left + right
	}
}
