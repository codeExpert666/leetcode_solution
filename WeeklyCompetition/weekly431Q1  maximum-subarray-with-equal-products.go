package weeklycompetition

func maxLength(nums []int) int {
	n := len(nums)
	maxLen := 0

	// 遍历所有可能的起点
	for i := 0; i < n; i++ {
		prod := nums[i]
		gcd := nums[i]
		lcm := nums[i]

		// 从起点开始扩展子数组
		for j := i; j < n; j++ {
			if j > i {
				prod *= nums[j]
				gcd = getGCD(gcd, nums[j])
				lcm = getLCM(lcm, nums[j])
			}

			// 检查当前子数组是否满足条件
			if prod == lcm*gcd {
				maxLen = max(maxLen, j-i+1)
			}
		}
	}

	return maxLen
}

// 计算最大公因数
func getGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 计算最小公倍数
func getLCM(a, b int) int {
	return a / getGCD(a, b) * b
}

// 返回较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
