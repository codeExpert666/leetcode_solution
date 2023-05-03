package DynamicProgramming

/**
 * 动态规划，1到n组成的二叉搜索数只有n种情况：n个节点分别作为根节点
 * 节点i作为根节点时，左子树i-1个节点，右子树有n-i个节点
 * 故递推公式为：dp[i] = dp[0]*dp[n-1] + dp[1]*dp[n-2] + ... + dp[n-1]*dp[0]
 * dp[i]代表1到i组成的二叉搜索数的种数
 * PS：注意dp[0]的初值，应设为1而不是0，应为后面有乘积操作，0会导致运算出错，0可以理解成空树，空树也是一种树
 */
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1 // 初始化

	for i := 2; i <= n; i++ {
		// 递推公式
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
