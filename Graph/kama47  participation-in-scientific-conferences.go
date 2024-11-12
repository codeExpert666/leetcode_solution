package main

import (
	"fmt"
)

const MaxInt int = 1<<31 - 1

type PartEdge struct { // 邻接表中的元素类型
	ToNode int
	Weight int
}

// 单源最短路径Dijkstra算法
func main() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	// 构建图（邻接表）
	adjTable := make([][]PartEdge, nodeNum+1)
	for i := 0; i < edgeNum; i++ {
		var n1, n2, weight int
		fmt.Scan(&n1, &n2, &weight)
		adjTable[n1] = append(adjTable[n1], PartEdge{n2, weight})
	}
	// 初始化Dijkstra
	minDist := make([]int, nodeNum+1)  // 所有节点到源点的最短路径
	visited := make([]bool, nodeNum+1) // 节点是否已访问
	for i := 0; i < len(minDist); i++ {
		if i == 1 {
			minDist[i] = 0
		} else {
			minDist[i] = MaxInt
		}
	}
	// Dijkstra运行过程
	for i := 0; i < nodeNum; i++ { // 获取每个节点的最短路径
		// 第一步：从minDist中找到距离源点最近的未访问节点
		minValue, cur := MaxInt, -1
		for j := 1; j < len(minDist); j++ {
			if !visited[j] && minDist[j] < minValue {
				minValue = minDist[j]
				cur = j
			}
		}
		if cur == -1 { // 未访问节点均无法到达
			break
		}
		// 第二步：将当前节点标记为访问
		visited[cur] = true
		// 第三步：更新minDist数组（与cur相连且未被访问节点）
		for _, pe := range adjTable[cur] {
			if !visited[pe.ToNode] && minDist[cur]+pe.Weight < minDist[pe.ToNode] {
				minDist[pe.ToNode] = minDist[cur] + pe.Weight
			}
		}
	}
	// 输出
	if minDist[nodeNum] == MaxInt { // 起点与终点没有路径
		fmt.Println(-1)
	} else { // 起点与终点存在路径
		fmt.Println(minDist[nodeNum])
	}
}
