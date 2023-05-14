package DynamicProgramming

/**
 * 动态规划
 * 同之前的股票题一样，每一天有两种状态：1、买入了股票 2、卖出了股票
 * 注意将每次交易的费用加在递推公式中（买入时收取交易费用）
 */
func maxProfit7(prices []int, fee int) int {
	// dp初始化，这里注意：dp[1]不能初始化为-fee，虽然dp[1]理解为当天先买入再卖出
	// 但毕竟没有真实的买入操作，不能初始化为-fee。如若初始化为-fee，在递推的过程中实际上会重复计算
	dp := [2]int{-prices[0] - fee, 0}

	for i := 1; i < len(prices); i++ {
		// 递推公式，注意每一笔交易都有费用
		tmp := dp[0]
		dp[0] = max(dp[0], dp[1]-prices[i]-fee)
		dp[1] = max(dp[1], tmp+prices[i])
	}
	// 初始化中有0，故不可能再出现负值，可省略这一步
	//// 这里dp[1]可能为负值，代表不可能赚钱，按照题目要求，应返回0
	//if dp[1] < 0 {
	//	return 0
	//}
	return dp[1]
}
