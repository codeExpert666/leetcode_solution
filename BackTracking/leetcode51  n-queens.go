package BackTracking

import "strings"

/**
 * 回溯（棋盘回溯问题）
 */
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)     // 存放最终结果
	chessBoard := make([][]string, n) // 记录棋盘
	// 初始化棋盘
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "."
		}
	}
	var back func(int) // 回溯函数
	back = func(row int) {
		// 终止条件
		if row == n {
			// 将字符串数组转换成题目所需格式
			//tmpBoard := make([]string, n)
			//for i := 0; i < n; i++ {
			//	var tmp string
			//	for j := 0; j < n; j++ {
			//		tmp += chessBoard[i][j]
			//	}
			//	// 这里注意，不要写成append函数，因为已经确定长度，再用append会在切片末尾添加
			//	tmpBoard[i] = tmp
			//}
			// 将字符串数组转换成题目所需格式（法二：利用字符串内置函数）
			tmpBoard := make([]string, n)
			for i, rowStr := range chessBoard {
				tmpBoard[i] = strings.Join(rowStr, "")
			}
			// 添加结果
			result = append(result, tmpBoard)
			return
		}
		// 遍历第row行的所有位置
		for column := 0; column < n; column++ {
			if isValid(chessBoard, row, column) { // 剪枝
				chessBoard[row][column] = "Q"
				back(row + 1)
				chessBoard[row][column] = "." // 回溯
			}
		}
	}
	// 执行回溯函数
	back(0)
	return result
}

// 判断该皇后位置是否安全
func isValid(chessBoard [][]string, row, column int) bool {
	// 检查同列有没有其他皇后
	for i := 0; i < row; i++ {
		if chessBoard[i][column] == "Q" {
			return false
		}
	}
	// 检查左上方是否有其他皇后
	for i, j := row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}
	// 检查右上方是否有其他皇后
	for i, j := row-1, column+1; i >= 0 && j <= len(chessBoard)-1; i, j = i-1, j+1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}
	return true
}
