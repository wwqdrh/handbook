package five

import "container/heap"

/**
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
*/
type minHeap []int

func (h *minHeap) Len() int             { return len((*h)) }
func (h *minHeap) Less(i, j int) bool   { return (*h)[i] < (*h)[j] }
func (h *minHeap) Swap(i, j int)        { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(val interface{}) { *h = append(*h, val.(int)) }
func (h *minHeap) Pop() interface{} {
	t := *h
	res := t[len(t)-1]
	*h = t[:len(t)-1]
	return res
}

type maxHeap []int

func (h *maxHeap) Len() int             { return len((*h)) }
func (h *maxHeap) Less(i, j int) bool   { return (*h)[i] > (*h)[j] }
func (h *maxHeap) Swap(i, j int)        { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *maxHeap) Push(val interface{}) { *h = append(*h, val.(int)) }
func (h *maxHeap) Pop() interface{} {
	t := *h
	res := t[len(t)-1]
	*h = t[:len(t)-1]
	return res
}

type MedianFinder struct {
	minData *minHeap
	maxData *maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minData: new(minHeap),
		maxData: new(maxHeap),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 相等的时候最终加在小堆，不等的时候加在大堆，保证两边数量一致
	if this.minData.Len() == this.maxData.Len() {
		heap.Push(this.maxData, num)
		heap.Push(this.minData, heap.Pop(this.maxData))
	} else {
		heap.Push(this.minData, num)
		heap.Push(this.maxData, heap.Pop(this.minData))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minData.Len() == 0 && this.maxData.Len() == 0 {
		return 0
	}
	minData, maxData := *(this.minData), *(this.maxData)
	if this.minData.Len() == this.maxData.Len() {
		return float64(minData[0]+maxData[0]) / 2
	} else {
		return float64(minData[0])
	}
}
