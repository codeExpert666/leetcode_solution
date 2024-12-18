package Array

import (
	"math"
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

// 排序+二分查找，每一次查询都是独立的，没有对一系列查询进行优化
func ClosestRoom(rooms [][]int, queries [][]int) []int {
	// 首先对房间按照面积进行升序排列
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][1] < rooms[j][1]
	})
	// 处理所有查询
	res := make([]int, len(queries))
	for i, query := range queries {
		// 二分查找满足面积要求的第一个房间
		l, r := 0, len(rooms)-1
		for l <= r {
			mid := l + (r-l)>>1
			if rooms[mid][1] < query[1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		// 此时 l 指向第一个满足面积条件的房间
		// 便利 l 之后的房间，找到编号最近的房间
		if l == len(rooms) { // 没有房间满足面积要求
			res[i] = -1
			continue
		}
		minDist, curNo := math.MaxInt, 0
		for j := l; j < len(rooms); j++ {
			if d := abs(rooms[j][0] - query[0]); d < minDist {
				minDist, curNo = d, rooms[j][0]
			} else if d == minDist && rooms[j][0] < curNo { // 距离相同选编号小的房间
				curNo = rooms[j][0]
			}
		}
		res[i] = curNo
	}
	return res
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// 离线算法+二分查找+有序集合
// 离线算法的含义是，事先已经知道了所有的查询，对一系列查询进行优化
// 有序集合利用了 github 的一个 go 语言数据结构的开源模块。很好的模块，要学会使用
func ClosestRoom1(rooms [][]int, queries [][]int) []int {
	q := len(queries)
	// 按照面积降序对房间进行排列
	sort.Slice(rooms, func(i, j int) bool { return rooms[i][1] > rooms[j][1] })
	// 按照面积降序对所有查询进行排列（离线优化的主要步骤）
	// 由于查询的顺序不能修改，这里选择对下标进行排序
	idx := make([]int, q) // 构建下标数组，然后进行排序
	for i := 0; i < q; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { return queries[idx[i]][1] > queries[idx[j]][1] })
	// 处理所有查询
	res := make([]int, q)
	for i := 0; i < q; i++ { // 查询结果初始化
		res[i] = -1
	}
	roomIds := redblacktree.NewWithIntComparator() // 构建有序集合，存储满足查询面积要求的所有房间编号
	j := 0                                         // 标记下一次查询要处理的起始房间
	for _, index := range idx {
		preferred, minSize := queries[index][0], queries[index][1]
		// 将待处理房间加入有序集合
		for j < len(rooms) && rooms[j][1] >= minSize {
			roomIds.Put(rooms[j][0], struct{}{})
			j++
		}
		// 从有序集合中找出与查询相近的房间编号
		minDist := math.MaxInt
		// 左侧房间编号（小于）
		if node, ok := roomIds.Floor(preferred); ok {
			res[index] = node.Key.(int)
			minDist = preferred - res[index]
		}
		// 右侧房间编号（大于）
		if node, ok := roomIds.Ceiling(preferred); ok && node.Key.(int)-preferred < minDist {
			res[index] = node.Key.(int)
		}
	}
	return res
}
