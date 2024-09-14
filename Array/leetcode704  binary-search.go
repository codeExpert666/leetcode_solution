package Array

// 写法一：左闭右闭区间，即target在[low, high]中
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1 // 双指针
	for low <= high {
		//mid := low + (high-low)/2 // 中间位置写法一
		mid := low + (high-low)>>1 // 中间位置写法二
		if nums[mid] < target {    // target在右半区
			low = mid + 1
		} else if nums[mid] > target { // target在左半区
			high = mid - 1
		} else { // 找到目标
			return mid
		}
	}
	// 能走到这一步，说明循环过程中没有找到目标，且此时low > high。
	// 表明找不到目标
	return -1
}

// 写法二：左闭右开区间，即target在[low, high)中
func search2(nums []int, target int) int {
	low, high := 0, len(nums) // 双指针
	for low < high {
		//mid := low + (high-low)/2 // 中间位置写法一
		mid := low + (high-low)>>1 // 中间位置写法二
		if nums[mid] < target {    // target在右半区
			low = mid + 1
		} else if nums[mid] > target { // target在左半区
			high = mid
		} else { // 找到目标
			return mid
		}
	}
	// 能走到这一步，说明循环过程中没有找到目标，且此时low > high。
	// 表明找不到目标
	return -1
}
