package Greedy

/**
 * 本题思路过于简单，但仍然是贪心的思想：局部最优：收集每天的正利润，全局最优：求得最大利润。
 */
func maxProfit(prices []int) int {
	var profit int // 累积利润

	// 从第二项开始遍历，以计算价格差值
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			profit += tmp
		}
	}
	return profit
}

/**
 * 动态规划dp
 */
func maxProfit2(prices []int) int {
	// 初始化
	dp := [2]int{-prices[0], 0}

	for i := 1; i < len(prices); i++ {
		// 递推公式
		tmp := dp[0]
		dp[0] = max(dp[0], dp[1]-prices[i])
		//dp[0] = max(dp[0], -prices[i])  这是只能买卖一次股票的写法
		dp[1] = max(dp[1], tmp+prices[i])
	}

	// 最后一天一定是不持有股票收益最大
	return dp[1]
}
