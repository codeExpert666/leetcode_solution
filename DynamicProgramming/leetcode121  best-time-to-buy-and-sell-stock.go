package DynamicProgramming

import "math"

/**
 * 方法一：暴力 O(n^2)
 * 也就是找到所有右侧减左侧值中的最大值
 */
func maxProfit(prices []int) int {
	var res int // 不初始化为最小值的原因是题目要求当亏损时需返回0，那不妨直接初始化为0，过程中不保留亏损情况

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			res = max(res, prices[j]-prices[i])
		}
	}
	return res
}

/**
 * 方法二：贪心 O(n)
 * 本质上是在暴力的基础上进行剪枝，考虑以下两种情况：
 * 1、当左侧固定时，遇到右侧有更小值时，记录完二者中间的所有差值后，就应将左侧调整到右侧当前位置（因为左侧的未来差值不可能比右侧未来差值大）
 * 2、当左侧固定时，右侧都比左侧大时，左侧无需移动，因为最大值只可能在当前左侧取到
 */
func maxProfit1(prices []int) int {
	low := math.MaxInt
	var res int

	for _, v := range prices {
		// 不断更新左端点，从而实现剪枝
		low = min(low, v)
		res = max(res, v-low)
	}

	return res
}

/**
 * 动态规划dp，利用滚动数组，节省空间
 * 每天的状态就两种：持有股票与不持有股票。故dp[i]代表每天的状态，进而不断递推
 * dp[i][0]代表持有股票的最大收益，dp[i][1]代表不持有股票的最大收益
 */
func maxProfit2(prices []int) int {
	// 初始化
	dp := [2]int{-prices[0], 0}

	for i := 1; i < len(prices); i++ {
		tmp := dp[0]
		// 注意：这样写是不对的，因为本题中的股票只买卖一次（先买后卖），第i天持有股票要么是之前就持有，要么是今天才持有
		// 而今天才持有收益一定是负数-price[i]，下面的写法实际上处理的是另一个问题“股票可买卖多次”，因为递推的过程会牵涉到以前买进卖出的过程
		// dp[0] = max(dp[0], dp[1]-prices[i])
		dp[0] = max(dp[0], -prices[i])
		dp[1] = max(dp[1], tmp+prices[i])
	}

	// 最后一天一定是不持有股票收益最大
	return dp[1]
}
