package sequence

type LinkedList struct {
	root *LinkedListNode
}

type LinkedListNode struct {
	val  interface{}
	next *LinkedListNode
}

func NewLinkedList(nums []interface{}) *LinkedList {
	root := new(LinkedListNode)
	preNode := root
	for _, num := range nums {
		curNode := &LinkedListNode{
			val: num,
		}
		preNode.next = curNode
		preNode = curNode
	}
	return &LinkedList{
		root: root.next,
	}
}

func (l *LinkedList) ToSlice() []interface{} {
	res := make([]interface{}, 0)

	for cur := l.root; cur != nil; cur = cur.next {
		res = append(res, cur.val)
	}

	return res
}
