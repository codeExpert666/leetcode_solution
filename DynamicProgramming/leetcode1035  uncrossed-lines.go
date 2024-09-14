package DynamicProgramming

/**
 * 问题可以转化为两个序列找最长子序列
 * dp[i][j]代表子序列nums[i-1]和子序列nums[j-1]的最长相同子串长度
 */
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([]int, len(nums2)+1)

	for i := 0; i < len(nums1); i++ {
		var tmp int
		for j := 1; j <= len(nums2); j++ {
			dpCopy := dp[j]
			if nums1[i] == nums2[j-1] {
				dp[j] = tmp + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			tmp = dpCopy
		}
	}
	return dp[len(nums2)]
}
