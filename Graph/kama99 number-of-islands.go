package main

import "fmt"

// 遍历地图，遇到未访问的陆地，就利用BFS找到整个岛屿
// func islandNumsBFS(matrix [][]int, visited [][]bool) int {
// 	var res int
// 	// BFS
// 	bfs := func(startx, starty int) {
// 		neighbors := [4][2]int{
// 			{-1, 0}, {0, 1}, {1, 0}, {0, -1},
// 		}
// 		queue := [][2]int{{startx, starty}}
// 		visited[startx][starty] = true
// 		for len(queue) > 0 {
// 			// 出队
// 			pos := queue[0]
// 			queue = queue[1:]
// 			// 入队
// 			for _, neighbor := range neighbors {
// 				newPos := [2]int{pos[0] + neighbor[0], pos[1] + neighbor[1]}
// 				if newPos[0] < 0 || newPos[0] >= len(matrix) || newPos[1] < 0 || newPos[1] >= len(matrix[0]) ||
// 					matrix[newPos[0]][newPos[1]] == 0 {
// 					continue // 超出地图边界或超出陆地范围
// 				}
// 				if !visited[newPos[0]][newPos[1]] {
// 					queue = append(queue, newPos)
// 					visited[newPos[0]][newPos[1]] = true
// 				}
// 			}
// 		}
// 	}
// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[i]); j++ {
// 			if matrix[i][j] == 1 && !visited[i][j] {
// 				res++
// 				// BFS遍历标记整片岛屿
// 				bfs(i, j)
// 			}
// 		}
// 	}
// 	return res
// }

// 遍历地图，遇到未访问的陆地，就利用DFS找到整个岛屿
func islandNumsDFS(matrix [][]int, visited [][]bool) int {
	var res int
	neighbors := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dfs func(int, int)
	dfs = func(startx, starty int) {
		visited[startx][starty] = true
		for _, neighbor := range neighbors {
			newx, newy := startx+neighbor[0], starty+neighbor[1]
			if newx < 0 || newx >= len(matrix) || newy < 0 || newy >= len(matrix[0]) ||
				matrix[newx][newy] == 0 {
				continue
			}
			if !visited[newx][newy] {
				dfs(newx, newy)
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 && !visited[i][j] {
				res++
				// DFS遍历标记整片岛屿
				dfs(i, j)
			}
		}
	}
	return res
}

func TestBFS() {
	// 输入
	var row, col int
	fmt.Scan(&row, &col)
	matrix := make([][]int, row)
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
		matrix[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	// 运算
	res := islandNumsDFS(matrix, visited)
	// 输出
	fmt.Println(res)
}
