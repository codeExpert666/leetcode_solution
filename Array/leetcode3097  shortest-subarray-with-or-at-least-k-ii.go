package Array

import "math"

// MinimumSubarrayLength 滑动窗口+位运算
func MinimumSubarrayLength(nums []int, k int) int {
	bitCount := make([]int, 30)
	orRes := 0
	ans := math.MaxInt

	for left, right := 0, 0; right < len(nums); right++ {
		orRes |= nums[right]
		for i := 0; i < 30; i++ {
			bitCount[i] += (nums[right] >> i) & 1
		}
		for orRes >= k && left <= right {
			ans = min(ans, right-left+1)
			for i := 0; i < 30; i++ {
				bitCount[i] -= (nums[left] >> i) & 1
			}
			orRes = getOrRes(bitCount)
			left++
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func getOrRes(bitCount []int) int {
	var res int
	for i := 0; i < len(bitCount); i++ {
		if bitCount[i] > 0 {
			res += 1 << i
		}
	}
	return res
}
