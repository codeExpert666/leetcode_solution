package DynamicProgramming

/**
 * 动态规划，本题有两个难点：一个是递推式，另一个是初始值
 * 数n的拆分可以分为如下情况：1 + (n-1)（拆分或不拆分）、2 + (n-2)（拆分或不拆分）、...、(n-1) + 1（拆分或不拆分）
 * 但从n/2后的拆分便没有意义了，因为拆分一个数n 使之乘积最大，那么一定是拆分成m个近似相同的子数相乘才是最大的。
 * 由上所述，递推式应为dp[i] = max({dp[i], (i - j) * j, dp[i - j] * j})，其中，dp[i]表示分拆数字i，可以得到的最大乘积为dp[i]
 * PS：遍历到n/2不再拆分还有一个好处，就是可以不为dp[0]、dp[1]赋初始值，因为访问不到。
 */
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1 // 初始化

	for i := 3; i <= n; i++ {
		dp[i] = -1
		for j := 1; j <= i/2; j++ {
			// 递推式
			dp[i] = max(max(dp[i], j*(i-j)), j*dp[i-j])
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
 * 贪心：每次拆成n个3，如果剩下是4，则保留4，然后相乘，但是这个结论需要数学证明其合理性！
 * 数学证明分为两步：1、证明只有当整数被拆分成相近的一些正整数时，乘积才可能取到最大（均值不等式）
 * 2、问题转化为拆分成几个相近数的问题，不妨设拆分成x个，则问题转化为求解y=(n/x)^x的最大值问题
 * 令t=n/x，则y=t^(n/t)，求得当t=e时，y取最大值，由于t为正整数，故t=2或t=3，其实从图像可以比较，应该是t=3时y更大
 * 也即，一个整数应尽量拆分成多个3
 */
func integerBreak1(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}
	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}
	res *= n
	return res
}
