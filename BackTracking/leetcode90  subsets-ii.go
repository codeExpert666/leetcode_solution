package BackTracking

import (
	"sort"
)

/**
 * 回溯
 */
var (
	res8  [][]int
	path8 []int
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // 先排序
	res8, path8 = make([][]int, 0), make([]int, 0, len(nums))
	backtracking7(nums, 0)
	return res8
}

func backtracking7(nums []int, startIndex int) {
	tmp := make([]int, len(path8))
	copy(tmp, path8)
	res8 = append(res8, tmp)

	for i := startIndex; i < len(nums); i++ {
		if i > startIndex && nums[i] == nums[i-1] { // 去重剪枝
			continue
		}
		path8 = append(path8, nums[i])
		backtracking7(nums, i+1)
		path8 = path8[:len(path8)-1]
	}
}
