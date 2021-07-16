package three

import "container/list"

/**
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
调用 min、push 及 pop 的时间复杂度都是 O(1)。

*/

type MinStack struct {
	data   *list.List
	helper *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:   list.New(),
		helper: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	if minEle := this.helper.Back(); minEle == nil || x <= minEle.Value.(int) {
		this.helper.PushBack(x)
	}
	this.data.PushBack(x)
}

func (this *MinStack) Pop() {
	dataTop := this.data.Back()
	if dataTop != nil {
		val := dataTop.Value.(int)
		if helpTop := this.helper.Back(); helpTop != nil && helpTop.Value.(int) == val {
			this.helper.Remove(helpTop)
		}
		this.data.Remove(dataTop)
	}
}

func (this *MinStack) Top() int {
	if dataTop := this.data.Back(); dataTop != nil {
		return dataTop.Value.(int)
	} else {
		return -1
	}
}

func (this *MinStack) Min() int {
	if helpTop := this.helper.Back(); helpTop != nil {
		return helpTop.Value.(int)
	} else {
		return -1
	}
}
