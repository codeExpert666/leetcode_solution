package BackTracking

import "sort"

/**
 * 回溯
 */
var (
	res4  [][]int
	path4 []int
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 先对数组进行排序
	res4, path4 = make([][]int, 0), make([]int, 0)
	backtracking3(candidates, target, 0)
	return res4

}

func backtracking3(candidates []int, target int, startIndex int) {
	// 终止条件
	if target == 0 {
		tmp := make([]int, len(path4))
		copy(tmp, path4)
		res4 = append(res4, tmp)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝
			break
		}
		if i > startIndex && candidates[i] == candidates[i-1] { // 剪枝，因为candidate是有序的，所以相同的元素应该在一起
			continue
		}
		path4 = append(path4, candidates[i])
		backtracking3(candidates, target-candidates[i], i+1)
		path4 = path4[:len(path4)-1]
	}
}
