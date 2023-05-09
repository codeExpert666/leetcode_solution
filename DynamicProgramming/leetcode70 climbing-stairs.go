package DynamicProgramming

/**
 * 动态规划，状态递推式为dp[i] = dp[i-1] + dp[i-2]
 * 本质上还是斐波那契数列
 */
func climbStairs(n int) int {
	if n <= 2 { // 可以一步走完的特殊处理
		return n
	}
	var dp [2]int // dp数组，第i项记录i层楼梯的走法
	dp[0], dp[1] = 1, 2
	for i := 3; i <= n; i++ {
		// 后续无法一步走完的楼梯的走法可分为两类：1、先走一层，再走剩余n-1层
		// 2、先走两层、再走剩余n-2层
		tmp := dp[1] + dp[0]
		dp[0] = dp[1]
		dp[1] = tmp
	}
	return dp[1]
}

/**
 * 还可以转化为完全背包问题（dp）
 * 用无限制的重量为1和2的物品装满承重为n的背包
 */
func climbStairs1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for j := 0; j <= n; j++ {
		for i := 1; i <= 2; i++ {
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}

	return dp[n]
}
