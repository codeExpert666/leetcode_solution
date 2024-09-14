package Array

// 方法一：左闭右闭区间
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}

// 方法一：左闭右开区间
func searchInsert2(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low // 返回low或high都可以，因为low与high此时指向同一位置
}
