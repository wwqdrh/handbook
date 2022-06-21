package tree

type RangeModule struct {
	root *Node
}

type Node struct {
	ls, rs   *Node
	sum, add int
}

var N = int(1e9) + 10

func Constructor() RangeModule {
	return RangeModule{
		root: new(Node),
	}
}

// 动态开点线段树模板
func (this *RangeModule) update(node *Node, lc, rc, l, r, v int) {
	length := rc - lc + 1
	if lc >= l && rc <= r {
		if v == 1 {
			node.sum = length
		} else {
			node.sum = 0
		}
		node.add = v
		return
	}
	this.pushdown(node, length)
	mid := (lc + rc) >> 1
	if l <= mid {
		this.update(node.ls, lc, mid, l, r, v)
	}
	if r > mid {
		this.update(node.rs, mid+1, rc, l, r, v)
	}
	this.pushup(node)
}

func (this *RangeModule) pushup(node *Node) {
	node.sum = node.ls.sum + node.rs.sum
}

func (this *RangeModule) pushdown(node *Node, length int) {
	if node.ls == nil {
		node.ls = new(Node)
	}
	if node.rs == nil {
		node.rs = new(Node)
	}
	if node.add == 0 {
		return
	}

	add := node.add
	if add == -1 {
		node.ls.sum = 0
		node.rs.sum = 0
	} else {
		node.ls.sum = length - length/2
		node.rs.sum = length / 2
	}
	node.ls.add = add
	node.rs.add = add
	node.add = 0
}

func (this *RangeModule) query(node *Node, lc, rc, l, r int) int {
	if lc >= l && rc <= r {
		return node.sum
	}
	this.pushdown(node, rc-lc+1)

	mid := (lc + rc) >> 1
	ans := 0
	if l <= mid {
		ans = this.query(node.ls, lc, mid, l, r)
	}
	if r > mid {
		ans += this.query(node.rs, mid+1, rc, l, r)
	}
	return ans
}

// <<动态开点线段树模板

func (this *RangeModule) AddRange(left int, right int) {
	this.update(this.root, 1, N-1, left, right-1, 1)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.query(this.root, 1, N-1, left, right-1) == right-left
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.update(this.root, 1, N-1, left, right-1, -1)
}
