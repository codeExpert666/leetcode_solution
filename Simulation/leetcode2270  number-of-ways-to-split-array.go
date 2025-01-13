package Simulation

func WaysToSplitArray(nums []int) int {
	// 初始化左右两部分的差值
	var dist int
	for _, v := range nums {
		dist -= v
	}

	// 不断遍历更新差值
	var res int
	for i := 0; i < len(nums)-1; i++ {
		dist += 2 * nums[i]
		if dist >= 0 {
			res++
		}
	}

	return res
}
