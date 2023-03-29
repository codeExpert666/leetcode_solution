package BackTracking

/**
 * 递归：从n个数中找k个数 = 从前n-1个数中找k个数 + 从前n-1个数中找k-1个数再与最后一个数拼接
 */
func combine(n int, k int) [][]int {
	// 终止条件
	if k == 1 {
		var res [][]int
		for i := 1; i <= n; i++ {
			res = append(res, []int{i})
		}
		return res
	}
	if n == k {
		var tmp []int
		for i := 1; i <= k; i++ {
			tmp = append(tmp, i)
		}
		return [][]int{tmp}
	}
	// 从前n-1个数中找k个数
	res := combine(n-1, k)
	// 从前n-1个数中找k-1个数
	tmp := combine(n-1, k-1)
	// 与最后一个数拼接
	for i := 0; i < len(tmp); i++ {
		tmp[i] = append(tmp[i], n)
	}
	// 整合结果
	res = append(res, tmp...)
	return res
}

/**
 * 回溯法（特殊的递归）：不断深度遍历，找到找不到都进行回溯
 * 如果把回溯过程看成是一个树结构，那么整个搜索过程就是“广度（for循环）”+“深度（递归）”的结合
 * 仔细体会回溯法以及剪枝操作
 */
var (
	path []int
	res  [][]int
)

func combine1(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n int, k int, start int) {
	if len(path) == k { // 说明已经满足了k个数的要求
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i <= n; i++ { // 从start开始，不往回走，避免出现重复组合
		if n-i+1 < k-len(path) { // 剪枝
			break
		}
		path = append(path, i)
		dfs(n, k, i+1)
		path = path[:len(path)-1]
	}
}
