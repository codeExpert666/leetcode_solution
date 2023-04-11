package BackTracking

/**
 * 回溯
 */

type Position struct {
	row, column int
}

func solveSudoku(board [][]byte) {
	// 先获取所有空白位置
	blankPos := getAllBlankPos(board)

	// 回溯函数
	var backtrack func(int) bool
	backtrack = func(cur int) bool {
		// 终止条件
		if cur == len(blankPos) {
			return true
		}
		curPos := blankPos[cur]
		for i := 1; i < 10; i++ {
			if isNumberValid(i, curPos, board) { // 剪枝
				board[curPos.row][curPos.column] = byte(i + 48)
				if !backtrack(cur + 1) {
					board[curPos.row][curPos.column] = '.' // 回溯
				} else {
					return true
				}
			}
		}
		return false
	}

	backtrack(0)
}

// 获取所有待填写数字的位置
func getAllBlankPos(board [][]byte) []*Position {
	blankPos := make([]*Position, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				// 记录位置
				blankPos = append(blankPos, &Position{
					row:    i,
					column: j,
				})
			}
		}
	}
	return blankPos
}

// 判断在某空白位置填写某数字是否合法
func isNumberValid(n int, pos *Position, board [][]byte) bool {
	// 检查同行是否有重复数字
	for i := 0; i < len(board); i++ {
		if i != pos.column {
			number := board[pos.row][i] - 48
			if int(number) == n {
				return false
			}
		}
	}
	// 检查同列是否有重复数字
	for i := 0; i < len(board); i++ {
		if i != pos.row {
			number := board[i][pos.column] - 48
			if int(number) == n {
				return false
			}
		}
	}
	// 判断该数字所在九宫格的起始结束位置
	//startRow, endRow := getStartEndPos(pos.row)
	//startCol, endCol := getStartEndPos(pos.column)
	// 判断该数字所在九宫格的起始结束位置(优化版本)
	startRow := pos.row / 3 * 3
	startCol := pos.column / 3 * 3
	// 检查九宫格是否有重复数字
	for i := startRow; i <= startRow+2; i++ {
		for j := startCol; j <= startCol+2; j++ {
			if i != pos.row && j != pos.column {
				number := board[i][j] - 48
				if int(number) == n {
					return false
				}
			}
		}
	}
	return true
}

//func getStartEndPos(index int) (int, int) {
//	if index < 3 {
//		return 0, 2
//	} else if index < 6 {
//		return 3, 5
//	} else {
//		return 6, 8
//	}
//}

/**
 * 回溯（方法二）
 * 个人认为这个方法并不如自己的方法，因为这种方法每次递归都需要从头遍历数独，时间复杂度较高
 */
func solveSudoku1(board [][]byte) {
	var backtracking func(board [][]byte) bool
	backtracking = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				//判断此位置是否适合填数字
				if board[i][j] != '.' {
					continue
				}
				//尝试填1-9
				for k := '1'; k <= '9'; k++ {
					if isvalid1(i, j, byte(k), board) == true { //如果满足要求就填
						board[i][j] = byte(k)
						if backtracking(board) == true {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtracking(board)
}

//判断填入数字是否满足要求
func isvalid1(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}
	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
