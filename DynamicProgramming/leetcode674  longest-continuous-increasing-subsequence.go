package DynamicProgramming

/**
 * 常规方法，一次遍历，比较所有可能的递增连续子序列长度
 */
func findLengthOfLCIS(nums []int) int {
	var res = 1
	var count = 1 // 记录当前连续递增子序列长度
	for i := 1; i < len(nums); i++ {
		// 更新递增子序列长度
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}
		// 更新结果
		if count > res {
			res = count
		}
	}
	return res
}

/**
 * 动态规划
 * dp[i]代表以nums[i]结尾的最长递增连续子序列的长度
 */
func findLengthOfLCIS1(nums []int) int {
	var res = 1
	// dp初始化
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
