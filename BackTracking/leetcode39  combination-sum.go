package BackTracking

import "sort"

/**
 * 回溯
 */
var (
	res3  [][]int
	path3 []int
	sum1  int
)

func combinationSum(candidates []int, target int) [][]int {
	res3, path3 = make([][]int, 0), make([]int, 0)
	backtracking2(candidates, target, 0)
	return res3
}

func backtracking2(candidates []int, target int, startPos int) {
	// 终止条件
	if sum1 >= target {
		if sum1 == target {
			tmp := make([]int, 0)
			tmp = append(tmp, path3...)
			res3 = append(res3, tmp)
		}
		// 剪枝
		return
	}
	for i := startPos; i < len(candidates); i++ {
		sum1 += candidates[i]
		path3 = append(path3, candidates[i])
		backtracking2(candidates, target, i)
		// 回溯
		path3 = path3[:len(path3)-1]
		sum1 -= candidates[i]
	}
}

/**
 * 回溯的优化版本，与上面回溯方法的区别在于先对候选数组排了序，这样可以进行剪枝操作
 */
func combinationSumCopy(candidates []int, target int) [][]int {
	res3, path3 = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates) // 排序，为剪枝做准备
	dfs1(candidates, 0, target)
	return res3
}

func dfs1(candidates []int, start int, target int) {
	if target == 0 { // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path3))
		copy(tmp, path3)
		res3 = append(res3, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝，提前返回
			break
		}
		path3 = append(path3, candidates[i])
		dfs1(candidates, i, target-candidates[i])
		path3 = path3[:len(path3)-1]
	}
}
