package one

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

input: 5, output: true
input: 20, output: false
*/

func Hand4(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	type Node [2]int
	m, n := len(matrix), len(matrix[0])

	node := Node{0, n - 1} // 行. 列
	toLeft := func(i *Node) { i[1]-- }
	toRight := func(i *Node) { i[0]++ }
	for node[0] < m && node[1] >= 0 {
		curVal := matrix[node[0]][node[1]]
		if curVal == target {
			return true
		} else if curVal < target {
			toRight(&node)
		} else {
			toLeft(&node)
		}
	}
	return false
}
