package Greedy

import "slices"

// 贪心：优先沿花费高的切割线进行切割，这样可以尽可能减少高花费切割次数
func MinimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	// 降序排列切割费用
	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	hi, vi := 0, 0 // 分别指向下一个待使用的水平/垂直切割线
	// 按照水平/垂直切割线的公共降序序列进行切割
	var res int
	// 优化1: 使用更清晰的变量命名
	// 优化2: 合并循环,减少重复代码
	// 优化3: 提前计算切割次数,避免重复计算
	for hi < m-1 || vi < n-1 {
		if vi == n-1 || (hi < m-1 && horizontalCut[hi] >= verticalCut[vi]) {
			// 水平切割,当前有vi+1个垂直分块
			res += horizontalCut[hi] * (vi + 1)
			hi++
		} else {
			// 垂直切割,当前有hi+1个水平分块
			res += verticalCut[vi] * (hi + 1)
			vi++
		}
	}
	return res
}
