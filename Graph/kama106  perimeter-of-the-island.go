package main

import "fmt"

// startx与starty是创建matrix时遇到的第一个陆地位置
func islandPerimeter(matrix [][]int, visited [][]bool, startx, starty int) int {
	var res int
	neighbors := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	bfs := func() {
		visited[startx][starty] = true // 先访问
		queue := []int{startx, starty}
		for len(queue) > 0 {
			x, y := queue[0], queue[1]
			queue = queue[2:]
			for _, neighbor := range neighbors {
				newx, newy := x+neighbor[0], y+neighbor[1]
				if newx < 0 || newx >= len(matrix) || newy < 0 || newy >= len(matrix[0]) ||
					matrix[newx][newy] == 0 { // 陆地在地图边缘或陆地周围有水域则累加周长
					res++
				} else if !visited[newx][newy] {
					visited[newx][newy] = true
					queue = append(queue, newx, newy)
				}
			}
		}
	}
	bfs()
	return res
}

func TsetIslandPerimeter() {
	var row, col int
	fmt.Scan(&row, &col)
	matrix := make([][]int, row)
	visited := make([][]bool, row)
	var startx, starty int // 第一块陆地位置
	var flag bool          // 标记是否找到第一块陆地
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
		matrix[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Scan(&matrix[i][j])
			if !flag && matrix[i][j] == 1 {
				startx, starty = i, j
				flag = true
			}
		}
	}
	fmt.Println(islandPerimeter(matrix, visited, startx, starty))
}

// 本题也可以不用dfs或bfs，甚至不用抽象成图来解决
// 正常按序二维遍历，对每个遇到的陆地计算可用边长数量，也是可以的
