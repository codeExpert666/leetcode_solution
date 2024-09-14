package Array

func spiralOrder(matrix [][]int) []int {
	row, col := len(matrix), len(matrix[0])
	var loop int // 螺旋的层数
	if row <= col {
		loop = row / 2
	} else {
		loop = col / 2
	}
	i, j := 0, 0                   // 记录遍历到的位置
	offset := 1                    // 每一层的遍历边界偏移量
	res := make([]int, 0, row*col) // 记录结果
	for loop > 0 {
		// 遍历上边
		for ; j < col-offset; j++ {
			res = append(res, matrix[i][j])
		}
		// 遍历右边
		for ; i < row-offset; i++ {
			res = append(res, matrix[i][j])
		}
		// 遍历下边
		for ; j >= offset; j-- {
			res = append(res, matrix[i][j])
		}
		// 遍历左边
		for ; i >= offset; i-- {
			res = append(res, matrix[i][j])
		}
		// 进入下一层，需更新相关值
		i++
		j++
		offset++
		loop--
	}
	// 特殊处理矩阵行列数的较小者为奇数的情况
	if row <= col && row%2 != 0 {
		for ; j <= col-offset; j++ {
			res = append(res, matrix[i][j])
		}
	} else if row > col && col%2 != 0 {
		for ; i <= row-offset; i++ {
			res = append(res, matrix[i][j])
		}
	}
	return res
}
