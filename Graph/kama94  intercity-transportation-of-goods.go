package main

import "fmt"

// Bellman_ford算法，能够处理带负权制的单源最短路径
func TestBellmanFord() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	// 存储图中所有边
	edgeList := make([][3]int, edgeNum)
	for i := 0; i < edgeNum; i++ {
		edgeList[i] = [3]int{}
		fmt.Scan(&edgeList[i][0], &edgeList[i][1], &edgeList[i][2])
	}
	// Bellman_ford算法初始化
	minDist := make([]int, nodeNum+1)
	for i := 2; i < len(minDist); i++ {
		minDist[i] = MaxInt
	}
	// Bellman_ford算法具体过程
	for i := 1; i < nodeNum; i++ {
		// 对所有边松弛一次，相当于计算 起点到达 与起点一条边相连的节点 的最短距离。
		// 起点与终点间最多相距n-1条边，故需松弛n-1次
		updated := false
		for _, edge := range edgeList { // 对所有边松弛一次
			n1, n2, w := edge[0], edge[1], edge[2]
			if minDist[n1] != MaxInt && minDist[n1]+w < minDist[n2] { // 松弛操作
				minDist[n2] = minDist[n1] + w
				updated = true
			}
		}
		if !updated {
			// 松弛n-1次对于任何图是通用的，一定能够保证所有长度（边的个数）的路径都能被考虑到
			// 然而，如若某图中从源点出发的最长路径中边的个数小于n-1，则不用松弛n-1次
			// 只要发生minDist不再更新，即可停止
			break
		}
	}
	// 输出结果
	if minDist[nodeNum] == MaxInt { // 终点无法到达
		fmt.Print("unconnected")
	} else {
		fmt.Print(minDist[nodeNum])
	}
}
