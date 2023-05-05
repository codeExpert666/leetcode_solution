package DynamicProgramming

/**
 * 完全背包问题的应用（dp）
 * 本题将总金额看成背包总重，硬币面额看成重量和价值，由于本题是求有多少种，故初始化需注意设置dp[0]为1
 * 在本题的情况下，两层循环不可交换，本题本质上是一种组合问题，如若交换，求解的就是排列问题了
 * 本题中的循环很好理解，硬币按序考虑，得到的一定是组合数
 * 但循环交换后，由于先按列遍历，导致每个amount都会先对全部硬币面额进行考虑，并带入后续列中，这就导致
 * 同样的一些硬币面额会以各种顺序出现（之前只会按照遍历的顺序出现），从而形成了排列数
 */
func change(amount int, coins []int) int {
	dp := make([]int, amount+1) // dp数组及初始化
	dp[0] = 1
	// dp过程
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			// 递推公式
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}
