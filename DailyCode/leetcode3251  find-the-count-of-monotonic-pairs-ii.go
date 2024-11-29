package DailyCode

func countOfPairs1(nums []int) int {
	mod := int(1e9 + 7)
	n := len(nums)
	m := 0
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	dp := make([]int, m+1)
	tmp := make([]int, m+1)
	for i := 0; i <= nums[0]; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		copy(tmp, dp)
		clear(dp)
		d := max(0, nums[i]-nums[i-1])
		for j := d; j <= nums[i]; j++ {
			if j == 0 {
				dp[j] = tmp[j]
			} else {
				dp[j] = (dp[j-1] + tmp[j-d]) % mod
			}
		}
	}
	var res int
	for i := 0; i <= nums[n-1]; i++ {
		res = (res + dp[i]) % mod
	}
	return res % mod
}
