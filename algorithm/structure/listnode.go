package structure

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
