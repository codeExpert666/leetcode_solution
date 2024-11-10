package main

import "fmt"

// Prim算法获取最小生成树
func getMinDis(adjMatrix [][]int) int {
	minDis := make([]int, len(adjMatrix)) // 记录图中顶点与生成树的最小距离
	for i := 1; i < len(minDis); i++ {
		minDis[i] = 10001 // 初始化最大距离，以保证后续更小距离的更新
	}
	isInTree := make([]bool, len(adjMatrix)) // 标记图中每个点是否在最小生成树中
	for i := 2; i < len(adjMatrix); i++ {    // i初始值不重要，该循环目的是运行nodeNum-1次，以获取生成树的所有边
		// 第一步：从minDis中找到不在生成树中且距离生成树最近的顶点
		cur, minVal := -1, 10002
		for j := 1; j < len(minDis); j++ {
			if !isInTree[j] && minDis[j] < minVal {
				cur, minVal = j, minDis[j]
			}
		}
		// 第二步：将该顶点加入生成树中
		isInTree[cur] = true
		// 第三步：考虑到树的更新，minDis也需要更新(更新不在树中的节点，更具体说是更新cur的邻接顶点)
		for j := 1; j < len(minDis); j++ {
			if !isInTree[j] && adjMatrix[cur][j] < minDis[j] {
				minDis[j] = adjMatrix[cur][j]
			}
		}
	}
	// 最小生成树获取完毕后，minDis中保留了该树所有边的权值
	// 节点1的minDis除外，因为该点是起始点，且没有入边，minDis未更新
	var res int
	for i := 2; i < len(minDis); i++ {
		res += minDis[i]
	}
	return res
}

func TestPrim() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	// 构建邻接矩阵（无向图）
	adjMatrix := make([][]int, nodeNum+1)
	for i := 1; i <= nodeNum; i++ {
		adjMatrix[i] = make([]int, nodeNum+1)
		for j := 1; j <= nodeNum; j++ {
			adjMatrix[i][j] = 10001 // 超出权值范围表示无连接
		}
	}
	for i := 0; i < edgeNum; i++ {
		var n1, n2, w int
		fmt.Scan(&n1, &n2, &w)
		adjMatrix[n1][n2] = w
		adjMatrix[n2][n1] = w
	}
	// 调用Prim算法获取最小生成树的权值和
	fmt.Println(getMinDis(adjMatrix))
}
