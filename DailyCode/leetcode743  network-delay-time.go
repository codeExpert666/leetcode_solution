package DailyCode

import (
	"container/heap"
	"math"
)

type partEdge struct { // 邻接表和优先队列的元素类型
	toNode int
	weight int
}

// 最好再写一个结构体，专门用于描述优先队列的元素，这样代码可读性强，不易混淆
type pQueue []partEdge

func (pq *pQueue) Len() int {
	return len(*pq)
}

func (pq *pQueue) Less(i, j int) bool {
	return (*pq)[i].weight < (*pq)[j].weight
}

func (pq *pQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *pQueue) Push(x any) {
	*pq = append(*pq, x.(partEdge))
}

func (pq *pQueue) Pop() any {
	top := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return top
}

func networkDelayTime(times [][]int, n int, k int) int {
	// 构建图结构（邻接表）
	adjTable := make([][]partEdge, n+1)
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		adjTable[u] = append(adjTable[u], partEdge{v, w})
	}
	// 队列优化版Dijkstra算法
	minTime := make([]int, n+1) // 源点到其他节点到最短时延
	for i := 1; i <= n; i++ {   // 初始化
		if i != k {
			minTime[i] = math.MaxInt
		}
	}
	determined := make([]bool, n+1) // 节点的最短时延是否确定
	pq := &pQueue{partEdge{k, 0}}
	heap.Init(pq) // 小根堆，每次获取最短时延的节点，源点入队
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(partEdge) // 最短时延节点出队
		if determined[cur.toNode] {    // 优先队列中节点可能重复加入，同一节点靠后弹出的无效
			continue
		}
		determined[cur.toNode] = true               // 该节点的最短时延已确定
		for _, next := range adjTable[cur.toNode] { // 访问邻居节点
			if !determined[next.toNode] && cur.weight+next.weight < minTime[next.toNode] {
				// 邻居节点未确定最短时延，且发现更短时延
				minTime[next.toNode] = cur.weight + next.weight            // 更新最短时延
				heap.Push(pq, partEdge{next.toNode, minTime[next.toNode]}) // 加入队列
			}
		}
	}
	// 所有节点收到信号的最短时延由最慢收到信号的节点确定
	var res int
	for _, t := range minTime {
		if t > res {
			res = t
		}
	}
	if res == math.MaxInt { // 存在节点无法收到信号
		res = -1
	}
	return res
}
