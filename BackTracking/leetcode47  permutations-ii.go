package BackTracking

import (
	"sort"
	_ "strconv"
)

/**
 * 回溯(需要去重)
 */
var (
	res11  [][]int
	path11 []int
)

func permuteUnique(nums []int) [][]int {
	res11, path11 = make([][]int, 0), make([]int, 0, len(nums))
	backtracking10(nums, make([]bool, len(nums)))
	return res11
}

func backtracking10(nums []int, used []bool) { // used用在不同层，用于全排列，访问过的元素不再访问，该数组利用下标访问
	// 终止条件
	if len(path11) == len(nums) {
		tmp := make([]int, len(path11))
		copy(tmp, path11)
		res11 = append(res11, tmp)
		return
	}

	used1 := [21]bool{} // used1用在同一层，去重，注意，该数组利用元素值访问
	for i := 0; i < len(nums); i++ {
		index := nums[i] + 10
		if !used1[index] && !used[i] {
			path11 = append(path11, nums[i])
			used1[index] = true
			used[i] = true
			backtracking10(nums, used)
			used[i] = false // 只有used需要回溯
			path11 = path11[:len(path11)-1]
		}
	}
}

/**
 * 先排序（去重效率高）再回溯
 */
func permuteUnique1(nums []int) [][]int {
	sort.Ints(nums) // 先排序
	res11, path11 = make([][]int, 0), make([]int, 0, len(nums))
	bp(nums, make([]bool, len(nums)))
	return res11
}

func bp(nums []int, used []bool) {
	if len(path11) == len(nums) {
		tmp := make([]int, len(path11))
		copy(tmp, path11)
		res11 = append(res11, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] { // 注意这个去重操作
			// ！！！上一结点存在并与当前值相等，且在全排列过程中未被遍历过
			continue
		}
		if !used[i] {
			path11 = append(path11, nums[i])
			used[i] = true
			bp(nums, used)
			used[i] = false
			path11 = path11[:len(path11)-1]
		}
	}
}
