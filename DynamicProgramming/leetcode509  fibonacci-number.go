package DynamicProgramming

/**
 * 递归，由于利用了空间换时间的方式，故空间复杂度为O(2n) = O(n)(递归调用栈以及hash表)
 * 时间复杂度大致是O(n)，应为是递归，但又有优化，故不好计算
 */
var m = make(map[int]int) // 记录已经计算过的斐波那契值，避免重复计算

func fib(n int) int {
	if n <= 1 {
		m[n] = n
		return n
	}
	// 斐波那契数列前两项值
	tmp1, e1 := m[n-2] // 注意，e代表是否存在，key存在，则为true，不存在，则为false
	tmp2, e2 := m[n-1]
	if !e1 {
		tmp1 = fib(n - 2)
		m[n-2] = tmp1
	}
	if !e2 {
		tmp2 = fib(n - 1)
		m[n-1] = tmp2
	}
	// 当前项等于前两项之和
	return tmp1 + tmp2
}

/**
 * 动态规划：本来dp数组的长度应是n+1，每个位置记录对应的斐波那契数值，
 * 但在实际求解中，只需要知道前两个斐波那契数，故无需记录整体的斐波那契数列，从而节省空间
 * 这种方法时间复杂度为O(n)，空间复杂度为O(1)
 */
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	var dp [2]int // dp数组，在本题中用于记录前两个斐波那契数
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := dp[0] + dp[1]
		// 不断更新前两个数
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}
