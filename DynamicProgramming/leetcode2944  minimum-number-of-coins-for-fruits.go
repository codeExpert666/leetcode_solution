package DynamicProgramming

import "slices"

// 参考灵神题解，dp[i] 表示在购买第 i 个水果的条件下，获取 i 及之后所有水果的最少金币数
// 购买第 i 个水果后，可以再购买第 i+1 个水果，此时花费金币 prices[i]+dp[i+1]
// 也可以不购买第 i+1 个水果，购买后续水果。但第 2i+1 个水果是一定要购买的。最终 dp[i] 取这些情况的最小值
// 需要注意的是：当 2i >= n 时，也即 i >= (n+1) / 2 时，dp[i] = prices[i]。
// 于是，可以直接利用 prices 作为 dp 数组
func minimumCoins(prices []int) int {
	n := len(prices)

	for i := (n+1)>>1 - 1; i > 0; i-- {
		prices[i-1] += slices.Min(prices[i : 2*i+1])
	}

	return prices[0]
}

// 灵神题解：基于 dp 的滑动窗口最小值问题。代码太优雅了，欣赏就完事了！
// 由于每次计算 dp[i] 都需要 [i+1:2i+2] 范围内的最小值，且该区间随着 i 的减小也会逐渐缩短并向左移动。
// 故该问题可转化为滑动窗口中的最小值问题，该问题的解决方式时单调双端队列（与单调栈相似，只是说多了总容量的限制）
func minimumCoins1(prices []int) int {
	n := len(prices)
	type pair struct{ i, f int }
	// 仔细体会这个哨兵的设计，既能保证 2i >= n 时 dp[i] = prices[i]（因为哨兵在这一段时间内一定不会被剔除，且始终在 q[0]，且其值为 0）
	// 又能在 2i < n 时及时从队列中移除
	q := []pair{{n + 1, 0}} // 哨兵
	for i := n; i > 0; i-- {
		for q[0].i > i*2+1 { // 右边离开窗口
			q = q[1:]
		}
		// 到这里为止 q 中存储的便是水果 i 的 dp[i] 所需的完整窗口
		f := prices[i-1] + q[0].f
		for f <= q[len(q)-1].f { // 这里不需要判断 len(q) > 0 的原因是 f 一定大于 q[0].f，队列不可能为空。
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, f}) // 左边进入窗口
	}
	return q[len(q)-1].f
}
