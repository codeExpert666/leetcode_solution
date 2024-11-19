package main

import "fmt"

const MaxInt int = 1<<31 - 1

// TestNegativeWeightingCircuit 向bellman_ford算法中添加负权回路的判断
// 正常的bellman_ford算法在经过nodeNum-1次循环后距离数组一定不再更新
// 若存在负权回路，距离可以无限小，故会一直更新
func TestNegativeWeightingCircuit() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	// 构建图结构（存储所有边）
	edges := make([][3]int, edgeNum)
	for i := 0; i < edgeNum; i++ {
		fmt.Scan(&edges[i][0], &edges[i][1], &edges[i][2])
	}
	// 算法初始化
	minDist := make([]int, nodeNum+1)
	for i := 2; i < len(minDist); i++ {
		minDist[i] = MaxInt
	}
	// 算法过程
	for loop := 1; loop <= nodeNum; loop++ { // 比普通bellman_ford算法多循环一次，观察minDist是否更新
		updated := false
		for i := 0; i < edgeNum; i++ {
			n1, n2, w := edges[i][0], edges[i][1], edges[i][2]
			if loop < nodeNum && minDist[n1] != MaxInt && minDist[n1]+w < minDist[n2] {
				minDist[n2] = minDist[n1] + w
				updated = true
			} else if loop == nodeNum && minDist[n1] != MaxInt && minDist[n1]+w < minDist[n2] {
				fmt.Print("circle")
				return
			}
		}
		if !updated {
			break
		}
	}
	// 输出结果
	if minDist[nodeNum] == MaxInt {
		fmt.Print("unconnected")
	} else {
		fmt.Print(minDist[nodeNum])
	}
}
