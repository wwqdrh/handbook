package sequence

import "errors"

// 使用两个栈实现队列
// 一个栈用来辅助 一个栈来存储数据

type StackQueue struct {
	dataStack []interface{}
	helpStack []interface{}
}

func NewStackQueue() *StackQueue {
	return &StackQueue{
		dataStack: make([]interface{}, 0),
		helpStack: make([]interface{}, 0),
	}
}

func (q *StackQueue) Push(val interface{}) {
	q.dataStack = append(q.dataStack, val)
}

func (q *StackQueue) Pop() (interface{}, error) {
	if len(q.dataStack) == 0 && len(q.helpStack) == 0 {
		return nil, errors.New("queue为空")
	}
	if len(q.helpStack) == 0 {
		for len(q.dataStack) > 0 {
			q.helpStack = append(q.helpStack, q.dataStack[len(q.dataStack)-1])
			q.dataStack = q.dataStack[:len(q.dataStack)-1]
		}
	}
	val := q.helpStack[len(q.helpStack)-1]
	q.helpStack = q.helpStack[:len(q.helpStack)-1]
	return val, nil
}
