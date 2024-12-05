package DynamicProgramming

import "strconv"

// 暴力解法：枚举每个数字，对于每个数字统计 1 出现的次数
func countDigitOne(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		for cur := i; cur > 0; cur /= 10 {
			if cur%10 == 1 {
				res++
			}
		}
	}
	return res
}

// CountDigitOne1 数位dp：灵神题解链接：https://leetcode.cn/problems/number-of-digit-one/solutions/1750339/by-endlesscheng-h9ua/
// dfs 或者回溯本质上就是一种遍历，而数位 dp 本质上就是对 dfs（回溯）的众多分支进行归类记录，对于相同类别的分支，可以归为一个状态
// 同一状态只用计算一次即可，然后再利用 dfs（回溯）天然的倒推作为递推公式，就完成了求解
// 这种解法的难点还是在于对状态的定义，也即 dfs（回溯）的参数列表
func CountDigitOne1(n int) int {
	s := strconv.Itoa(n)
	l := len(s)
	// 定义状态数组
	// 只需要记录前面数字未到达极限时（i，count）组合所能构造出的数字中 1 的总数量
	// 同理，也只有前面数字未到达极限时才能访问该数组（记录或取值）
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, i+1) // 编号 i 的数字前面有 i 个数字，1 的个数取值范围是 0 到 i
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1 // 表示该状态还未记录
		}
	}
	var dfs func(int, int, bool) int
	// i 表示处理第 i 位数字，count 表示前 i 个数字中 1 的个数，isLimit 表示前 i 个数字是否都到达极限
	dfs = func(i int, count int, isLimit bool) (res int) {
		if i == l { // 退出条件
			return count
		}
		if !isLimit { // 前面所有数字还未到达极限
			p := &dp[i][count]
			if *p != -1 {
				res = *p
				return
			}
			defer func() { *p = res }() // 记录状态
		}
		up := 9      // 当前位的取值上限
		if isLimit { // 上限受到前面数字的影响
			up = int(s[i] - '0')
		}
		for j := 0; j <= up; j++ { // 枚举当前数字
			c := count
			if j == 1 { // 前 i 位数字中 1 的个数更新
				c++
			}
			res += dfs(i+1, c, isLimit && j == up)
		}
		return
	}
	return dfs(0, 0, true)
}
