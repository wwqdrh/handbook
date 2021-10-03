package three

/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

input: [[1,2,3],[4,5,6],[7,8,9]]
output: [1,2,3,6,9,8,7,4,5]
*/

func Hand29(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
