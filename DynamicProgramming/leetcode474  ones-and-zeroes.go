package DynamicProgramming

/**
 * 01背包问题应用（dp）
 */
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		// 先获取当前字符串中0与1的个数
		var c0, c1 int
		s := strs[i]
		for j := 0; j < len(s); j++ {
			if s[j] == '0' {
				c0++
			} else {
				c1++
			}
		}
		// dp过程
		for k := m; k >= c0; k-- {
			for p := n; p >= c1; p-- {
				dp[k][p] = max(dp[k][p], dp[k-c0][p-c1]+1)
			}
		}
	}
	return dp[m][n]
}
