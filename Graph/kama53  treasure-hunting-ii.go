package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	N1, N2 int // 边两端顶点
	Weight int // 边权重
}

// Kruskal最小生成树：核心是边，向树中逐个添加边
func TestKruskal() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	edges := make([]Edge, edgeNum) // 存储所有边
	var s UnionFindSet             // 并查集判断添加的边是否会在树中形成环
	s.Init(nodeNum + 1)
	// 保存所有边
	for i := 0; i < edgeNum; i++ {
		var n1, n2, w int
		fmt.Scan(&n1, &n2, &w)
		edges[i] = Edge{n1, n2, w}
	}
	// 对所有边按照权值由小到大排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight <= edges[j].Weight
	})
	// 遍历所有边，按照贪心策略逐步将边加入树中，但不能形成环
	var res int
	for _, e := range edges {
		if !s.IsConnected(e.N1, e.N2) {
			res += e.Weight
			s.Join(e.N1, e.N2)
		}
	}
	fmt.Println(res)
}
