package weeklycompetition

func minimumOperations(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	res := 0
	for c := 0; c < col; c++ {
		for r := 1; r < row; r++ {
			if grid[r][c] <= grid[r-1][c] {
				tmp := grid[r][c]
				grid[r][c] = grid[r-1][c] + 1
				res += grid[r][c] - tmp
			}
		}
	}
	return res
}
