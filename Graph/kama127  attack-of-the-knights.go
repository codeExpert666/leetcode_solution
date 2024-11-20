package main

import (
	"container/heap"
	"fmt"
)

// Knight 棋子所处位置
type Knight struct {
	x, y    int // 位置
	f, g, h int // 启发式函数h = f + g，f是起点到当前节点到距离，g是当前节点与终点的距离，距离计算均采用欧式距离
}

// PQueue 类型为Knight的优先队列
type PQueue []Knight

func (pq *PQueue) Len() int {
	return len(*pq)
}

func (pq *PQueue) Less(i, j int) bool {
	return (*pq)[i].h < (*pq)[j].h
}

func (pq *PQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Knight))
}

func (pq *PQueue) Pop() interface{} {
	l := len(*pq)
	top := (*pq)[l-1]
	*pq = (*pq)[:l-1]
	return top
}

// AStar A*算法，本题的启发式函数是距离，启发式分数越小，优先级越高
func AStar(startx, starty, endx, endy int) int {
	var grid [1001][1001]int                                                                       // 定义地图
	neighbors := [8][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}} // 棋子移动规则
	pq := make(PQueue, 0)                                                                          // 启发式优先队列
	// 初始化队列
	heap.Init(&pq)
	disToEnd := (endx-startx)*(endx-startx) + (endy-starty)*(endy-starty) // 当前节点与终点的欧式距离（省略开方，便于计算且不影响结果）
	heap.Push(&pq, Knight{startx, starty, 0, disToEnd, disToEnd})
	grid[startx][starty]++ // 防止起点为0与未访问节点也为0混淆，赋值为1
	for pq.Len() > 0 {
		// 取出启发式函数值最大的节点
		cur := heap.Pop(&pq).(Knight)
		if cur.x == endx && cur.y == endy { // 到达终点
			return grid[endx][endy] - 1 // 返回步长（减1是因为起点设置为了1）
		}
		for i := 0; i < 8; i++ { // 遍历下一步所有可能位置
			nextx, nexty := cur.x+neighbors[i][0], cur.y+neighbors[i][1]
			if nextx < 1 || nextx > 1000 || nexty < 1 || nexty > 1000 { // 下一位置超出地图边界
				continue
			}
			if grid[nextx][nexty] == 0 { // 下一位置未访问
				disToEnd = (nextx-endx)*(nextx-endx) + (nexty-endy)*(nexty-endy)
				heap.Push(&pq, Knight{ // 位置加入队列
					x: nextx,
					y: nexty,
					f: cur.f + 5,
					g: disToEnd,
					h: cur.f + 5 + disToEnd,
				})
				grid[nextx][nexty] = grid[cur.x][cur.y] + 1 // 步长增加
			}
		}
	}
	return -1 // 终点无法到达，在本题中不会出现
}

func main() {
	var queryNum int
	_, _ = fmt.Scan(&queryNum)
	answer := make([]int, queryNum)
	for i := 0; i < queryNum; i++ {
		var startx, starty, endx, endy int
		_, _ = fmt.Scan(&startx, &starty, &endx, &endy)
		answer[i] = AStar(startx, starty, endx, endy)
	}
	// 输出结果
	for _, v := range answer {
		fmt.Println(v)
	}
}
