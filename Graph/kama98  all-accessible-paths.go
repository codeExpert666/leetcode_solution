package main

import "fmt"

// 将路径转化为字符串
func pathToStr(path []int) string {
	var res string
	for _, node := range path {
		res += fmt.Sprintf("%d ", node)
	}
	return res[:len(res)-1]
}

// graph的存储形式是邻接表
// 采用dfs获取所有路径，不存储直接输出
// 注意是有向无环图
func getAllPath(graph [][]int, start int, end int) [][]int {
	res := make([][]int, 0)
	path := []int{start}
	// 定义深度优先搜索函数
	var dfs func()
	dfs = func() {
		if path[len(path)-1] == end {
			res = append(res, append([]int{}, path...))
			return
		}
		// 遍历所有未访问过的邻居节点
		for _, node := range graph[path[len(path)-1]-1] {
			path = append(path, node)
			dfs()
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}

func TestDFS() {
	// 处理输入
	var n, m int // 节点数与边数
	fmt.Scan(&n, &m)
	graph := make([][]int, n) // 邻接表
	for i := 0; i < m; i++ {  // 遍历所有边，填充邻接表
		var s, e int // 边的两个顶点
		fmt.Scan(&s, &e)
		graph[s-1] = append(graph[s-1], e)
	}
	// 获取路径
	res := getAllPath(graph, 1, n)
	// 输出
	if len(res) == 0 {
		fmt.Println(-1)
	} else {
		for _, path := range res {
			fmt.Println(pathToStr(path))
		}
	}
}
