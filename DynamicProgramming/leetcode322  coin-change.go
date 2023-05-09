package DynamicProgramming

/**
 * 完全背包问题应用（dp）
 * 本题的难点在于初始化dp数组，因为递推公式需要用到min，为了不影响值的更新，初始值就要设的足够大
 * 这里最大不过amount取最大，只用1元硬币，硬币个数即为10000，故设成10001，不会影响值的更新
 * 此外，dp[0]自身不会递推，但为了不影响计算其余位置的正确值，需要设置成0
 * PS：本题求钱币最小个数，那么钱币有顺序和没有顺序都可以，都不影响钱币的最小个数。所以本题并不强调集合是组合还是排列。
 *    也就是说两种遍历顺序均可
 */
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 注意初始化
	for i := 1; i <= amount; i++ {
		dp[i] = 10001
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			// 递推公式
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	// 如若找不到硬币组合，dp[amount]会保持初始值
	if dp[amount] == 10001 {
		return -1
	}
	return dp[amount]
}
