package tree

// 推荐阅读《算法竞赛进阶指南》0x43 和 0x48 节
// https://oi-wiki.org/ds/seg/
// https://cp-algorithms.com/data_structures/segment_tree.html
// 总结得比较详细 https://www.acwing.com/blog/content/1684/
// https://codeforces.com/blog/entry/18051
// https://codeforces.com/blog/entry/89313
// https://codeforces.com/blog/entry/15890
// todo 高效线段树 crazySegmentTree https://codeforces.com/blog/entry/89399
//  像使用 STL 一样使用线段树 https://zhuanlan.zhihu.com/p/459679512 https://zhuanlan.zhihu.com/p/459880950

// 注：对于指针写法，必要时禁止 GC，能加速不少
// func init() { debug.SetGCPercent(-1) }

// 最值及其下标 https://codeforces.com/contest/474/problem/E
// 最大子段和 https://codeforces.com/edu/course/2/lesson/4/2/practice/contest/273278/problem/A https://www.acwing.com/problem/content/246/ https://www.luogu.com.cn/problem/P4513
// 最大子段和+按位或 https://www.luogu.com.cn/problem/P7492 (https://www.luogu.com.cn/contest/42328)
// 开方 https://codeforces.com/problemset/problem/920/F https://www.luogu.com.cn/problem/P4145 http://acm.hdu.edu.cn/showproblem.php?pid=4027
// 取模 https://codeforces.com/problemset/problem/438/D
// 转换的好题 https://codeforces.com/problemset/problem/1187/D
// 合并 https://codeforces.com/problemset/problem/380/C
// k 维曼哈顿（单点修改+区间最大值）https://codeforces.com/problemset/problem/1093/G
// 区间 mex https://www.luogu.com.cn/problem/P4137
// - 做法之一是离线+线段树二分 https://www.luogu.com.cn/blog/user7035/solution-p4137
// - 也可以用树状数组 https://www.luogu.com.cn/blog/Atalod/ti-xie-p4137-post
// - 反向构造题 https://www.luogu.com.cn/problem/P6852
// 区间（绝对）众数及其次数（摩尔投票算法）https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
//     LC1157 https://leetcode-cn.com/problems/online-majority-element-in-subarray/
//     https://www.luogu.com.cn/problem/P3567
//     https://www.luogu.com.cn/problem/P3765
//     https://codeforces.com/contest/1514/problem/D
// GCD https://codeforces.com/problemset/problem/914/D
// 最小差值 https://codeforces.com/problemset/problem/765/F
// 题目推荐 https://cp-algorithms.com/data_structures/segment_tree.html#toc-tgt-12
// LC https://leetcode-cn.com/tag/segment-tree/
// 另见 dp.go 的动态 DP 部分
// todo http://poj.org/problem?id=2991
// 变换成值域 https://www.luogu.com.cn/problem/SP1684 https://www.luogu.com.cn/problem/UVA11235 http://poj.org/problem?id=3368
// http://poj.org/problem?id=3470
// todo http://poj.org/problem?id=1201

// EXTRA: 权值线段树
// 讲解与习题 https://www.luogu.com.cn/blog/bfqaq/qian-tan-quan-zhi-xian-duan-shu
// 浅谈权值线段树到主席树 https://www.luogu.com.cn/blog/your-alpha1022/WeightSegmentTree-ChairmanTree
// 谈树状数组套权值树 https://www.luogu.com.cn/blog/bfqaq/qian-tan-shu-zhuang-shuo-zu-quan-zhi-shu

// EXTRA: 线段树优化建图
// 每个位置对应着 O(logn) 个线段树上的节点，每个区间可以拆分成至多 O(logn) 个线段树上的区间
// 这个性质可以用于优化建图
// https://www.luogu.com.cn/blog/forever-captain/DS-optimize-graph
// https://codeforces.com/problemset/problem/786/B
// todo https://www.luogu.com.cn/problem/P6348

// l 和 r 也可以写到方法参数上，实测二者在执行效率上无异
// 考虑到 debug 和 bug free 上的优点，写到结构体参数中
type seg []struct {
	l, r int
	val  int
}

// 单点更新：build 和 update 通用
func (t seg) set(o, val int) {
	t[o].val = val
}

