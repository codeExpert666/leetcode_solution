package weeklycompetition

// 前缀和
func countPartitions(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i, v := range nums {
		prefixSum[i+1] = prefixSum[i] + v
	}

	var res int
	for i := 1; i < n; i++ {
		if (prefixSum[i]<<1-prefixSum[n])&1 == 0 {
			res++
		}
	}

	return res
}
