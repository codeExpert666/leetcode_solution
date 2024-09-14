package Array

// 快慢双指针法（一个指针负责指向新数组的末尾我，另一个指针负责寻找不是val的元素）
func removeElement(nums []int, val int) int {
	var left, right int
	// 首先遍历找到第一个val值元素
	for ; left < len(nums); left++ {
		if nums[left] == val {
			break
		}
	}
	// 从left开始往后遍历，不断将后续非val值元素替换到left位置
	for right = left + 1; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	// left始终指向新数组末尾
	return left
}

// 快慢双指针法（简洁版本，一个循环即可）
func removeElement1(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
