package weeklycompetition

func subarraySum(nums []int) int {
	var res int

	for i, v := range nums {
		// 确定起始下标
		start := max(0, i-v)
		// 计算前缀和
		if i > 0 {
			nums[i] += nums[i-1]
		}
		// 累计结果
		if start > 0 {
			res += nums[i] - nums[start-1]
		} else {
			res += nums[i]
		}
	}

	return res
}

// 灵神前缀和题解
// 这里注意一点：对于数组 nums，前缀和数组 prefix 的长度一般取 len(nums)+1,
// 于是 prefix[i] = nums[0] + ... + nums[i-1], prefix[0] 无意义
// 这样的设计可以保证不用再对 nums[0] 进行特殊处理，从而确保代码的简洁统一
func subarraySum1(nums []int) (ans int) {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	for i, num := range nums {
		ans += s[i+1] - s[max(i-num, 0)]
	}
	return
}
