package Simulation

// 灵神题解：贪心
// 任意顺序都能交易的最小金额 => 最差顺序能够交易的金额
// 这里的最差顺序并不是一个绝对的固定排列，而是一个排列规则（先做亏钱交易，再做赚钱交易）
// 最差顺序是针对每笔交易来说的，对于一个亏钱交易，当其作为最后一笔亏钱交易时，其所需要的起始资金是最大的
// 对于一个赚钱交易，当其作为亏欠后的第一笔赚钱交易时，其所需要的起始资金是最大的
// 于是，所求资金就是这些最大起始资金的最大值。因为该值可能满足任意一笔资金排列在任意顺序的情况。
// 具体公式计算见灵神题解 https://leetcode.cn/problems/minimum-money-required-before-transactions/solutions/1830862/by-endlesscheng-lvym/?envType=daily-question&envId=2025-01-25
func minimumMoney(transactions [][]int) int64 {
	totalLose, mx := 0, 0
	for _, t := range transactions {
		totalLose += max(t[0]-t[1], 0)
		mx = max(mx, min(t[0], t[1]))
	}
	return int64(totalLose + mx)
}
