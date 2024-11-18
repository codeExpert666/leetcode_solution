package main

import (
	"fmt"
)

type PartEdge struct { // 邻接表中的元素类型
	ToNode int
	Weight int
}

// 队列优化版Bellman_ford算法，能够处理带负权制的单源最短路径
// 每次遍历无需再松弛所有边，只需要对 上一次松弛的时候更新过的节点作为出发节点所连接的边 进行松弛就够了。
func TestQueueBellmanFord() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	// 构建图结构（邻接表）
	adjTable := make([][]PartEdge, nodeNum+1)
	for i := 0; i < edgeNum; i++ {
		var n1, n2, weight int
		fmt.Scan(&n1, &n2, &weight)
		adjTable[n1] = append(adjTable[n1], PartEdge{n2, weight})
	}
	// 队列优化版Bellman_ford算法初始化
	minDist := make([]int, nodeNum+1) // 节点到源点的最近距离
	for i := 2; i < len(minDist); i++ {
		minDist[i] = MaxInt
	}
	queue := []int{1}                    // 存储后续需要访问的节点，精简了边松弛的数量
	isInQueue := make([]bool, nodeNum+1) // 节点是否已在队列中
	for len(queue) > 0 {
		// 弹出节点
		node := queue[0]
		queue = queue[1:]
		isInQueue[node] = false
		// 仅访问邻居（当前节点的最短距离发生改变，可能会引发邻居的最短距离改变。反之，对于那些没有改变的节点，则不用考虑）
		// 精简就体现在这里
		for _, pe := range adjTable[node] {
			if minDist[node]+pe.Weight < minDist[pe.ToNode] { // 发现到达pe.ToNode的更近路径
				minDist[pe.ToNode] = minDist[node] + pe.Weight // 更新路径长度
				// 加入队列（已经在队列中的节点不再重复加入）
				if !isInQueue[pe.ToNode] {
					queue = append(queue, pe.ToNode)
					isInQueue[pe.ToNode] = true
				}
			}
		}
	}
	// 输出距离
	if minDist[nodeNum] == MaxInt { // 终点无法到达
		fmt.Print("unconnected")
	} else {
		fmt.Print(minDist[nodeNum])
	}
}
