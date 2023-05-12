package DynamicProgramming

/**
 * 动态规划dp
 * 由于最多买卖两次，故直接暴力套用一次买卖股票的代码，按照两次买卖的分割点遍历一次价格数组
 * 对两部分分别套用买卖一次股票代码
 * 很遗憾，超时了
 */
func maxProfit3(prices []int) int {
	var res int
	for i := 0; i < len(prices); i++ {
		res = max(res, part(prices[:i])+part(prices[i:]))
	}
	return res
}

// 只能买卖一次股票的代码
func part(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
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

/**
 * 真 动态规划 滚动数组
 * 一天一共就有五个状态:
 * 1、没有进行过股票买卖操作 （其实我们也可以不设置这个状态）
 * 2、第一次持有股票
 * 3、第一次不持有股票
 * 4、第二次持有股票
 * 5、第二次不持有股票
 * dp[i][j]中 i表示第i天，j为 [0 - 4] 五个状态，dp[i][j]表示第i天状态j所剩最大现金。
 * Tips：动态规划问题在设计dp数组的时候一定要考虑每个子问题有哪些可能性，或者说有哪些状态
 *       状态列全了，剩下就是对这些状态递推
 */
func maxProfit4(prices []int) int {
	// 这里的初始化需要注意：第一天的第一次卖出可以理解为当天买了，然后又卖出了，所以是0
	// 第一天的第二次买入，可以理解为先买入，再卖出，再买入，剩下一个同理
	dp := [5]int{0, -prices[0], 0, -prices[0], 0}

	for i := 1; i < len(prices); i++ {
		// 递推公式，dp[0]无需改变
		dp[4] = max(dp[4], dp[3]+prices[i])
		dp[3] = max(dp[3], dp[2]-prices[i])
		dp[2] = max(dp[2], dp[1]+prices[i])
		dp[1] = max(dp[1], dp[0]-prices[i])
	}

	//return max(dp[2], dp[4])
	// 代码可以优化
	// 现在最大的时候一定是卖出的状态，而两次卖出的状态现金最大一定是最后一次卖出。如果想不明白的录友也可以这么理解：
	// 如果第一次卖出已经是最大值了，那么我们可以在当天立刻买入再立刻卖出。
	// 所以dp[4][4]已经包含了dp[4][2]的情况。也就是说第二次卖出手里所剩的钱一定是最多的。
	// 但是我仍然不是很理解这段话（好像题目确实允许当天买当天卖）
	return dp[4]
}
