package Simulation

import "slices"

// 本题的核心是求出 meetings 中所有区间的并集
func CountDays(days int, meetings [][]int) int {
	// 先将所有区间按照左端点进行升序排序
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})
	// 遍历所有区间，逐个求取并集
	unionSet := [][]int{meetings[0]}                // 并集
	unionLen := meetings[0][1] - meetings[0][0] + 1 // 并集的长度
	for i := 1; i < len(meetings); i++ {
		last := unionSet[len(unionSet)-1] // 并集可能不连续，获取最后一个区间
		if meetings[i][0] <= last[1] {    // 当前区间与最后一个区间相交
			if meetings[i][1] > last[1] { // 当前区间并不完全被最后一个区间包含
				unionLen += meetings[i][1] - last[1] // 先算长度，防止 last[1] 被更新
				last[1] = meetings[i][1]
			}
		} else { // 当前区间与最后一个区间不相交
			unionSet = append(unionSet, meetings[i])
			unionLen += meetings[i][1] - meetings[i][0] + 1
		}
	}
	return days - unionLen
}
