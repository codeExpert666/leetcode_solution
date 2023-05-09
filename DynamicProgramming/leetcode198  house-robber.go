package DynamicProgramming

/**
 * 传统动态规划（dp）
 */
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 2; i < len(dp); i++ {
		// 递推公式
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}
