package DailyCode

func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// 创建一个 (m+1) x (n+1) 的二维布尔数组
	// dp[i][j] 表示 s 的前 i 个字符和 p 的前 j 个字符是否匹配
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 空字符串匹配空模式
	dp[0][0] = true

	// 处理模式 p 以星号开始的情况
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			// 星号可以匹配零个字符，所以如果前两个字符匹配，这里也匹配
			dp[0][j] = dp[0][j-2]
		}
	}

	// 填充 dp 数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// 如果当前字符匹配（相同或为点号），则结果取决于之前的匹配情况
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// 如果当前字符是星号，有两种情况：
				// 1. 忽略星号和它前面的字符
				dp[i][j] = dp[i][j-2]
				// 2. 使用星号匹配当前字符
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					// 如果星号前面的字符匹配当前字符，我们可以选择匹配一个或多个
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
			// 如果以上条件都不满足，dp[i][j] 保持 false
		}
	}

	// 返回整个字符串和模式是否匹配
	return dp[m][n]
}
