package BackTracking

/**
 * 回溯
 */
var (
	res10  [][]int
	path10 []int
)

func permute(nums []int) [][]int {
	res10, path10 = make([][]int, 0), make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	backtracking9(nums, used)
	return res10
}

func backtracking9(nums []int, used []bool) { // used用于记录已经遍历过的元素
	// 终止条件
	if len(nums) == len(path10) {
		tmp := make([]int, len(path10))
		copy(tmp, path10)
		res10 = append(res10, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path10 = append(path10, nums[i])
		used[i] = true // 顺序很重要，一定要在递归前记录
		backtracking9(nums, used)
		path10 = path10[:len(path10)-1]
		used[i] = false // 标记数组也要回溯
	}
}
