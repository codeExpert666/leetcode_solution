package Greedy

import "sort"

/**
 * 贪心法
 * 本题涉及两个维度，同样，还是先考虑一个维度，再在第一个维度的基础上考虑第二维度
 * 先只考虑气球左端点，对其进行排序，再考虑右端点的交叠情况，交叠情况的出现可以减少箭的数量
 */
func findMinArrowShots(points [][]int) int {
	// 按照左端点排序
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			// 左端点相同的，右端点越大越靠前，因为一支箭射这两个气球，主要取决于范围小的气球
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	// 遍历数组
	var res int
	for i := 0; i < len(points)-1; i++ {
		if points[i][1] < points[i+1][0] {
			// 前后两个气球没有交叠，只能单独为当前气球发射一支箭
			res++
		} else if points[i][1] < points[i+1][1] {
			// 前后气球有部分交叠，当且仅当箭射在交集中，可减少箭的数量
			// 但现在还不能确定增加一支箭，因为该交集也许与后面气球范围有交集
			// 直到前面所有气球的交集与下一气球没有交集时，才能确定发射一支箭
			// 贪心就贪在这里，尽可能减少箭的数量
			points[i+1][1] = points[i][1]
		}
		// 当前后气球完全重叠时，无需处理，因为根据前面的排序，两气球的交集就是下一个气球
	}
	// 最后一个气球或最后一串气球还未处理，仍需要一支箭
	res++
	return res
}
