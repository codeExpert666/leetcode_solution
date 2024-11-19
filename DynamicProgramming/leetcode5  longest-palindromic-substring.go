package DynamicProgramming

func LongestPalindrome(s string) string {
	var left, right int // 最长回文子串左右边界
	// 构建dp数组，dp[i][j]表示s[i:j+1]是否为回文串
	// 由于j>=i，dp数组本质是上三角矩阵
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	// 遍历dp数组递推
	for i := len(s) - 1; i >= 0; i-- { // 注意遍历顺序
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j+1-i > right-left { // 找到更长回文子串，更新左右边界
					left, right = i, j+1
				}
			}
		}
	}
	return s[left:right]
}
