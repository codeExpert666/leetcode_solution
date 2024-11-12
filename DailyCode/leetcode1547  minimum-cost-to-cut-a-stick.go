package DailyCode

import (
	"math"
	"sort"
)

// 获取当前切割的成本
func getPartCost(n int, i int, used []bool, cuts []int) int {
	left, right := 0, n
	for j := i - 1; j >= 0; j-- { //找到被切割棍子的左端点
		if used[j] {
			left = cuts[j]
			break
		}
	}
	for j := i + 1; j < len(used); j++ {
		if used[j] {
			right = cuts[j]
			break
		}
	}
	return right - left
}

// 回溯：思路正确，但超时了，复杂度太高
func MinCost(n int, cuts []int) int {
	res, cost := math.MaxInt, 0     // cost是当前成本
	sort.Ints(cuts)                 // 便于寻找最近的切开位置
	used := make([]bool, len(cuts)) // 指示切开位置是否使用
	var backTracking func(int)
	backTracking = func(count int) { // count表示切开位置被使用的个数
		if count == len(cuts) {
			if cost < res {
				res = cost
			}
			return
		}
		for i := 0; i < len(cuts); i++ {
			if !used[i] { // 仅考虑未被使用的切开位置
				partCost := getPartCost(n, i, used, cuts)
				cost += partCost
				used[i] = true
				backTracking(count + 1)
				used[i] = false
				cost -= partCost
			}
		}
	}
	backTracking(0)
	return res
}

// 方法二：动态规划，很明显，一个大问题可以拆解为若干子问题
// 二维dp（上三角），dp[i][j]表示棍子（cuts[i-1]至cuts[j+1]）在cuts[i:j+1]所指示的切割位置下的最小成本
func MinCost1(n int, cuts []int) int {
	sort.Ints(cuts) // 方便获取一段棍子的所有切割位置
	cuts = append([]int{0}, cuts...)
	cuts = append(cuts, n) // 一段棍子的所有可能端点
	// 这里dp数组多加两行和一列的原因是dp[i][k-1]与dp[k+1][j]的访问会超出dp数组的有意义范围
	dp := make([][]int, len(cuts))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(cuts)-1)
	}
	for i := len(dp) - 2; i > 0; i-- { // 注意遍历顺序
		for j := i; j < len(cuts)-1; j++ {
			if j == i { // 棍子上只有一个切割位置
				dp[i][j] = cuts[i+1] - cuts[i-1]
			} else {
				dp[i][j] = math.MaxInt
				for k := i; k <= j; k++ { // 获取左右两段棍子的最小切割成本
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
				}
				dp[i][j] += cuts[j+1] - cuts[i-1] // 加上自身切割成本
			}
		}
	}
	return dp[1][len(cuts)-2]
}
