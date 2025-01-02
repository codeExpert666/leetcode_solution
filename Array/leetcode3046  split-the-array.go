package Array

// 问题等价于数组中元素的重复次数不能大于 2
func IsPossibleToSplit(nums []int) bool {
	count := [101]byte{}
	for _, v := range nums {
		count[v]++
		if count[v] > 2 {
			return false
		}
	}
	return true
}
