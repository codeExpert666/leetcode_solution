package DynamicProgramming

/**
 * 动态规划 滚动数组节省空间
 * 这里dp[i][j]代表长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列长度
 */
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)

	for i := 0; i < len(text1); i++ {
		var tmp int // tmp与dpCopy都是为了记录上一层dp[j-1]的值，因为在遍历的过程中会被覆盖
		for j := 1; j <= len(text2); j++ {
			// 递推公式
			dpCopy := dp[j]
			if text2[j-1] == text1[i] {
				dp[j] = tmp + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			tmp = dpCopy
		}
	}
	return dp[len(text2)]
}
