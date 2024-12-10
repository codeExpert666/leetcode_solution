package DynamicProgramming

// 方法一：动态规划
// dp[i][j] 表示 s[i:j+1] 最少需要添加多少个字母能够变为回文串
func MinInsertions(s string) int {
	l := len(s)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if j-i <= 1 && s[j] != s[i] {
				dp[i][j] = 1
			} else if j-i > 1 {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
				if s[i] == s[j] {
					dp[i][j] = min(dp[i][j], dp[i+1][j-1])
				}
			}
		}
	}
	return dp[0][l-1]
}

// 方法二：添加最少的字符变成回文串 => 找出当前串的最大回文序列（序列是不连续的）=> 剩余字符均需要添加字符进行匹配
