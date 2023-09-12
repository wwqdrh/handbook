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

// 倒数第k个节点
func (l *LinkedList) LastK(k int) interface{} {
	slow, fast := l.root, l.root
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.next
	}

	for fast != nil {
		slow, fast = slow.next, fast.next
	}

	return slow.val
}

func (l *LinkedList) Reversed() {
	var prev *LinkedListNode

	for cur := l.root; cur != nil; {
		temp := cur.next

		cur.next = prev
		prev = cur
		cur = temp
	}

	l.root = prev
}

func (l *LinkedList) Merge(other *LinkedList) {
	cur := l.root
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = other.root
}
