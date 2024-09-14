package Array

// 相向双指针，原理：一串非递减数字的最大平方值只可能出现在两个端点上，
// 且去掉最大值剩下的这串数字仍然非递减。
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	for low, high, ins := 0, len(nums)-1, len(res)-1; low <= high; ins-- {
		lowVal, highVal := nums[low]*nums[low], nums[high]*nums[high]
		if lowVal >= highVal {
			res[ins] = lowVal
			low++
		} else {
			res[ins] = highVal
			high--
		}
	}
	return res
}
