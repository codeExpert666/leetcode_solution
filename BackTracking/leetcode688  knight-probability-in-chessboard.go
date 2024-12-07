package BackTracking

// 记忆化回溯
func KnightProbability(n int, k int, row int, column int) float64 {
	rules := [8][2]int{
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	}
	// memo := make(map[[3]int]float64) // row, column, k -> 不出界走法
	memo := make([][][]float64, k+1) // 用三维数组替代 hash 表
	for i := 1; i <= k; i++ {
		memo[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]float64, n)
			for l := 0; l < n; l++ {
				memo[i][j][l] = -1
			}
		}
	}
	var dfs func(int, int, int) float64
	dfs = func(i, j, remainStep int) (goodMoves float64) { // 为了防止溢出，goodMoves 实际存储的是不出界的概率
		if i < 0 || i >= n || j < 0 || j >= n { // 走出棋盘
			return 0
		}
		if remainStep == 0 { // 步数用完
			return 1
		}
		p := &memo[remainStep][i][j]
		if *p != -1 { // 当前状态已搜索过
			return *p
		}
		// 状态未搜索，保存状态
		defer func() { *p = goodMoves }()
		// 遍历下一步可能位置
		for _, r := range rules {
			tmp := dfs(i+r[0], j+r[1], remainStep-1)
			goodMoves += tmp
		}
		goodMoves /= 8
		return
	}
	goodMoves := dfs(row, column, k)
	// fmt.Println(goodMoves)
	return goodMoves
}
