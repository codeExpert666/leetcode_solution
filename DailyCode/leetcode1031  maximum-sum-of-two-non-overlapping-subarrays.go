package DailyCode

/**
 * 暴力枚举法，以较长的子数组为为基准，枚举较短子数组
 */
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	var res, sum int // 记录最大和与当前和
	for i := 0; i <= len(nums)-firstLen; i++ {
		// 计算第一个子数组的和
		sum = getSum(nums, i, i+firstLen-1)
		sumCopy := sum // 备份
		// 计算第二个子数组的和
		// 处理第二个子数组在第一个子数组左侧的情况
		if i >= secondLen {
			for j := 0; j <= i-secondLen; j++ {
				if j == 0 {
					sum += getSum(nums, j, j+secondLen-1)
				} else {
					sum -= nums[j-1]
					sum += nums[j+secondLen-1]
				}
				// 更新最大和
				if sum > res {
					res = sum
				}
			}
		}
		sum = sumCopy // 复原
		// 处理第二个子数组在第一个子数组右侧的情况
		if len(nums)-i-firstLen >= secondLen {
			for j := i + firstLen; j <= len(nums)-secondLen; j++ {
				if j == i+firstLen {
					sum += getSum(nums, j, j+secondLen-1)
				} else {
					sum -= nums[j-1]
					sum += nums[j+secondLen-1]
				}
				// 更新最大和
				if sum > res {
					res = sum
				}
			}
		}
	}
	return res
}

// 求[start，end]的子数组的和
func getSum(nums []int, start, end int) int {
	var sum int
	for i := start; i <= end; i++ {
		sum += nums[i]
	}
	return sum
}

/**
 * 动态规划：待补充
 */
