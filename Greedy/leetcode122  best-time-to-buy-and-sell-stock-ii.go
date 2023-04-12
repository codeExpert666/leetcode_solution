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
