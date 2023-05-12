package main

import "fmt"

func maxProfit4(prices []int) int {
	// 这里的初始化需要注意：第一天的第一次卖出可以理解为当天买了，然后又卖出了，所以是0
	// 第一天的第二次买入，可以理解为先买入，再卖出，再买入，剩下一个同理
	dp := [5]int{0, -prices[0], 0, -prices[0], 0}
	fmt.Println(dp)

	for i := 1; i < len(prices); i++ {
		// 递推公式，dp[0]无需改变
		dp[4] = max(dp[4], dp[3]+prices[i])
		dp[3] = max(dp[3], dp[2]-prices[i])
		dp[2] = max(dp[2], dp[1]+prices[i])
		dp[1] = max(dp[1], dp[0]-prices[i])
		fmt.Println(dp)
	}

	return max(max(dp[2], dp[4]), dp[0])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProfit4([]int{7, 6, 5, 4, 3, 1, 7}))
}
