package main

import (
	"fmt"
)

// Floyd多源最短路径，动态规划思想
func main() {
	var nodeNum, edgeNum int
	_, _ = fmt.Scan(&nodeNum, &edgeNum)
	// 构建图结构（邻接矩阵充当dp数组）
	adjMatrix := make([][]int, nodeNum+1)
	for i := 1; i <= nodeNum; i++ {
		adjMatrix[i] = make([]int, nodeNum+1)
		for j := 1; j <= nodeNum; j++ {
			adjMatrix[i][j] = MaxInt // 后续不断记录最小距离，故初始化为最大值
		}
	}
	for i := 0; i < edgeNum; i++ {
		var n1, n2, w int
		_, _ = fmt.Scan(&n1, &n2, &w)
		adjMatrix[n1][n2] = w
		adjMatrix[n2][n1] = w
	}
	// Floyd算法过程，dp本该是三维矩阵，空间优化后可用二维矩阵替代
	for k := 1; k <= nodeNum; k++ {
		for i := 1; i <= nodeNum; i++ {
			for j := 1; j <= nodeNum; j++ {
				if adjMatrix[i][k]+adjMatrix[k][j] < adjMatrix[i][j] {
					adjMatrix[i][j] = adjMatrix[i][k] + adjMatrix[k][j]
				}
			}
		}
	}
	// 此时邻接矩阵中存储了任意两点间的最短路径，输出所需结果
	var res []int
	var queryNum int
	_, _ = fmt.Scan(&queryNum)
	for i := 0; i < queryNum; i++ {
		var src, dst int
		_, _ = fmt.Scan(&src, &dst)
		res = append(res, adjMatrix[src][dst])
	}
	for _, l := range res {
		if l == MaxInt {
			fmt.Println(-1)
		} else {
			fmt.Println(l)
		}
	}
}
