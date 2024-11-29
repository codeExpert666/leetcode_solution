package DailyCode

// 二维数组动态规划，dp[i][j] 表示 arr1 的第 i 项为 j 时，nums[:i+1] 所能构成的单调数组对数
// 见官方题解：https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-i/solutions/2992035/dan-diao-shu-zu-dui-de-shu-mu-i-by-leetc-7x5r/?envType=daily-question&envId=2024-11-28
// 以下为为 dp 数组优化为一维的版本
func countOfPairs(nums []int) int {
	n := len(nums)
	m := 0
	mod := int(1e9 + 7)
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	dp := make([]int, m+1)
	tmpDp := make([]int, m+1) // 备份
	for i := 0; i <= nums[0]; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		copy(tmpDp, dp)
		clear(dp)
		d := max(0, nums[i]-nums[i-1])
		for j := d; j <= nums[i]; j++ {
			if j == 0 {
				dp[j] = tmpDp[j-d]
			} else {
				dp[j] = (dp[j-1] + tmpDp[j-d]) % mod // 防止溢出
			}
		}
	}
	// 返回结果
	var res int
	for i := 0; i <= nums[n-1]; i++ {
		res = (res + dp[i]) % mod // 防止溢出
	}
	return res
}
