package DynamicProgramming

/**
 * 本题也是动态规划解决子序列问题的经典问题
 * 本题的难点同样还是dp数组的意义，这里dp数组定义为二维数组，dp[i][j]代表以nums1[i-1]结尾的子序列与以nums2[j-1]结尾的子序列的最长重复长度
 * 当然，可以利用滚动数组节省空间
 */
func findLength(nums1 []int, nums2 []int) int {
	var res int
	// 这里初始化无需特殊处理，为0即可
	dp := make([]int, len(nums2)+1)

	for i := 0; i < len(nums1); i++ {
		// 倒序更新，避免之前的值被覆盖
		for j := len(nums2); j >= 1; j-- {
			if nums1[i] == nums2[j-1] {
				// 递推公式
				dp[j] = dp[j-1] + 1
			} else { // 注意：这一步不可省略，因为当二者不相等时，不代表dp[j]就为0，只是初始化时满足这个条件而已
				dp[j] = 0
			}
			// 遇到更大的结果就更新
			if dp[j] > res {
				res = dp[j]
			}
		}
	}
	return res
}
