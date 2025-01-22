package weeklycompetition

import (
	"math"
)

// 参考灵神题解：记忆化递归
// 记忆化递归可进一步写出递归公式，故而可以转化为动态规划。
// 二者本质上是一致的，掌握一种方法即可。
// 本题由于对称位置的房屋颜色不能相同，故可将对称位置的房屋看成一个整体，一次处理两个对称房屋
// 从最中间两个房屋开始，不断向两边遍历，枚举颜色，获得最小成本
func minCost(n int, cost [][]int) int64 {
	// 初始化记忆数组
	halfLength := n >> 1
	// 这里取 4 是因为递归入口状态的 color 会取到 3
	// 而单独区分该状态，使其不参与 memo 需要再加一个 if，导致所有的状态
	// 都要经过该 if 判断，得不偿失，不如空间换时间
	memo := make([][4][4]int, halfLength)
	for i := 0; i < halfLength; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				memo[i][j][k] = -1 // 初始值一定是无效值，0 可能是有效值
			}
		}
	}

	// 前一对房屋的颜色为 preJ, preK 时，房屋对 i, n-1-i 及之后所有房屋上色的最小成本
	var dfs func(int, int, int) int
	dfs = func(i int, preJ int, preK int) (res int) {
		if i == -1 { // 递归边界
			return
		}

		if memo[i][preJ][preK] != -1 { // 该状态已记录
			return memo[i][preJ][preK]
		}

		res = math.MaxInt
		for j, c1 := range cost[i] { // 枚举当前对房屋中的第一个房屋的颜色
			if j == preJ {
				continue
			}
			for k, c2 := range cost[n-1-i] { // 枚举当前对房屋中的第二个房屋的颜色
				if k != preK && k != j { // 非法涂色方案
					res = min(res, dfs(i-1, j, k)+c1+c2)
				}
			}
		}
		memo[i][preJ][preK] = res

		return
	}

	return int64(dfs(halfLength-1, 3, 3))
}
