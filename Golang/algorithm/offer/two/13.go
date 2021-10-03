package two

import "container/list"

/**
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

input: m = 2, n = 3, k = 1
output: 3
*/

func Hand13(m, n, k int) int {
	queue := list.New()
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	queue.PushBack([]int{0, 0, 0, 0})
	res := 0
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		bfs := front.Value.([]int)
		i, j, si, sj := bfs[0], bfs[1], bfs[2], bfs[3]
		if i >= m || j >= n || si+sj > k || visited[i][j] {
			continue
		}
		res++
		visited[i][j] = true
		// sj1, si1 := sj + 1, si + 1
		var sj1, si1 int
		if (j+1)%10 == 0 {
			sj1 = sj - 8
		} else {
			sj1 = sj + 1
		}
		if (i+1)%10 == 0 {
			si1 = si - 8
		} else {
			si1 = si + 1
		}
		queue.PushBack([]int{i + 1, j, si1, sj})
		queue.PushBack([]int{i, j + 1, si, sj1})
	}
	return res
}
