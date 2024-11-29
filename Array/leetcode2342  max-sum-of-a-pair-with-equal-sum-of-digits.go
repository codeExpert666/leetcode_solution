package Array

import (
	"slices"
)

func maximumSum(nums []int) int {
	// 数组降序排列，保证较大值先访问
	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	// 记录所有数字的数位和
	bitSum := [82][]int{}
	for _, v := range nums {
		s := 0 // 数位和
		for tmp := v; tmp != 0; tmp /= 10 {
			s += tmp % 10
		}
		bitSum[s] = append(bitSum[s], v)
	}
	// 对于每个可能的数位和，取出最大的两个数做和
	res := -1
	for i := 1; i <= 81; i++ {
		if len(bitSum[i]) <= 1 { // 该数位和对应的数字数量不足 2 个
			continue
		}
		if s := bitSum[i][0] + bitSum[i][1]; s > res { // 由于已排序，前两个元素最大
			res = s
		}
	}
	return res
}

// 优化思路：该问题的一个子问题是：给定数组，任意取出两个元素，获取最大和。
// 对于这个子问题，可采用一次遍历法。考虑到最大和只可能是一个元素和其前方（数组头部方向）最大元素的和
// 故可以顺序遍历，记录当前最大值，每遍历到一个新元素，判定当前元素与当前最大值的和是否大于最大和，同时更新当前最大值
func maximumSum1(nums []int) int {
	res := -1
	// 记录所有数字的数位和
	bitSum := [82]int{}
	for _, v := range nums {
		// 获取数位和
		s := 0
		for tmp := v; tmp != 0; tmp /= 10 {
			s += tmp % 10
		}
		// 更新最大和
		if bitSum[s] != 0 && v+bitSum[s] > res {
			res = v + bitSum[s]
		}
		// 更新数位和当前最大值
		if v > bitSum[s] {
			bitSum[s] = v
		}
	}
	return res
}
