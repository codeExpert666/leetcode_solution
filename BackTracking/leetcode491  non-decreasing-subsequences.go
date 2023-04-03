package BackTracking

/**
 * 回溯
 */
var (
	res9  [][]int
	path9 []int
)

func findSubsequences(nums []int) [][]int {
	res9, path9 = make([][]int, 0), make([]int, 0, len(nums))
	backtracking8(nums, 0)
	return res9
}

func backtracking8(nums []int, startIndex int) {
	// 终止条件
	if len(path9) >= 2 {
		tmp := make([]int, len(path9))
		copy(tmp, path9)
		res9 = append(res9, tmp)
	}
	// 因为本题不能排序，去重的任务只能由hash表完成
	//used := make(map[int]bool, len(nums)) // ！！！很重要，用于对同层元素去重
	used := [201]int{} // 本题的元素范围是确定的，也可以使用数组来代替hash表，可以有效降低时间复杂度
	for i := startIndex; i < len(nums); i++ {
		//if used[nums[i]] {
		//	continue
		//}
		if used[nums[i]+100] == 1 {
			continue
		}
		if len(path9) == 0 || nums[i] >= path9[len(path9)-1] {
			// 该元素为路径第一个元素或满足递增要求时，加入路径
			path9 = append(path9, nums[i])
			backtracking8(nums, i+1)
			path9 = path9[:len(path9)-1]
			//used[nums[i]] = true // 已访问过，进行标记
			used[nums[i]+100] = 1
		}
	}
}
