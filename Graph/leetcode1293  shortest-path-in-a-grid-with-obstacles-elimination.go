package main

import "fmt"

type status struct {
	x, y int
	rest int // 剩余还可经过的障碍物数量
}

// ShortestPath BFS扩展题：常规 BFS 求最短路径时，队列中维护了下一步要访问的位置，同时还可以维护源点到该位置的距离信息
// 既然可以多维护一维信息，那就可以再维护任意维信息。也即，维护“还可以经过的障碍物数量”信息
func ShortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if k >= m+n-3 { // 在没有障碍物的情况下，最短路径为 m+n-2，如果可以经过的障碍物数量大于最短路径中间节点数量，直接走最短路径即可
		return m + n - 2
	}
	// BFS 加一维状态
	neighbors := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	// visit设计特殊，这样的设计能够让访问过的节点再被访问，但同时也能过滤掉无效重复，主要是把 rest 当做一种状态
	visited := make(map[status]struct{})
	visited[status{0, 0, k}] = struct{}{} // 先访问，再入队，保证相同状态不会重复入队
	queue := []status{{0, 0, k}}
	res := -1 // 记录路径长度
	for l := len(queue); l > 0; l = len(queue) {
		res++
		for ; l > 0; l-- {
			cur := queue[0]
			queue = queue[1:]
			if cur.x == m-1 && cur.y == n-1 { // 找到最短路径
				return res
			}
			for _, neighbor := range neighbors { // 访问邻居
				nStatus := status{cur.x + neighbor[0], cur.y + neighbor[1], cur.rest}
				if nStatus.x < 0 || nStatus.x >= m || nStatus.y < 0 || nStatus.y >= n { // 超出地图边界
					continue
				}
				if grid[nStatus.x][nStatus.y] == 1 { // 经过障碍物
					nStatus.rest--
				}
				if _, exist := visited[nStatus]; exist || nStatus.rest < 0 {
					continue // 已访问过的节点状态不再访问，无法再经过障碍物时遇到障碍物不再访问
				}
				visited[nStatus] = struct{}{} // 访问
				queue = append(queue, nStatus)
			}
		}
	}
	return -1 // 无法到达终点
}

func main() {
	grid := [][]int{{0, 1, 0, 0, 0, 1, 0, 0}, {0, 1, 0, 1, 0, 1, 0, 1}, {0, 0, 0, 1, 0, 0, 1, 0}}
	fmt.Println(ShortestPath(grid, 1))
}
