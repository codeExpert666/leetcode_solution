package DailyCode

func ResultsArray2(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	count := 0
	for i, v := range nums {
		if i == 0 || v != nums[i-1]+1 {
			count = 1
		} else {
			count++
		}
		if count >= k {
			res[i-k+1] = nums[i]
		} else if i-k+1 >= 0 { // 连续递增元素不超过k个，也有可能是此时子数组总元素个数不超过k
			res[i-k+1] = -1
		}
	}
	return res
}
