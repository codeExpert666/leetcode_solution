package Simulation

func FindBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])

	res := make([]int, n)
	for ball := range n { // 遍历每个球
		column := ball         // 小球所在列
		for layer := range m { // 遍历每一层下落
			if grid[layer][column] == 1 {
				column++
				if column == n || grid[layer][column] == -1 {
					column = -1
					break
				}
			} else {
				column--
				if column == -1 || grid[layer][column] == 1 {
					column = -1
					break
				}
			}
		}
		res[ball] = column
	}

	return res
}
