package structure

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(nums []string) *TreeNode {
	if len(nums) == 0 || nums[0] == "null" {
		return new(TreeNode)
	}
	queue := list.New()
	val, err := strconv.Atoi(nums[0])
	if err != nil {
		return new(TreeNode)
	}

	root := &TreeNode{Val: val}
	queue.PushBack(root)
	idx, length := 1, len(nums)
	for idx < length {
		cur := queue.Front()

		if nums[idx] != "null" {
			if val, err := strconv.Atoi(nums[idx]); err == nil {
				t := &TreeNode{Val: val}
				cur.Value.(*TreeNode).Left = t
				queue.PushBack(t)
			}
		}
		idx++
		if idx == length {
			break
		}

		if nums[idx] != "null" {
			if val, err := strconv.Atoi(nums[idx]); err == nil {
				t := &TreeNode{Val: val}
				cur.Value.(*TreeNode).Right = t
				queue.PushBack(t)
			}
		}
		idx++

		queue.Remove(cur)
	}
	return root
}

func (t *TreeNode) ToString() string {
	res := make([]string, 0)

	queue := list.New()
	queue.PushBack(t)
	for queue.Len() > 0 {
		cur := queue.Front()

		if curVal := cur.Value.(*TreeNode); curVal != nil {
			res = append(res, strconv.FormatInt(int64(curVal.Val), 10))
			queue.PushBack(curVal.Left)
			queue.PushBack(curVal.Right)
		} else {
			res = append(res, "null")
		}

		queue.Remove(cur)
	}
	idx := len(res) - 1
	for ; idx >= 0; idx-- {
		if res[idx] == "null" {
			idx--
		} else {
			break
		}
	}
	return strings.Join(res[:idx+1], ",")
}
