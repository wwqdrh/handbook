package tree

import (
	"errors"
)

// 线段树，区间树，长度不变的平衡树结构
// 线段树，在区间统计，区间染色这类区间不变的问题中，线段树更好gen
// 线段树，维护区间信息，在O(logN)时间复杂度内实现单点修改，区间修改，区间查询(区间求和，求区间最大值，求区间最小值等)

// SegmentTree 使用数组实现线段树结构
// 如果有n个元素，如果n=2^k, 只需要2n空间，如果n=2^k+1, 需要有4*n个节点
// 使用线段树时，不考虑添加元素，一般采用4n的静态空间即可
type SegmentTree struct {
	tree   []int                // 线段树
	data   []int                // 数组数据
	merger func(v1, v2 int) int // 线段树功能函数，求和求余等等
}

func NewSegmentTree(arrs []int, merger func(i1, i2 int) int) *SegmentTree {
	length := len(arrs)
	tree := &SegmentTree{
		tree:   make([]int, length*4),
		data:   arrs,
		merger: merger,
	}
	tree.buildSegmentTree(0, 0, length-1)
	return tree
}

func (SegmentTree) leftChild(i int) int {
	return 2*i + 1
}

func (tree *SegmentTree) buildSegmentTree(index, l, r int) int {
	if l == r {
		tree.tree[index] = tree.data[l]
		return tree.data[l]
	}
	leftI := tree.leftChild(index)
	rightI := leftI + 1
	mid := l + (r-l)/2
	leftResp := tree.buildSegmentTree(leftI, l, mid)
	rightResp := tree.buildSegmentTree(rightI, mid+1, r)

	tree.tree[index] = tree.merger(leftResp, rightResp)
	return tree.tree[index]
}

// Query 查询数据
func (tree *SegmentTree) Query(queryL, queryR int) (int, error) {
	length := len(tree.data)
	if queryL < 0 || queryL > queryR || queryR >= length {
		return 0, errors.New("index is illegal")
	}
	return tree.queryRange(0, 0, length-1, queryL, queryR), nil
}

// queryRange 具体的查询逻辑
func (tree *SegmentTree) queryRange(index, l, r, queryL, queryR int) int {
	if l == queryL && r == queryR {
		return tree.tree[index]
	}

	leftI := tree.leftChild(index)
	rightI := leftI + 1
	mid := l + (r-l)/2
	if queryL > mid {
		return tree.queryRange(rightI, mid+1, r, queryL, queryR)
	}
	if queryR <= mid {
		return tree.queryRange(leftI, l, mid, queryL, queryR)
	}

	leftResp := tree.queryRange(leftI, l, mid, queryL, mid)
	rightResp := tree.queryRange(rightI, mid+1, r, mid+1, queryR)
	return tree.merger(leftResp, rightResp)
}

func (tree *SegmentTree) Update(k, v int) {
	length := len(tree.data)
	if k < 0 || k >= length {
		return
	}
	tree.set(0, 0, length-1, k, v)
}

func (tree *SegmentTree) set(treeIndex, l, r, k, v int) {
	if l == r {
		tree.tree[treeIndex] = v
		return
	}

	leftI := tree.leftChild(treeIndex)
	rightI := leftI + 1
	midI := l + (r-l)/2
	if k > midI {
		tree.set(rightI, midI+1, r, k, v)
	} else {
		tree.set(leftI, l, midI, k, v)
	}
	tree.tree[treeIndex] = tree.merger(tree.tree[leftI], tree.tree[rightI])
}
