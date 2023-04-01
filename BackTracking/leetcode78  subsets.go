package BackTracking

/**
 * 回溯
 */
var (
	res7  [][]int
	path7 []int
)

func subsets(nums []int) [][]int {
	res7, path7 = make([][]int, 0), make([]int, 0, len(nums))
	backtracking6(nums, 0)
	return res7
}

func backtracking6(nums []int, index int) {
	// 终止条件
	if index == len(nums) {
		tmp := make([]int, len(path7))
		tmp = append(tmp, path7...)
		// copy(tmp, path7)  使用copy方法时一定要保证源切片与目标切片长度len相同
		res7 = append(res7, tmp)
		return
	}
	for i := 0; i < 2; i++ { // 每个位置的元素都有两种可能，在或不在子集中
		if i == 0 {
			// 在子集中
			path7 = append(path7, nums[index])
			backtracking6(nums, index+1)
			path7 = path7[:len(path7)-1]
		} else {
			// 不在子集中
			backtracking6(nums, index+1)
		}
	}
}

/**
 * 回溯二：更精妙的回溯法，不设置终止条件，但程序不会陷入死循环，反而会遍历所有正确结果
 * 仔细体会（模拟一下就明白了）这种方法，非常简洁
 */
func subsets1(nums []int) [][]int {
	res7, path7 = make([][]int, 0), make([]int, 0, len(nums))
	dfs3(nums, 0)
	return res7
}

func dfs3(nums []int, start int) {
	tmp := make([]int, len(path7))
	copy(tmp, path7)
	res7 = append(res7, tmp)

	for i := start; i < len(nums); i++ {
		path7 = append(path7, nums[i])
		dfs3(nums, i+1)
		path7 = path7[:len(path7)-1]
	}
}
