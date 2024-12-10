package BackTracking

// 记忆化回溯
func KnightDialer(n int) int {
	mod := int(1e9 + 7)
	// 每个号码的邻居
	nextNum := [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 3, 9}, {},
		{0, 1, 7}, {2, 6}, {1, 3}, {2, 4}}
	// 状态记忆
	memo := make([][]int, n+1)
	for i := 1; i < n; i++ {
		memo[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			memo[i][j] = -1 // 初始化
		}
	}
	// 已经走了 i 步，上一步走的是 num，后续还能形成的电话号码个数
	var dfs func(int, int) int
	dfs = func(i, num int) (res int) {
		if i == n { // 当前号码本身
			return 1
		}
		p := &memo[i][num]
		if *p != -1 { // 该状态已遍历过
			return *p
		}
		defer func() { *p = res }() // 保存状态
		// 遍历邻居
		for _, next := range nextNum[num] {
			res += dfs(i+1, next)
		}
		res %= mod
		return
	}
	// 获取所有结果
	var res int
	for n := range 10 {
		res += dfs(1, n)
	}
	return res % mod
}
