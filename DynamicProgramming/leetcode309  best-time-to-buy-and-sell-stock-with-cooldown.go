package DynamicProgramming

/**
 * 动态规划
 * 每一天有两种状态：买入股票、卖出股票
 */
func maxProfit6(prices []int) int {
	// dp初始化
	//dp := [2]int{max(-prices[0], -prices[1]), max(prices[1]-prices[0], 0)} // 第二天的状态
	dp := [2]int{-prices[0], 0} // 第一天的状态
	var dpValue int             // 第零天的卖出状态

	for i := 1; i < len(prices); i++ {
		// 递推公式，当天的状态需要依靠前一天的状态和前两天的卖出状态
		tmp := dp[1]
		dp[1] = max(dp[1], dp[0]+prices[i])
		dp[0] = max(dp[0], dpValue-prices[i])
		dpValue = tmp
	}

	return dp[1]
}
