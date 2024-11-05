package Array

func PivotIndex(nums []int) int {
	// 先完整遍历一遍获取数组总和
	var sum int
	for _, v := range nums {
		sum += v
	}
	// 再遍历一遍，不断比较左右两侧的元素和
	var leftSum int // 左侧元素和与左侧最近元素
	for i, v := range nums {
		if leftSum<<1 == sum-v { // 左右侧元素和相等
			return i
		}
		leftSum += v
	}
	return -1
}
