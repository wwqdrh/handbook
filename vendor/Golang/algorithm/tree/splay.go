package tree

import (
	. "fmt"
	"strings"
)

/*
伸展树 splay
https://oi-wiki.org/ds/splay/
https://www.cnblogs.com/cjyyb/p/7499020.html
普通平衡树 https://www.luogu.com.cn/problem/P3369 https://www.luogu.com.cn/problem/P6136
文艺平衡树 https://www.luogu.com.cn/problem/P3391

有关 Link Cut Tree 的部分见 link_cut_tree.go
*/

// 下面的代码参考了刘汝佳的实现，即不使用父节点指针的方案
// 若想看使用父节点指针的方案，可以见 link_cut_tree.go

type spKeyType int
type spValueType int

type spNode struct {
	lr  [2]*spNode
	sz  int
	key spKeyType
	val spValueType
}

// 设置如下返回值是为了方便使用 spNode 中的 lr 数组
func (o *spNode) cmpKth(k int) int {
	switch d := k - o.lr[0].size() - 1; {
	case d < 0:
		return 0 // 左儿子
	case d > 0:
		return 1 // 右儿子
	default:
		return -1
	}
}

func (o *spNode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

// 对于取名叫 maintain 还是 pushUp，由于操作的对象是当前节点，个人认为取名 maintain 更为准确
func (o *spNode) maintain() {
	o.sz = 1 + o.lr[0].size() + o.lr[1].size()
}

func (o *spNode) pushDown() {
	// custom ...

}

// 旋转，并维护子树大小
// d=0：左旋，返回 o 的右儿子
// d=1：右旋，返回 o 的左儿子
func (o *spNode) rotate(d int) *spNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	// x.sz = o.sz; o.maintain()
	o.maintain()
	x.maintain()
	return x
}

// 将子树 o 的第 k 小节点伸展到 o，返回该节点
// k 必须为正
func (o *spNode) splay(k int) (kth *spNode) {
	o.pushDown()
	d := o.cmpKth(k)
	if d < 0 {
		return o
	}
	k -= d * (o.lr[0].size() + 1)
	c := o.lr[d]
	c.pushDown()
	if d2 := c.cmpKth(k); d2 >= 0 {
		c.lr[d2] = c.lr[d2].splay(k - d2*(c.lr[0].size()+1))
		if d2 == d {
			o = o.rotate(d ^ 1)
		} else {
			o.lr[d] = c.rotate(d)
		}
	}
	return o.rotate(d ^ 1)
}
func (o *spNode) splayMin() *spNode { return o.splay(1) }
func (o *spNode) splayMax() *spNode { return o.splay(o.size()) }

// 分裂子树 o，把 o 的前 k 小个节点放在 lo 子树，其他的放在 ro 子树（lo 节点为 o 的第 k 小节点）
// 0 < k <= o.size()，取等号时 ro 为 nil
func (o *spNode) split(k int) (lo, ro *spNode) {
	lo = o.splay(k)
	ro = lo.lr[1]
	lo.lr[1] = nil
	lo.maintain()
	return
}

// 把子树 ro 合并进子树 o，返回合并前 o 的最大节点
// 子树 o 的所有元素比子树 ro 中的小
// o != nil
func (o *spNode) merge(ro *spNode) *spNode {
	// 把最大节点伸展上来，这样会空出一个右儿子用来合并 ro
	o = o.splayMax()
	o.lr[1] = ro
	o.maintain()
	return o
}

type splay struct{ root *spNode }

const (
	splayMin spKeyType = -2e9
	splayMax spKeyType = 2e9
)

func newSplay() *splay {
	// 放入两个哨兵节点 min max，以简化 put delete 的逻辑
	// 注意哨兵对 size() 的影响
	root := &spNode{key: splayMin}      // value: 1
	root.lr[1] = &spNode{key: splayMax} // value: 1
	t := &splay{root}
	t.maintain()
	return t
}

func (t *splay) maintain() {
	t.root.lr[1].maintain()
	t.root.maintain()
}

