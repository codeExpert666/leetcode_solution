package DailyCode

// ShortestDistanceAfterQueries 动态规划，复杂度较高，没控制好。我这里多计算了一些题目不需要的内容
// 题目只用求起点到终点的最短距离，我的dp数组给出了任意两个节点的最近距离，所以复杂的高了
func ShortestDistanceAfterQueries(n int, queries [][]int) []int {
	// 初始化每两个城市间的最短距离
	minDist := make([][]int, n)
	for i := 0; i < n; i++ {
		minDist[i] = make([]int, n)
		for j := i; j < n; j++ {
			minDist[i][j] = j - i
		}
	}
	// 每添加一条边，更新minDist数组
	var res []int
	for _, query := range queries {
		from, to := query[0], query[1]
		// 更新城市0,...,from到城市to,...,n-1的最短距离
		for i := 0; i <= from; i++ {
			for j := to; j < n; j++ {
				minDist[i][j] = min(minDist[i][j], minDist[i][from]+1+minDist[to][j])
			}
		}
		res = append(res, minDist[0][n-1])
	}
	return res
}

// ShortestDistanceAfterQueries1 复杂度更低的动态规划：题目要求起点到终点的最短距离，这种方法的dp数组给出了起点到任意点的最短距离，所以复杂度更低。
// 根据题意，对于任一单向道路的起始点 u，终止点 v，都有 u<v，那么从城市 0 到任一城市的路径上，所经过的城市编号是单调递增的。
// 令 dp[i] 表示城市 0 到城市 i 的最短路径，同时使用 prev[i] 记录通往城市 i 的所有单向道路的起始城市集合，那么对于 i>0，有 dp[i]=min(dp[j]+1)
// 其中j∈prev[i]。根据以上推论，我们可以遍历 queries，在每次查询时，更新 prev 数组，然后更新 dp 数组。
// 注意到，每次新建一条从城市 u 到城市 v 的单向道路时，只有 i≥v 的 dp[i] 会发生变化，因此更新 dp 可以从 v 开始更新。
func ShortestDistanceAfterQueries1(n int, queries [][]int) []int {
	prev := make([][]int, n)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		prev[i] = append(prev[i], i-1)
		dp[i] = i
	}
	var res []int
	for _, query := range queries {
		prev[query[1]] = append(prev[query[1]], query[0])
		for v := query[1]; v < n; v++ {
			for _, u := range prev[v] {
				dp[v] = min(dp[v], dp[u]+1)
			}
		}
		res = append(res, dp[n-1])
	}
	return res
}
