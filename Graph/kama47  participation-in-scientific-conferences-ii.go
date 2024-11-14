package main

import (
	"container/heap"
	"fmt"
)

// PriorityQueue 实现了 heap.Interface 并保存 PartEdge 元素
type PriorityQueue []PartEdge

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 根据 Weight 小的优先级更高
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// 注意：这里需要使用指针接收，因为需要对pq本身做出改变
// 如果只是对其内部元素做出改变，值接收即可
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(PartEdge))
}

// 堆顶元素位于切片索引0的位置。然而，在堆的实现中，推出堆顶元素会经历以下步骤：
// 1、heap.Pop 函数首先将堆顶元素与切片的最后一个元素交换位置。
// 2、然后调用 Pop() 方法移除并返回切片的最后一个元素，即原本的堆顶元素。
// 通过这种方式，堆的性质得以保持，剩余的元素重新调整为有效的堆结构。
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	topHeap := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return topHeap
}

// 堆优化版本Dijkstra算法
func TestDijkstraHeap() {
	var nodeNum, edgeNum int
	fmt.Scan(&nodeNum, &edgeNum)
	// 构建图结构（邻接表）
	adjTable := make([][]PartEdge, nodeNum+1)
	for i := 0; i < edgeNum; i++ {
		var n1, n2, weight int
		fmt.Scan(&n1, &n2, &weight)
		adjTable[n1] = append(adjTable[n1], PartEdge{n2, weight})
	}
	// Dijkstra算法初始化
	minDis := make([]int, nodeNum+1) // 所有节点到源点的最短距离
	for i := 2; i < len(minDis); i++ {
		minDis[i] = MaxInt
	}
	visited := make([]bool, nodeNum+1) // 节点是否访问
	pq := &PriorityQueue{}             // 优先队列，保证距离源点最近的节点在堆顶 为何指针?
	heap.Init(pq)
	heap.Push(pq, PartEdge{1, 0})
	// Dijkstra算法过程
	for pq.Len() > 0 {
		// 第一步：寻找距离源点最近节点
		cur := heap.Pop(pq).(PartEdge)
		// 第二步：标记当前节点已访问
		visited[cur.ToNode] = true
		// 第三步：更新minDis与pq
		for _, nextNode := range adjTable[cur.ToNode] {
			if !visited[nextNode.ToNode] {
				newDis := cur.Weight + nextNode.Weight
				if newDis < minDis[nextNode.ToNode] {
					minDis[nextNode.ToNode] = newDis
					heap.Push(pq, PartEdge{nextNode.ToNode, newDis})
				}
			}
		}
	}
	// 输出最短距离
	if minDis[nodeNum] == MaxInt { // 孤立节点
		fmt.Println(-1)
	} else { // 起点能够到达终点
		fmt.Println(minDis[nodeNum])
	}
}
