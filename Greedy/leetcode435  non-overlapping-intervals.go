package Greedy

import (
	"sort"
)

/**
 * 贪心法
 * 依旧是左右端点两个维度，先考虑一个维度，再在基础上考虑另一个维度
 * 先考虑左端点，按其对所有区间排序，然后再顺序遍历，处理交叠情况：当某一个区间与前一区间交叠时
 * 保留靠前的那个区间，这样可以尽可能的减少移除的区间数量，贪心便体现在这。
 */
func eraseOverlapIntervals(intervals [][]int) int {
	// 先对区间按照左端点由小到大进行排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	// 遍历数组
	var res int
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] { // 当前区间与前一区间出现交叠
			res++ // 需要移除一个区间
			if intervals[i-1][1] < intervals[i][1] {
				// 保留靠前的区间
				intervals[i][0] = intervals[i-1][0] // 这句可以不要，因为左端点顺序排列，左端点不会影响结果
				intervals[i][1] = intervals[i-1][1]
			}
		}
		// 其他情况无需特殊处理：不重叠时不用处理，完全重叠时排序保证了当前区间一定是靠前区间
	}
	return res
}
