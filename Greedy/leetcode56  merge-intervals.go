package Greedy

import "sort"

/**
 * 贪心法
 * 两个维度，先按一个维度左端点进行排序，再逐个遍历，判断与前一区间是否有重叠，然后合并
 */
func merge(intervals [][]int) [][]int {
	// 先对区间按照左端点由小到大进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 遍历所有区间，不断合并
	res := make([][]int, 0, len(intervals))
	res = append(res, intervals[0]) // 初始化
	for i := 1; i < len(intervals); i++ {
		l := len(res)
		if intervals[i][0] <= res[l-1][1] {
			// 当前区间与前一区间有重叠，更新前一区间的右端点
			res[l-1][1] = max(intervals[i][1], res[l-1][1])
		} else {
			// 当前区间与前一区间没有重叠，直接添加当前区间到结果中
			res = append(res, intervals[i])
		}
	}
	return res
}
