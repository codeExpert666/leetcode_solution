package DynamicProgramming

// 递增子序列类型题目
func findNumberOfLIS(nums []int) int {
	var maxLen, count int
	// dp[i]表示以nums[i]结尾的最长递增子序列的长度和个数
	dp := make([][2]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = [2]int{1, 1} // 自身是递增子序列
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				if dp[j][0]+1 == dp[i][0] { // 递增子序列长度相同，增加个数
					dp[i][1] += dp[j][1]
				} else if dp[j][0]+1 > dp[i][0] { // 更长的递增子序列，更新长度，个数重制
					dp[i][0], dp[i][1] = dp[j][0]+1, dp[j][1]
				}
			}
		}
		// 更新递增子序列最大长度和个数
		if dp[i][0] > maxLen { // 更长的递增子序列，更新长度，个数重制
			maxLen, count = dp[i][0], dp[i][1]
		} else if dp[i][0] == maxLen { // 最长递增子序列长度相同，增加个数
			count += dp[i][1]
		}
	}
	return count
}
