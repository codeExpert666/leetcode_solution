package BackTracking

// 判断当前位置放置棋子是否合理
func isLegal(board [][]bool, row, col int) bool {
	for i := 1; i <= row; i++ { // 因为按行遍历，所以只需判断上方棋子位置即可
		if board[row-i][col] { // 正上方
			return false
		}
		if col-i >= 0 && board[row-i][col-i] { // 左上方
			return false
		}
		if col+i < len(board) && board[row-i][col+i] { // 右上方
			return false
		}
	}
	return true
}

func TotalNQueens(n int) int {
	var res int
	// 构建棋盘
	board := make([][]bool, n)
	for i := 0; i < len(board); i++ {
		board[i] = make([]bool, n)
	}
	var backTracking func(int)
	backTracking = func(row int) {
		if row == n { // 超出棋盘范围，统计当前盘面
			res++
			return
		}
		for col := 0; col < n; col++ { // 遍历列
			if isLegal(board, row, col) {
				board[row][col] = true  // 放置棋子
				backTracking(row + 1)   // 下一行下一棋子
				board[row][col] = false // 回溯
			}
		}
	}
	backTracking(0)
	return res
}
