package graph

import (
	"math"
	"sort"
)

////////////////////
// 最小生成树
////////////////////

// 给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。

// 连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。

// 请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。

type unionFind struct {
	parent, rank []int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{parent, rank}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.parent[fy] = fx
	return true
}

func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

func KruskalMinCostConnectPoints(points [][]int) (ans int) {
	n := len(points)
	type edge struct{ v, w, dis int }
	edges := []edge{}
	for i, p := range points {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(p, points[j])})
		}
	}

	sort.Slice(edges, func(i, j int) bool { return edges[i].dis < edges[j].dis })

	uf := newUnionFind(n)
	left := n - 1
	for _, e := range edges {
		if uf.union(e.v, e.w) {
			ans += e.dis
			left--
			if left == 0 {
				break
			}
		}
	}
	return
}

// Prim算法思想：按点来生成最小生成树，先划分两个点集，已访问过的点集S和为访问过的点集V。
// 先随机选择一个起始点加入已访问过的点集S中，访问未访问过且和该点存在路径的点
// ，存下这些可达点（未访问过，与已访问过的点集中的点存在路径）与已访问过的点集的距离（换种说法就是连通两个点集的边），选择最短距离的边的未访问过的顶点，加入已访问过的点集中，直到把所有点都加入到已访问过的点集中。
func PrimMinCostConnectPoints(points [][]int) int {
	n := len(points)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
		}
	} //邻接矩阵
	visited := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = math.MaxInt32
	}
	d[0] = 0
	res := 0
	for {
		minCost := math.MaxInt32
		u := -1
		for i := 0; i < n; i++ {
			if visited[i] == 0 && d[i] < minCost { //寻找最短路径对应的点及路径长度
				minCost = d[i]
				u = i
			}
		}
		if u == -1 { //点都加入了，退出循环
			break
		}
		res += minCost
		visited[u] = 1
		for i := 0; i < n; i++ {
			if visited[i] == 0 && grid[u][i] > 0 {
				if grid[u][i] < d[i] {
					d[i] = grid[u][i] //更新距离数组
				}
			}
		}
	}
	return res
}
