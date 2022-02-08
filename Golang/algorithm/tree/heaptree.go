package tree

import (
	"container/heap"
	"sort"
)

// 贪心+大根堆的应用
// 每次只能选一个 那就选结束最早的最有益
// 然后在选择的组合上，之前想用递归发现超时了
// 使用大根堆，遍历的时候总时间+当前的持续时间<当前的最晚时间就直接把元素加进去
// 如果不是，则比较大根堆中最大的就是[0]是否比当前的大
// 如果比当前的大说明替换掉最大的 这样能留出更多的空间
// 需要用到的api就是heap替换后的fix操作
// 最后优先队列有多长就有多少元素

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := &Heap{}
	total := 0 // 优先队列中所有课程的总时间
	for _, course := range courses {
		if t := course[0]; total+t <= course[1] {
			total += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.IntSlice[0] {
			total += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}
