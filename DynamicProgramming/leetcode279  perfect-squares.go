package DynamicProgramming

/**
 * 完全背包问题应用（dp）
 * 可以转化为将1到100无限制装入容量为n的背包中
 * 两点需要注意：一是dp数组的初始化，二是对于物品的遍历，本题需要自己确定物品列表
 */
func numSquares(n int) int {
	dp := make([]int, n+1)
	// 初始化
	for i := 1; i < len(dp); i++ {
		dp[i] = 10001
	}

	//for i := 1; i <= 100; i++ {
	//	for j := i * i; j <= n; j++ {
	//		// 递推公式
	//		dp[j] = min(dp[j], dp[j-i*i]+1)
	//	}
	//}
	for i := 1; i*i <= n; i++ { // 外层循环在上一版本的基础上做了剪枝，若i*i>n。则背包必不可能包含i
		for j := i * i; j <= n; j++ {
			// 递推公式
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	//if dp[n] == 10001 {  不需要，因为任意一个正整数都能写成多个完全平方数的和
	//	return 0
	//}
	return dp[n]
}

/**
 * 本题还可以用纯数学的方法解决，四平方定理，这里不再赘述，后续补充
 */
