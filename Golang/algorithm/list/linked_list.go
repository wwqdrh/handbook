package list

// 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	var (
		root *ListNode
		prev *ListNode
	)
	for _, num := range nums {
		cur := &ListNode{Val: num}
		if root == nil {
			root, prev = cur, cur
		} else {
			prev.Next = cur
			prev = cur
		}
	}
	return root
}

func (l *ListNode) ToSlice() (res []int) {
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return
}

// 随机节点
type ListRandNode struct {
	Val    int
	Next   *ListRandNode
	Random *ListRandNode
}
