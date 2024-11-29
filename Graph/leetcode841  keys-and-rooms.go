package main

import "fmt"

// CanVisitAllRooms rooms本质上就是邻接表，BFS遍历
func CanVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms)) // 标记节点是否访问
	visited[0] = true                   // 起点默认已访问
	count := 1                          // 已访问的节点数
	queue := []int{0}                   // 节点先访问，再入队列
	for len(queue) > 0 {
		// 出队
		cur := queue[0]
		queue = queue[1:]
		// 将未访问的邻接节点入队
		for _, next := range rooms[cur] {
			if !visited[next] {
				visited[next] = true
				count++
				queue = append(queue, next)
			}
		}
	}
	return count == len(rooms) // 所有节点是否都被访问
}

func TestCanVisitAllRooms() {
	fmt.Println(CanVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
