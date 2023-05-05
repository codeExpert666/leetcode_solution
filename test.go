package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1) // dp数组及初始化
	dp[0] = 1
	// dp过程
	for j := 0; j <= amount; j++ { // 遍历背包容量
		for i := 0; i < len(coins); i++ { // 遍历物品
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
			fmt.Println(dp)
		}

	}

	return dp[amount]
}

func change1(amount int, coins []int) int {
	dp := make([]int, amount+1) // dp数组及初始化
	dp[0] = 1
	// dp过程
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			// 递推公式
			dp[j] += dp[j-coins[i]]
			fmt.Println(dp)
		}
	}

	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Print(change1(5, []int{1, 2, 5}))
}
