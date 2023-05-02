package DynamicProgramming

/**
 * 动态规划，注意，这里dp[i]的定义：到达第i+1台阶所花费的最少体力为dp[i]。
 */
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 { // 可以一步到达的特殊处理
		return min(cost[0], cost[1])
	}
	var dp [2]int                           // 记录到达前两层楼梯的最低花费
	dp[0], dp[1] = 0, min(cost[0], cost[1]) // 注意dp[0]的初始值为0，这是因为题目明确提出可以从下标为1的台阶开始爬楼梯，故无需任何花费

	for i := 2; i < len(cost); i++ {
		// 到达每一层楼梯只有两种情况：1、由前一层楼梯爬一步到达 2、由前前楼梯一次性爬两步到达
		// 当前楼梯的花费取上述两种情况的最小值
		tmp := min(dp[0]+cost[i-1], dp[1]+cost[i])
		dp[0] = dp[1]
		dp[1] = tmp
	}
	return dp[1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
