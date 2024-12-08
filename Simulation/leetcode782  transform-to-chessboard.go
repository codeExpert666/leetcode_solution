package Simulation

// Claude 生成的，看不懂，直接 Pass 了，牵涉到很多数学论证和推理，实际不可能考这种题目的
func MovesToChessboard(board [][]int) int {
	n := len(board)
	// 第一步：检查矩阵是否合法
	// 统计第一行和第一列的1的数量
	rowSum, colSum := 0, 0
	// 统计行和列的差异数
	rowDiff, colDiff := 0, 0

	// 检查每一行是否都与第一行相同或完全相反
	for i := 0; i < n; i++ {
		rowSum += board[0][i]
		colSum += board[i][0]

		for j := 0; j < n; j++ {
			// 使用异或运算(^)检查2x2子矩阵的合法性
			if board[0][0]^board[0][j]^board[i][0]^board[i][j] == 1 {
				return -1
			}
		}
	}

	// 检查1和0的数量是否合法
	if rowSum != n/2 && rowSum != (n+1)/2 {
		return -1
	}
	if colSum != n/2 && colSum != (n+1)/2 {
		return -1
	}

	// 计算每一行和标准棋盘的差异
	for i := 0; i < n; i++ {
		if board[0][i] == i%2 {
			rowDiff++
		}
		if board[i][0] == i%2 {
			colDiff++
		}
	}

	// 如果n是奇数
	if n%2 == 1 {
		if rowDiff%2 == 1 {
			rowDiff = n - rowDiff
		}
		if colDiff%2 == 1 {
			colDiff = n - colDiff
		}
	} else {
		// 如果n是偶数，取最小值
		rowDiff = min(rowDiff, n-rowDiff)
		colDiff = min(colDiff, n-colDiff)
	}

	// 返回所需的最小移动次数
	moves := (rowDiff + colDiff) / 2
	if moves*2 != rowDiff+colDiff {
		return -1
	}
	return moves
}