// 合并两个节点上的数据：maintain 和 query 通用
// 要求操作满足区间可加性
// 例如 + * | & ^ min max gcd mulMatrix 摩尔投票 最大子段和 ...
func (seg) op(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].val = t.op(lo.val, ro.val)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// o=1  1<=i<=n
func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	if m := (t[o].l + t[o].r) >> 1; i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.op(vl, vr)
}

func (t seg) queryAll() int { return t[1].val }

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

// EXTRA: 查询整个区间小于 v 的最靠左的位置
// 这里线段树维护的是区间最小值
// 需要先判断 t[1].min < v
func (t seg) queryFirstLessPos(o, v int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	if t[o<<1].val < v {
		return t.queryFirstLessPos(o<<1, v)
	}
	return t.queryFirstLessPos(o<<1|1, v)
}

// EXTRA: 查询 [l,r] 上小于 v 的最靠左的位置
// 这里线段树维护的是区间最小值
// 不存在时返回 0
func (t seg) queryFirstLessPosInRange(o, l, r, v int) int {
	if t[o].val >= v {
		return 0
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		if pos := t.queryFirstLessPosInRange(o<<1, l, r, v); pos > 0 {
			return pos
		}
	}
	if m < r {
		if pos := t.queryFirstLessPosInRange(o<<1|1, l, r, v); pos > 0 { // 注：这里 pos > 0 的判断可以省略，因为 pos == 0 时最后仍然会返回 0
			return pos
		}
	}
	return 0
}

//

// 延迟标记（区间修改）
// 单个更新操作：
// + min/max https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/A
//           https://codeforces.com/problemset/problem/1321/E
//           https://codeforces.com/problemset/problem/52/C
// + min/max 转换 https://codeforces.com/gym/294041/problem/E
// + max DP https://atcoder.jp/contests/dp/tasks/dp_w
// + ∑ https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/D https://www.luogu.com.cn/problem/P3372
// | & https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/C
// = min https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/E
// = ∑ https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/F https://codeforces.com/problemset/problem/558/E
// max max 离散化 https://codeforces.com/contest/1557/problem/D
// https://codeforces.com/problemset/problem/1114/F
// + 某个区间的不小于 x 的最小下标 https://codeforces.com/edu/course/2/lesson/5/3/practice/contest/280799/problem/C
// =max 求和的 O(log^2) 性质 https://codeforces.com/contest/1439/problem/C
// 矩阵乘法 ∑ https://codeforces.com/problemset/problem/718/C
//
// 多个更新操作复合：
// * + ∑ https://www.luogu.com.cn/problem/P3373 https://leetcode-cn.com/problems/fancy-sequence/
// = + ∑ https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/A
//
// 吉老师线段树 Segment Tree Beats
// todo https://oi-wiki.org/ds/seg-beats/
//  https://codeforces.com/blog/entry/57319
//  区间最值操作与区间历史最值详解 https://www.luogu.com.cn/blog/Hakurei-Reimu/seg-beats
//  模板题 https://www.luogu.com.cn/problem/P6242
//
// EXTRA: 多项式更新 Competitive Programmer’s Handbook Ch.28
// 比如区间加等差数列（差分法）https://www.luogu.com.cn/problem/P1438 https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/B
type lazyST []struct {
	l, r int
	todo int64
	sum  int64
}

func (lazyST) op(a, b int64) int64 {
	return a + b // % mod
}

func (t lazyST) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = t.op(lo.sum, ro.sum)
}

func (t lazyST) build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazyST) do(o int, add int64) {
	to := &t[o]
	to.todo += add                     // % mod
	to.sum += int64(to.r-to.l+1) * add // % mod
}

func (t lazyST) spread(o int) {
	if add := t[o].todo; add != 0 {
		t.do(o<<1, add)
		t.do(o<<1|1, add)
		t[o].todo = 0
	}
}

// 如果维护的数据（或者判断条件）具有单调性，我们就可以在线段树上二分
// 未找到时返回 n+1
// o=1  [l,r] 1<=l<=r<=n
// https://codeforces.com/problemset/problem/1179/C
func (t lazyST) binarySearch(o, l, r int, val int64) int {
	if t[o].l == t[o].r {
		if val <= t[o].sum {
			return t[o].l
		}
		return t[o].l + 1
	}
	t.spread(o)
	// 注意判断对象是当前节点还是子节点
	if val <= t[o<<1].sum {
		return t.binarySearch(o<<1, l, r, val)
	}
	return t.binarySearch(o<<1|1, l, r, val)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) update(o, l, r int, add int64) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, add)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, add)
	}
	if m < r {
		t.update(o<<1|1, l, r, add)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.op(vl, vr)
}

