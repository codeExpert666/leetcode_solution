package DynamicProgramming

/**
 * 在leetcode198的基础上考虑，就比较简单（分成两种情况，分别利用lc198的函数，取较大者）
 */
func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 两种情况：1、不打劫最后一个房屋
	tmp1 := rob(nums[:len(nums)-1])
	// 2、不打劫第一间房屋
	tmp2 := rob(nums[1:])

	// 返回二者中的较大者
	return max(tmp1, tmp2)
}
