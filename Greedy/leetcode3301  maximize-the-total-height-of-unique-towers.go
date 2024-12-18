package Greedy

import (
	"math"
	"slices"
)

// 贪心：降序排列后尽可能使得每栋楼的高度最大
func MaximumTotalSum(maximumHeight []int) int64 {
	// 降序排列最大高度数组
	slices.SortFunc(maximumHeight, func(a, b int) int { return b - a })
	// 初步判断
	if maximumHeight[0] < len(maximumHeight) {
		return -1
	}
	// 遍历数组，让每栋楼的高度尽可能大
	res, maxHeight := 0, math.MaxInt // maxHeight 表示在前面楼的影响下，当前楼的最大高度
	for _, h := range maximumHeight {
		if maxHeight == 0 { // 高度必须为正整数
			return -1
		}
		cur := min(maxHeight, h) // 每栋楼的高度收到自身和前面楼的共同限制
		res += cur
		maxHeight = cur - 1
	}
	return int64(res)
}
