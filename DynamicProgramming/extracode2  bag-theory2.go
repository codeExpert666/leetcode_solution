package DynamicProgramming

/**
 * 先遍历物品, 再遍历背包实现完全背包问题
 * 完全背包描述如下：有N件物品和一个最多能背重量为W的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。
 * 每件物品都有无限个（也就是可以放入背包多次），求解将哪些物品装入背包里物品价值总和最大。
 * 完全背包和01背包问题唯一不同的地方就是，每种物品有无限件
 * 二者在实现代码上的区别只有遍历顺序不同！！！
 */
func testCompletePack1(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ {
			// 推导公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

/*
 * 先遍历背包, 在遍历物品实现完全背包问题
 * 这种遍历顺序不太符合常理，但也是对的，与上一种遍历方式的理解有很大不同
 */
func testCompletePack2(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	// j从0 开始
	for j := 0; j <= bagWeight; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] {
				// 推导公式
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}
