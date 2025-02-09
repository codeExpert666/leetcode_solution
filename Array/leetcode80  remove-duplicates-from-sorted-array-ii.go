package Array

// RemoveDuplicates 快慢指针
func RemoveDuplicates(nums []int) int {
	count := 0                                // 当前元素的重复出现次数
	slow := 0                                 // 慢指针指向待插入位置
	for fast := 0; fast < len(nums); fast++ { // 快指针遍历数组
		if fast > 0 && nums[fast] != nums[fast-1] { // 遍历到下一元素
			count = 0
		}
		count++
		if count <= 2 { // 每种元素至多保留两个
			nums[slow] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow]
	return slow
}

// 灵神基于栈的方法，也很巧妙
func removeDuplicates1(nums []int) int {
	stackSize := 2 // 栈的大小，前两个元素默认保留
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[stackSize-2] { // 和栈顶下方的元素比较
			nums[stackSize] = nums[i] // 入栈
			stackSize++
		}
	}
	return min(stackSize, len(nums))
}
