package main

import "fmt"

// 图的拓扑排序
func TestTopologicalSorting() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	res := make([]int, 0)              // 节点的拓扑排序结果
	inDegrees := make([]int, nodeNum)  // 记录顶点的入度
	adjTable := make([][]int, nodeNum) // 图的邻接表
	// 处理输入
	for i := 0; i < edgeNum; i++ {
		var n1, n2 int
		fmt.Scan(&n1, &n2)
		inDegrees[n2]++                         // 更新入度
		adjTable[n1] = append(adjTable[n1], n2) // 更新邻接表
	}
	// 拓扑排序初始化（从入度为0的节点开始）
	queue := make([]int, 0)
	for i, v := range inDegrees {
		if v == 0 { // 入度为0，加入队列
			queue = append(queue, i)
		}
	}
	// 拓扑排序过程：1、找到入度为0节点 2、删除入度为0节点 1与2不断循环
	for len(queue) > 0 { // queue中始终存储入度为0的节点
		node := queue[0]
		queue = queue[1:]
		res = append(res, node) // 添加到排序结果中
		// 删除node，主要工作是减少其指向顶点的入度
		for _, n := range adjTable[node] {
			inDegrees[n]--
			if inDegrees[n] == 0 { // 指向顶点入度变为0，加入队列
				queue = append(queue, n)
			}
		}
	}
	if len(res) == nodeNum { // 图中不存在环
		for i, n := range res {
			if i != len(res)-1 {
				fmt.Printf("%d ", n)
			} else {
				fmt.Print(n)
			}
		}
	} else { // 图中存在环
		fmt.Print(-1)
	}
}
