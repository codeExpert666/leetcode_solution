package DynamicProgramming

/**
 * 二维dp数组实现01背包
 * 01背包问题可以这样描述：有n件物品和一个最多能背重量为w 的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。
 * 每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。
 */
func test2WeiBagProblem1(weight, value []int, bagWeight int) int {
	// 定义dp数组
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagWeight+1)
	}
	// 初始化
	for j := bagWeight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// 递推公式
	for i := 1; i < len(weight); i++ {
		//正序,也可以倒序
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagWeight]
}

/**
 * 一维数组实现01背包
 * 一维数组要想实现二维数组的效果，就必须不停的滚动覆盖，更新相应位置的值
 * 在二维数组中，递推公式中当前值的求解都是利用了上一层的值，故一维数组只要可以保证其
 * 存储的数据一直是上一层的值即可，这也是一维数组实现过程中倒序遍历的原因
 */
func test1WeiBagProblem(weight, value []int, bagWeight int) int {
	// 定义 and 初始化
	dp := make([]int, bagWeight+1)
	// 递推顺序
	for i := 0; i < len(weight); i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j := bagWeight; j >= weight[i]; j-- {
			// 递推公式
			// 注意：原二维数组中的第一层的值的更新恰好也满足以下递推公式，故不用单独处理
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	//fmt.Println(dp)
	return dp[bagWeight]
}
