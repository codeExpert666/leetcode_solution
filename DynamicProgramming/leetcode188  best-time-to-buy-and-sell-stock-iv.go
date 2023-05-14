package DynamicProgramming

/**
 * 动态规划
 * 本题是leetcode123的扩展，区别在于dp数组更长，涉及到的状态更多
 * 这些状态主要包括 1、没有进行任何操作 2、第i次买入股票 3、第i次卖出股票
 */
func maxProfit5(k int, prices []int) int {
	// 根据题目要求，price的长度可能为0，而长度为0的price数组会使下面代码出错
	// 故特殊处理，另外，k也可以为0，但k为0不影响程序的输出，故不做特殊处理
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, 2*k+1)
	for i := 1; i < len(dp); i += 2 { // dp初始化
		dp[i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		// 递推公式，由于状态太多，但有规律，故写成循环
		for j := k; j >= 1; j-- {
			index := 2 * j
			dp[index] = max(dp[index], dp[index-1]+prices[i])
			dp[index-1] = max(dp[index-1], dp[index-2]-prices[i])
		}
	}

	return dp[2*k]
}
