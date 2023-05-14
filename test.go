package main

import "fmt"

func maxProfit7(prices []int, fee int) int {
	// dp初始化
	dp := [2]int{-prices[0] - fee, 0}
	fmt.Println(dp)

	for i := 1; i < len(prices); i++ {
		// 递推公式，注意每一笔交易都有费用
		tmp := dp[0]
		dp[0] = max(dp[0], dp[1]-prices[i]-fee)
		dp[1] = max(dp[1], tmp+prices[i])
		fmt.Println(dp)
	}
	// 这里dp[1]可能为负值，代表不可能赚钱，按照题目要求，应返回0
	if dp[1] < 0 {
		return 0
	}
	return dp[1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProfit7([]int{2, 1, 4, 4, 2, 3, 2, 5, 1, 2}, 1))
}
