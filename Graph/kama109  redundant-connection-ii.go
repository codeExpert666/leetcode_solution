package main

import "fmt"

func isTreeAfterRemoveEdge(edges [][2]int, removeEdge [2]int) bool {
	var s UnionFindSet
	s.Init(len(edges) + 1)
	for _, edge := range edges {
		if edge != removeEdge { // 删除边不加入并查集
			if s.IsConnected(edge[0], edge[1]) { // 出现环
				return false
			}
			s.Join(edge[0], edge[1])
		}
	}
	return true
}

func getRemoveEdge(edges [][2]int) int { // 返回边的下标
	var s UnionFindSet
	s.Init(len(edges) + 1)
	for i, edge := range edges {
		if s.IsConnected(edge[0], edge[1]) {
			return i
		}
		s.Join(edge[0], edge[1])
	}
	return -1
}

// 思路：有向树即只有根节点入度为0，其他节点入度都为1。多一条边后只可能出现以下三种情况之一
// 1、存在入度为2的点，删除其中一条指向该节点的边（删除哪条都行，题目要求靠后输入的边）
// 2、存在入度为2的点，只能固定删除其中一条边（删除另一条导致不是有向树）
// 3、所有点入度均为1，说明环存在，删除引发环的最后一条边
// 对于1和2，尝试删除一条边后若形成有向树，则输出该边，这里删除边后对有向树的判断需要利用并查集（主要判断剩余边是否形成环）
// 对于3，需要检查输入边形成环的那一刻，需要并查集（也是判断环的形成）
// 该方法既需要判断度，又需要并查集，有点复杂
func TestRedundantConnectionii() {
	// var nodeNum int
	// fmt.Scan(&nodeNum)
	// edges := make([][2]int, nodeNum)    // 记录所有输入边
	// inDegrees := make([]int, nodeNum+1) // 记录所有顶点的入度
	// // 统计所有边和入度
	// for i := 0; i < nodeNum; i++ {
	// 	var n1, n2 int
	// 	fmt.Scan(&n1, &n2)
	// 	edges[i] = [2]int{n1, n2}
	// 	inDegrees[n2]++
	// }
	nodeNum := 3
	edges := [][2]int{{1, 2}, {1, 3}, {2, 3}}
	inDegrees := []int{0, 0, 1, 2}
	// 寻找引发入度为2的边，采用倒序遍历，目的是优先判断后输入的边
	var redundantEdges [][2]int // 边要么不存在，要么存在两条
	for i := nodeNum - 1; i >= 0; i-- {
		if inDegrees[edges[i][1]] == 2 {
			redundantEdges = append(redundantEdges, edges[i])
		}
	}
	if len(redundantEdges) > 0 { // 存在引发2入度的边
		// 判断删除两条边中的其中一条能否形成有向树
		if isTreeAfterRemoveEdge(edges, redundantEdges[0]) {
			fmt.Printf("%d %d", redundantEdges[0][0], redundantEdges[0][1])
		} else {
			fmt.Printf("%d %d", redundantEdges[1][0], redundantEdges[1][1])
		}
	} else { // 不存在引发2入度的边，说明形成了环
		// 删除引发环的那条边
		index := getRemoveEdge(edges)
		fmt.Printf("%d %d", edges[index][0], edges[index][1])
	}
}
