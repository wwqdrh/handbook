package four

import "container/heap"

/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，
则最小的4个数字是1、2、3、4。
*/

// 建堆Less、Swap、Len、Push、Pop、Peek
type heapInt []int

func (h *heapInt) Less(i, j int) bool { return (*h)[i] > (*h)[j] } // <就是小根堆，>就是大根堆
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *heapInt) Pop() interface{} {
	idx := len(*h) - 1
	t := (*h)[idx]
	*h = (*h)[:idx]
	return t
}
func (h *heapInt) Peek() int {
	return (*h)[0] // 堆顶，最大值
}

func Hand40(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	d := &heapInt{} // 大根堆
	for _, v := range arr {
		if d.Len() < k {
			heap.Push(d, v)
		} else {
			if d.Peek() > v {
				heap.Pop(d) // 删除最大值
				heap.Push(d, v)
			}
		}
	}
	return *d
}
