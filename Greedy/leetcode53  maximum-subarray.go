package Greedy

import "math"

/**
 * 方法一：暴力解法，找到所有可能的子序列，记录比较它们的和，从而获得最大子序列
 * 这种方法时间复杂度为O(n^2)，空间复杂度为O(1)
 */
func maxSubArray(nums []int) int {
	res := math.MinInt
	// 两层遍历，寻找所有连续子序列
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			// 遇到更大的和就更新
			if sum > res {
				res = sum
			}
		}
	}
	return res
}

/**
 * 方法二：贪心法，这种方法基于上面的暴力法，暴力法找到所有连续子序列，而贪心法会进行一定的“剪枝”
 * 贪心法通过调整连续子序列的区间端点从而提前规避了很多不必要的搜索
 */
func maxSubArray1(nums []int) int {
	res := math.MinInt
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 遇到更大的和就更新
		if sum > res {
			res = sum
		}
		// !!! 剪枝体现在这里
		// 一旦和变为负数，立即将区间端点调整至i的下一个位置，重新开始计算子序列和，这便是贪心的体现
		// 可以证明，当和变为负数时，区间端点不可能是i前的任意一个位置（包括i）
		if sum < 0 { // 这里条件写成sum <= 0也可以
			sum = 0
		}
	}
	return res
}
