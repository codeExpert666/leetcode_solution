package Greedy

import (
	"slices"
)

// 尽可能移除数量多的元素，这样可以使移除的元素种类最少
func MinSetSize(arr []int) int {
	n := len(arr)
	// 统计元素出现频数
	elemToFrequency := make(map[int]int)
	for _, v := range arr {
		elemToFrequency[v]++
	}
	// 对哈希表的频数进行降序排列
	frequency := make([]int, 0, len(elemToFrequency))
	for _, v := range elemToFrequency {
		frequency = append(frequency, v)
	}
	slices.SortFunc(frequency, func(a, b int) int {
		return b - a
	})
	// 统计移除的最少元素种类
	var removeCount int
	for i, v := range frequency {
		removeCount += v
		if removeCount >= n>>1 {
			return i + 1
		}
	}
	return -1 // 不会运行到这
}
