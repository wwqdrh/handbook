package one

/**
用两个栈实现一个队列。
队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。
(若队列中没有元素，deleteHead 操作返回 -1 )

input: ["CQueue","appendTail","deleteHead","deleteHead"] [[],[3],[],[]]
output: [null,null,3,-1]
*/
import "container/list"

type CQueue struct {
	inEnd, outEnd *list.List
}

func Constructor() CQueue {
	return CQueue{
		inEnd:  list.New(),
		outEnd: list.New(),
	}
}

func (q *CQueue) AppendTail(value int) {
	q.inEnd.PushBack(value)
}

func (q *CQueue) DeleteHead() int {
	if q.outEnd.Len() == 0 {
		for q.inEnd.Len() > 0 {
			q.outEnd.PushBack(q.inEnd.Remove(q.inEnd.Back()))
		}
	}
	if q.outEnd.Len() != 0 {
		e := q.outEnd.Back()
		q.outEnd.Remove(e)
		return e.Value.(int)
	}
	return -1
}
