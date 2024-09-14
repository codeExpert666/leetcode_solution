package Array

// 快慢双指针
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			//tmp := nums[slow]
			//nums[slow] = nums[fast]
			//nums[fast] = tmp
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
