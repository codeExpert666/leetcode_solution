package main

import "fmt"

// TestRestrictedSingleSourceShortestPath 单源有限最短路径
// bellman_ford可以通过在每轮松弛时对minDist数组进行备份实现（主要是为了应对负权回路的情况）
// SPFA（队列优化版bellman）可以通过类似于层序遍历分层输出节点的方式实现（保证每次往处理向前走一步的所有节点，无需备份minDist数组）
func TestRestrictedSingleSourceShortestPath() {
	var nodeNum, edgeNum int
	_, _ = fmt.Scan(&nodeNum, &edgeNum)
	// 构建邻接表
	adjTable := make([][]PartEdge, nodeNum+1)
	for i := 0; i < edgeNum; i++ {
		var n1, n2, w int
		_, _ = fmt.Scan(&n1, &n2, &w)
		adjTable[n1] = append(adjTable[n1], PartEdge{n2, w})
	}
	// 初始化
	var src, dst, maxLen int
	_, _ = fmt.Scan(&src, &dst, &maxLen)
	minDist := make([]int, nodeNum+1)
	for i := 1; i < len(minDist); i++ {
		if i != src {
			minDist[i] = MaxInt
		}
	}
	queue := []int{src}
	isInQueue := make([]bool, nodeNum+1)
	// 算法过程：每一轮向前走一步，最多走maxLen步
	for count, l := 0, len(queue); count < maxLen && l > 0; count, l = count+1, len(queue) {
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			isInQueue[cur] = false
			for _, pe := range adjTable[cur] {
				if minDist[cur]+pe.Weight < minDist[pe.ToNode] {
					minDist[pe.ToNode] = minDist[cur] + pe.Weight
					if !isInQueue[pe.ToNode] { // 注意这个条件放在哪，不管在不在队列里，发现更小距离都要更新，所以放在上一层if里是不对的
						queue = append(queue, pe.ToNode)
						isInQueue[pe.ToNode] = true
					}
				}
			}
		}
	}
	// 输出结果
	if minDist[dst] == MaxInt {
		fmt.Print("unreachable")
	} else {
		fmt.Print(minDist[dst])
	}
}
