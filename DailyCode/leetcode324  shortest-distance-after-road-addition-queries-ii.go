package DailyCode

// 时间复杂度为O(n*q)，过高，提交超时
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	res := make([]int, len(queries))
	dp := make([]int, n)      // dp[i]表示从节点0到节点i的最短路径的长度
	parents := make([]int, n) // 标记dp[i]最短路径中节点i的父节点
	// 初始化
	for i := 0; i < n; i++ {
		dp[i], parents[i] = i, i-1
	}
	// 每添加一条边，更新parents和dp的状态
	for i, q := range queries {
		from, to := q[0], q[1]
		// 若from > parents[to]，则新路径不如先前路径，没必要更新
		// 若from < parents[to]，本题添加的边不会交叉，那么有from <= parents[parents[to]]，
		// 显然添加的路径是到达to的更近路径
		if from < parents[to] {
			parents[to] = from    // 更新to节点的父节点
			dp[to] = dp[from] + 1 // 更新to节点的最短路径长度
			// 更新to后续节点的最短路径长度
			for j := to + 1; j < n; j++ {
				// 由于路径不交叉，to后续节点的父节点要么在to及之后，要么在from及之前
				if parents[j] >= to { // 父节点在to及之后的需更新
					dp[j] = dp[parents[j]] + 1
				} else { // 一旦遇到节点的父节点在from之前，后续不用更新
					break
				}
			}
		}
		res[i] = dp[n-1]
	}
	return res
}

// 官方题解：https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-ii/solutions/2984427/xin-zeng-dao-lu-cha-xun-hou-de-zui-duan-1n4hq/?envType=daily-question&envId=2024-11-20
// 时间复杂度为O(n+q)，官解思路很好，但是代码有误，更正如下：
func shortestDistanceAfterQueries1(n int, queries [][]int) []int {
	roads := make([]int, n) // 记录起点到终点的最短路径中每个节点的子节点
	for i := 0; i < n; i++ {
		roads[i] = i + 1
	}
	var res []int
	dist := n - 1 // 记录起点到终点最短路径长度
	for _, query := range queries {
		if k := roads[query[0]]; k != -1 && k < query[1] { // 路径不被包含
			roads[query[0]] = query[1]    // query[0]原子节点被覆盖，dist不变
			for k != -1 && k < query[1] { // 逐步清空路径所包含的更短路径
				roads[k], k = -1, roads[k]
				dist--
			}
		}
		res = append(res, dist)
	}
	return res
}
