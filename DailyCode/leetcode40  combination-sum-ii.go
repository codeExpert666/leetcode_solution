package DailyCode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	var path []int
	var dfs func(int, int)
	dfs = func(s, t int) {
		if t == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		//if s == n { // 可以省略，因为这种情况下后续的循环会直接跳过，达到了直接return的效果
		//	return
		//}
		for i := s; i < len(candidates) && candidates[i] <= t; i++ {
			if i > s && candidates[i] == candidates[i-1] {
				continue // 剪枝
			}
			path = append(path, candidates[i])
			dfs(i+1, t-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)

	return res
}
