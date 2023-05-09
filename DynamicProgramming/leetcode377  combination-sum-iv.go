package DynamicProgramming

/**
 * 完全背包应用（dp）
 * 纯粹套模板，注意两个点：1、完全背包 2、排列而不是组合
 */
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}
