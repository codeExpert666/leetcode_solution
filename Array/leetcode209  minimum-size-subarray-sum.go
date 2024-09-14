package Array

// 滑动窗口（左闭右闭）
func minSubArrayLen(target int, nums []int) int {
	sum, left, right := nums[0], 0, 0 // 记录当前窗口的边界与内部元素的和
	minLen := 0                       // 记录找到的最小窗口
	for {
		if sum < target {
			right++
			if right == len(nums) {
				break // 窗口超出数组范围
			}
			sum += nums[right] // 右边界向右扩充窗口
		} else {
			if minLen == 0 || right-left+1 < minLen {
				minLen = right - left + 1 // 更新最小窗口
			}
			// 右边界继续向右扩充没有意义，因为要找满足条件的最小窗口
			// 左边界向右缩减窗口以探索新的可能的最小窗口
			sum -= nums[left]
			left++
		}
	}
	return minLen
}