// <= key 的元素个数
func (t *splay) rank(key spKeyType) (kth int) {
	for o := t.root; o != nil; {
		switch {
		case key < o.key:
			o = o.lr[0]
		case key > o.key:
			kth += 1 + o.lr[0].size()
			o = o.lr[1]
		default:
			kth += 1 + o.lr[0].size()
			return
		}
	}
	return
}

func (t *splay) put(key spKeyType, value spValueType) {
	t.root = t.root.splay(t.rank(key))
	if t.root.key == key {
		t.root.val += value
	} else {
		// 把右子树的最小节点伸展上来，这样它就会空出一个左儿子用来插入
		t.root.lr[1] = t.root.lr[1].splayMin()
		t.root.lr[1].lr[0] = &spNode{sz: 1, key: key, val: value}
	}
	t.maintain()
}

func (t *splay) delete(key spKeyType) {
	t.root = t.root.splay(t.rank(key))
	if t.root.key != key {
		return
	}
	if t.root.val > 1 {
		t.root.val--
	} else {
		// 把右子树的最小节点伸展上来，这样它就会空出一个左儿子用来插入
		t.root.lr[1] = t.root.lr[1].splayMin()
		t.root.lr[1].lr[0] = t.root.lr[0]
		t.root = t.root.lr[1]
	}
	t.root.maintain()
}

// 其余和 BST 有关的方法见 bst.go
// 注意每次调用之前或之后都要执行一下 t.root = t.root.splay(t.rank(key))，以确保均摊复杂度为 O(logn)
// 注意 min max 哨兵对 rank() kth() 等方法的影响

//

func (o *spNode) String() (s string) {
	if o.key == splayMin {
		return "-∞"
	}
	if o.key == splayMax {
		return "+∞"
	}
	//return strconv.Itoa(int(o.key))
	if o.val == 1 {
		s = Sprintf("%v", o.key)
	} else {
		s = Sprintf("%v(%v)", o.key, o.val)
	}
	s += Sprintf("[sz:%d]", o.sz)
	return
}

/* 逆时针旋转 90° 打印这棵树：根节点在最左侧，右子树在上侧，左子树在下侧

效果如下（只打印 key）

Root
│                           ┌── +∞
│                       ┌── 96
│                   ┌── 92
│               ┌── 90
│               │   └── 78
│           ┌── 77
│           │   └── 70
│           │       └── 62
│       ┌── 58
│       │   │   ┌── 55
│       │   └── 53
│       │       └── 51
│       │           └── 48
│   ┌── 47
└── 43
    │       ┌── 40
    │   ┌── 39
    │   │   │           ┌── 37
    │   │   │       ┌── 31
    │   │   │   ┌── 30
    │   │   └── 27
    └── 25
        │   ┌── 17
        └── 10
            │   ┌── 9
            └── 8
                └── -∞

*/
func (o *spNode) draw(treeSB, prefixSB *strings.Builder, isTail bool) {
	prefix := prefixSB.String()
	if o.lr[1] != nil {
		newPrefixSB := &strings.Builder{}
		newPrefixSB.WriteString(prefix)
		if isTail {
			newPrefixSB.WriteString("│   ")
		} else {
			newPrefixSB.WriteString("    ")
		}
		o.lr[1].draw(treeSB, newPrefixSB, false)
	}
	treeSB.WriteString(prefix)
	if isTail {
		treeSB.WriteString("└── ")
	} else {
		treeSB.WriteString("┌── ")
	}
	treeSB.WriteString(o.String())
	treeSB.WriteByte('\n')
	if o.lr[0] != nil {
		newPrefixSB := &strings.Builder{}
		newPrefixSB.WriteString(prefix)
		if isTail {
			newPrefixSB.WriteString("    ")
		} else {
			newPrefixSB.WriteString("│   ")
		}
		o.lr[0].draw(treeSB, newPrefixSB, true)
	}
}

func (t *splay) String() string {
	if t.root == nil {
		return "Empty\n"
	}
	treeSB := &strings.Builder{}
	treeSB.WriteString("Root\n")
	t.root.draw(treeSB, &strings.Builder{}, true)
	return treeSB.String()
}
