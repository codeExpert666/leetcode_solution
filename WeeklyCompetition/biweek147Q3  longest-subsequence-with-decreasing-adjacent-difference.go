package weeklycompetition

// dp[i][j] 表示尾部为（ nums[j], nums[i] ）时，子序列的最大长度
func longestSubsequence(nums []int) int {
	var res, l = 2, len(nums)

	dp := make([][]*dpElem, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]*dpElem, l)
	}
	dp[0][0] = &dpElem{1, 0}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			dp[i][j] = &dpElem{2, absDis(nums[j], nums[i])}
			for k := 0; k < j; k++ {
				if dp[j][k].lasDis >= dp[i][j].lasDis {
					dp[i][j].maxLen = max(dp[i][j].maxLen, dp[j][k].maxLen+1)
					res = max(res, dp[i][j].maxLen)
				}
			}
		}
	}

	return res
}

func absDis(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

type dpElem struct {
	maxLen int
	lasDis int
}
