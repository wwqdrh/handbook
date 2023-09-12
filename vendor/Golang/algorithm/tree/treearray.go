package tree

// 树状数组
// 多次修改单个值，求区间

type TreeArray struct {
	nums, tree []int
}

func NewTreeArray(nums []int) *TreeArray {
	tree := make([]int, len(nums)+1)
	na := &TreeArray{
		nums, tree,
	}
	for i, num := range nums {
		na.add(i+1, num)
	}
	return na
}

func (t *TreeArray) Update(index, val int) {
	t.add(index+1, val-t.nums[index])
	t.nums[index] = val
}

// 爬树修改
func (t *TreeArray) add(index, val int) {
	// 求出最低位的1的位置
	for ; index < len(t.tree); index += index & -index {
		t.tree[index] += val
	}
}

func (t *TreeArray) SumRange(left int, right int) int {
	return t.prefixSum(right+1) - t.prefixSum(left)
}

func (t *TreeArray) prefixSum(index int) (sum int) {
	for ; index > 0; index &= index - 1 {
		sum += t.tree[index]
	}
	return
}
