package one

/**
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。

input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
output: true

input: board = [["a","b"],["c","d"]], word = "abcd"
output: true
*/

func Hand12(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	rows, cols := len(board), len(board[0])
	var exist func(i, j, idx int) bool
	exist = func(i, j, idx int) bool {
		if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}

		val := board[i][j]
		board[i][j] = '$'
		res := exist(i-1, j, idx+1) || exist(i+1, j, idx+1) || exist(i, j-1, idx+1) || exist(i, j+1, idx+1)
		board[i][j] = val // 回溯
		return res
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if exist(i, j, 0) {
				return true
			}
		}
	}
	return false
}
