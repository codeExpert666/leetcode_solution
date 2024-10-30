package main

import "fmt"

// 对于每个坐标，也可不封装成[2]int{}，直接将横纵坐标放入队列，只要保证每次出队时一次性出两个元素即可
func getMaxAreaBFS(matrix [][]int, visited [][]bool) int {
	var res int
	bfs := func(startx, starty int) int {
		var area = 1
		neighbors := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
		visited[startx][starty] = true
		queue := [][2]int{{startx, starty}}
		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]
			for _, neighbor := range neighbors {
				newPos := [2]int{pos[0] + neighbor[0], pos[1] + neighbor[1]}
				if newPos[0] < 0 || newPos[0] >= len(matrix) || newPos[1] < 0 || newPos[1] >= len(matrix[0]) ||
					matrix[newPos[0]][newPos[1]] == 0 {
					continue
				}
				if !visited[newPos[0]][newPos[1]] {
					area++
					visited[newPos[0]][newPos[1]] = true
					queue = append(queue, newPos)
				}
			}
		}
		return area
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && !visited[i][j] {
				area := bfs(i, j)
				if area > res {
					res = area
				}
			}
		}
	}
	return res
}

func main() {
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
	res := getMaxAreaBFS(matrix, visited)
	fmt.Println(res)
}
