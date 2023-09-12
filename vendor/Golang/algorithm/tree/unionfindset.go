package tree

type UnionFind struct {
	g      []int
	cnt    []int
	groups int
	maxG   int
}

func InitialUF(n int) *UnionFind {
	uf := &UnionFind{groups: n, maxG: 1}
	uf.g = make([]int, n)
	uf.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		uf.g[i] = i
		uf.cnt[i] = 1
	}
	return uf
}

func (u UnionFind) GetMaxG() int {
	return u.maxG
}

func (u *UnionFind) Find(x int) int {
	if u.g[x] == x {
		return x
	}
	u.g[x] = u.Find(u.g[x])
	return u.g[x]
}

func (u *UnionFind) Union(x, y int) bool {
	xp := u.Find(x)
	yp := u.Find(y)
	if xp == yp {
		return false
	}
	if u.cnt[yp] > u.cnt[xp] {
		xp, yp = yp, xp
	}
	u.cnt[xp] += u.cnt[yp]
	u.maxG = max(u.maxG, u.cnt[xp])
	u.cnt[yp] = 0
	u.g[yp] = xp
	u.groups--
	return true
}