func (t lazyST) queryAll() int64 { return t[1].sum }

// a 从 0 开始
func newLazySegmentTree(a []int64) lazyST {
	t := make(lazyST, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

// EXTRA: 适用于需要提取所有元素值的场景
func (t lazyST) spreadAll(o int) {
	if t[o].l == t[o].r {
		return
	}
	t.spread(o)
	t.spreadAll(o << 1)
	t.spreadAll(o<<1 | 1)
}

//

// 动态开点线段树·其一·单点修改
// LC327 https://leetcode-cn.com/problems/count-of-range-sum/
// rt := &stNode{l: 1, r: 1e9}
type stNode struct {
	lo, ro *stNode
	l, r   int
	sum    int64
}

func (o *stNode) get() int64 {
	if o != nil {
		return o.sum
	}
	return 0 // inf
}

func (stNode) op(a, b int64) int64 {
	return a + b //
}

func (o *stNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

func (o *stNode) build(a []int64, l, r int) {
	o.l, o.r = l, r
	if l == r {
		o.sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	o.lo = &stNode{}
	o.lo.build(a, l, m)
	o.ro = &stNode{}
	o.ro.build(a, m+1, r)
	o.maintain()
}

func (o *stNode) update(i int, add int64) {
	if o.l == o.r {
		o.sum += add //
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == nil {
			o.lo = &stNode{l: o.l, r: m}
		}
		o.lo.update(i, add)
	} else {
		if o.ro == nil {
			o.ro = &stNode{l: m + 1, r: o.r}
		}
		o.ro.update(i, add)
	}
	o.maintain()
}

func (o *stNode) query(l, r int) int64 {
	if o == nil || l > o.r || r < o.l {
		return 0 // inf
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}

// 动态开点线段树·其二·延迟标记（区间修改）
// https://codeforces.com/problemset/problem/915/E（注：此题有多种解法）
// https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/F https://www.luogu.com.cn/problem/P5848
//（内存受限）https://codeforces.com/problemset/problem/1557/D
// rt := &lazyNode{l: 1, r: 1e9}
type lazyNode struct {
	lo, ro *lazyNode
	l, r   int
	sum    int64
	todo   int64
}

func (o *lazyNode) get() int64 {
	if o != nil {
		return o.sum
	}
	return 0 // inf
}

func (lazyNode) op(a, b int64) int64 {
	return a + b //
}

func (o *lazyNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

func (o *lazyNode) build(a []int64, l, r int) {
	o.l, o.r = l, r
	if l == r {
		o.sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	o.lo = &lazyNode{}
	o.lo.build(a, l, m)
	o.ro = &lazyNode{}
	o.ro.build(a, m+1, r)
	o.maintain()
}

func (o *lazyNode) do(add int64) {
	o.todo += add                   // % mod
	o.sum += int64(o.r-o.l+1) * add // % mod
}

func (o *lazyNode) spread() {
	m := (o.l + o.r) >> 1
	if o.lo == nil {
		o.lo = &lazyNode{l: o.l, r: m}
	}
	if o.ro == nil {
		o.ro = &lazyNode{l: m + 1, r: o.r}
	}
	if add := o.todo; add != 0 {
		o.lo.do(add)
		o.ro.do(add)
		o.todo = 0 // -1
	}
}

func (o *lazyNode) update(l, r int, add int64) {
	if l <= o.l && o.r <= r {
		o.do(add)
		return
	}
	o.spread()
	m := (o.l + o.r) >> 1
	if l <= m {
		o.lo.update(l, r, add)
	}
	if m < r {
		o.ro.update(l, r, add)
	}
	o.maintain()
}

func (o *lazyNode) query(l, r int) int64 {
	// 对于不在线段树中的点，应按照题意来返回
	if o == nil || l > o.r || r < o.l {
		return 0 // inf
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	o.spread()
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}

// EXTRA: 线段树合并
// https://www.luogu.com.cn/problem/P5494
// todo 一些题目 https://www.luogu.com.cn/blog/styx-ferryman/xian-duan-shu-ge-bing-zong-ru-men-dao-fang-qi
//   https://codeforces.com/blog/entry/83969
//   https://www.luogu.com.cn/problem/P4556
//   https://www.luogu.com.cn/problem/P5298
//   https://codeforces.com/problemset/problem/600/E
// rt = rt.merge(rt2)
func (o *stNode) merge(b *stNode) *stNode {
	if o == nil {
		return b
	}
	if b == nil {
		return o
	}
	if o.l == o.r {
		// 按照所需合并，如加法
		o.sum += b.sum
		return o
	}
	o.lo = o.lo.merge(b.lo)
	o.ro = o.ro.merge(b.ro)
	o.maintain()
	return o
}

// EXTRA: 线段树分裂
// 将区间 [l,r] 从 o 中分离到 b 上
// https://www.luogu.com.cn/blog/cyffff/talk-about-segument-trees-split
// https://www.luogu.com.cn/problem/P5494
// rt, rt2 := rt.split(nil, l, r)
func (o *stNode) split(b *stNode, l, r int) (*stNode, *stNode) {
	if o == nil || l > o.r || r < o.l {
		return o, nil
	}
	if l <= o.l && o.r <= r {
		return nil, o
	}
	if b == nil {
		b = &stNode{l: o.l, r: o.r}
	}
	o.lo, b.lo = o.lo.split(b.lo, l, r)
	o.ro, b.ro = o.ro.split(b.ro, l, r)
	o.maintain()
	b.maintain()
	return o, b
}

// 权值线段树求第 k 小
// 调用前需保证 1 <= k <= rt.get()
func (o *stNode) kth(k int64) int {
	if o.l == o.r {
		return o.l
	}
	if cntL := o.lo.get(); k <= cntL {
		return o.lo.kth(k)
	} else {
		return o.ro.kth(k - cntL)
	}
}

//

// 可持久化线段树（又称函数式线段树、主席树） Persistent Segment Tree
// https://oi-wiki.org/ds/persistent-seg/
// 静态+动态 https://www.acwing.com/blog/content/487/
// https://zhuanlan.zhihu.com/p/250565583
// https://blog.csdn.net/weixin_43914593/article/details/108861279
//
// 另见 union_find.go 中的「可持久化并查集」
//
// 模板题 https://www.luogu.com.cn/problem/P3919
//       https://www.luogu.com.cn/problem/P3834 https://www.acwing.com/problem/content/257/ https://ac.nowcoder.com/acm/contest/7613/C
// 二分，转换成找最长的已填入数字的区间，做法类似最大子段和 https://codeforces.com/problemset/problem/484/E
// 与 DFS序+深度 结合 https://codeforces.com/problemset/problem/893/F
// todo 种类数 https://codeforces.com/problemset/problem/620/E
//  https://codeforces.com/problemset/problem/786/C
//  差分 https://codeforces.com/problemset/problem/813/E
//  https://codeforces.com/problemset/problem/837/G
//  https://codeforces.com/problemset/problem/840/D
//  https://codeforces.com/problemset/problem/893/F
//  https://codeforces.com/problemset/problem/961/E（不止一种做法）
// 在线做法 https://codeforces.com/problemset/problem/1262/D2
type pstNode struct {
	lo, ro *pstNode
	l, r   int
	sum    int64
}

func (pstNode) op(a, b int64) int64 {
	return a + b //
}

func (o *pstNode) maintain() {
	o.sum = o.op(o.lo.sum, o.ro.sum)
}

// t := make([]*pstNode, 1, maxVersion+1)
// t[0] = buildPST(a, 1, len(a))
// 或者 t := []*pstNode{buildPST(a, 1, len(a))}
func buildPST(a []int64, l, r int) *pstNode {
	o := &pstNode{l: l, r: r}
	if l == r {
		o.sum = a[l-1]
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPST(a, l, m)
	o.ro = buildPST(a, m+1, r)
	o.maintain()
	return o
}

// 一般写法是更新到当前版本，然后把返回的新版本加在 t 的末尾，即
// t = append(t, t[len(t)-1].update(i, add))
// 注意为了拷贝一份 pstNode，这里的接收器不是指针
func (o pstNode) update(i int, add int64) *pstNode {
	if o.l == o.r {
		o.sum += add
		return &o
	}
	if m := o.lo.r; i <= m {
		o.lo = o.lo.update(i, add)
	} else {
		o.ro = o.ro.update(i, add)
	}
	o.maintain()
	return &o
}

func (o *pstNode) query(l, r int) int64 {
	if l <= o.l && o.r <= r {
		return o.sum
	}
	m := o.lo.r
	if r <= m {
		return o.lo.query(l, r)
	}
	if m < l {
		return o.ro.query(l, r)
	}
	vl := o.lo.query(l, r)
	vr := o.ro.query(l, r)
	return o.op(vl, vr)
}

// 主席树相当于对数组的每个前缀建立一颗线段树
// 离散化时，求 kth 需要将相同元素也视作不同的

// EXTRA: 查询区间 [l,r] 中第 k 小在整个数组上的名次（从 1 开始）
// 注意返回的是（排序去重后的数组的）下标，不是元素值
// 初始 t[0] = buildPST(1, len(a))
//     t[i+1] = t[i].update(kth[i], 1)   kth[i] 为 a[i] 离散化后的值（从 1 开始）
// 查询 t[r].kth(t[l-1], k)               类似前缀和 [l,r] 1<=l<=r<=n
func (o *pstNode) kth(old *pstNode, k int) int {
	if o.l == o.r {
		return o.l
	}
	if cntL := o.lo.sum - old.lo.sum; k <= int(cntL) {
		return o.lo.kth(old.lo, k)
	} else {
		return o.ro.kth(old.ro, k-int(cntL))
	}
}

// EXTRA: 查询区间 [l,r] 中在 [low,high] 范围内的元素个数
// low 和 high 为离散化后的值（从 1 开始）
// http://acm.hdu.edu.cn/showproblem.php?pid=4417
// https://codeforces.com/problemset/problem/538/F
// 调用方式同上 t[r].countRange(t[l-1], low, high)
func (o *pstNode) countRange(old *pstNode, low, high int) int {
	if high < o.l || o.r < low {
		return 0
	}
	if low <= o.l && o.r <= high {
		return int(o.sum - old.sum)
	}
	m := o.lo.r
	if high <= m {
		return o.lo.countRange(old.lo, low, high)
	}
	if m < low {
		return o.ro.countRange(old.ro, low, high)
	}
	return o.lo.countRange(old.lo, low, high) + o.ro.countRange(old.ro, low, high)
}

// EXTRA: countDiff 区间不同元素个数
// 做法是维护左侧相同元素的位置
// todo 参考 https://www.luogu.com.cn/problem/solution/SP3267
// https://www.luogu.com.cn/problem/P1972
// https://www.luogu.com.cn/problem/SP3267

// EXTRA: 区间绝对众数及其出现次数（注意返回的众数是离散化后的值）
// 没找到时返回的 mode = -1
// 题目链接见最上
// 调用方式 t[r].countMode(t[l-1], k)，取 k = (r-l+2)/2，即区间长度一半的上取整
// 这种 k 值使得递归时左右儿子至多调用其中一个
func (o *pstNode) countMode(old *pstNode, k int) (mode, count int) {
	if o.l == o.r {
		return o.l, int(o.sum - old.sum)
	}
	if int(o.lo.sum-old.lo.sum) >= k {
		return o.lo.countMode(old.lo, k)
	}
	if int(o.ro.sum-old.ro.sum) >= k {
		return o.ro.countMode(old.ro, k)
	}
	return -1, 0
}

// EXTRA: 子序列自动机 Sequence Automation
// https://www.luogu.com.cn/problem/P5826

// EXTRA: 标记永久化
// todo 一本通 p239

//

// 自底向上（zkw）线段树
// todo  张昆玮《统计的力量》
//   https://zhuanlan.zhihu.com/p/361935620
//   https://zhuanlan.zhihu.com/p/29876526
//   https://zhuanlan.zhihu.com/p/29937723
//   https://codeforces.com/blog/entry/18051 Efficient and easy segment trees

//

// 李超线段树
// todo https://oi-wiki.org/ds/li-chao-tree/
//   https://zhuanlan.zhihu.com/p/64946571
//   https://www.luogu.com.cn/blog/fzber0103/Li-Chao-Tree

//

// Wavelet Trees
// todo https://www.geeksforgeeks.org/wavelet-trees-introduction/
// https://codeforces.com/blog/entry/52854
// https://www.youtube.com/watch?v=4aSv9PcecDw
