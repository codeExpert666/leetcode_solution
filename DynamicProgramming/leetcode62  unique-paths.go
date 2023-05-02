package DynamicProgramming

/**
 * 动态规划
 */
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	} // 这部分可以省略
	// dp数组，用于存储到达每个位置的路径条数
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// dp数组初始化
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 每个位置都可由左边位置右移一步到达，也可由上边位置下移一步到达
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}

/**
 * 动态规划优化版本：将二维数组优化为一维数组
 */
func uniquePaths1(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			dp[i] += dp[i-1]
		}
	}

	return dp[n-1]
}

/**
 * 还有一种方法：排列组合
 * 无论怎么走，走到终点都需要 m + n - 2 步，并且在这m + n - 2 步中，一定有 m - 1 步是要向下走的，不用管什么时候向下走
 * 这种方法看起来简单，但计算的过程并不简单，主要是会出现溢出的情况。整个过程的难点就在于对于溢出的处理
 * 实际上，计算组合问题的代码还是有难度的，特别是处理溢出的情况！
 * 这里不再对这种方法进行实现，有时间过来补一下 TODO
 */
