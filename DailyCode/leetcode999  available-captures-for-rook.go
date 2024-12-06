package DailyCode

func NumRookCaptures(board [][]byte) int {
	// 定义 capture 函数，用于判断车在某个方向上能吃到的卒的数量
	capture := func(x, y, dx, dy int) int {
		for x >= 0 && x < 8 && y >= 0 && y < 8 {
			if board[x][y] == 'B' {
				break
			}
			if board[x][y] == 'p' {
				return 1
			}
			x += dx
			y += dy
		}
		return 0
	}
	// 先找到车的位置
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				// 找到车后，分别向上下左右找
				return capture(i, j, 0, 1) + capture(i, j, 0, -1) + capture(i, j, 1, 0) + capture(i, j, -1, 0)
			}
		}
	}
	return 0 // 棋盘中没有车
}
