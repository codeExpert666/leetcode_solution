package Array

// 解法一：利用二分法（左闭右闭）找到某一个target的位置，然后向左右扩展以寻找边界
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			// 从mid向左扩展，找到结果区间的左端点
			for i := 0; mid-i >= 0 && nums[mid-i] == target; i++ {
				res[0] = mid - i
			}
			// 从mid向右扩展，找到结果区间的右端点
			for i := 0; mid+i < len(nums) && nums[mid+i] == target; i++ {
				res[1] = mid + i
			}
			return res
		}
	}
	// 程序运行到这里说明target未找到，将返回默认值
	return res
}

// 解法二：利用二分法分别获取左右边界
func searchRange2(nums []int, target int) []int {
	leftBoard := findLeft(nums, target)
	rightBoard := findRight(nums, target)

	if leftBoard == -2 || rightBoard == -2 { // target不在nums范围内
		return []int{-1, -1}
	}
	if rightBoard-leftBoard >= 2 { // nums包含target
		return []int{leftBoard + 1, rightBoard - 1}
	}
	return []int{-1, -1} // target在nums范围内，但nums不包含target
}

// 寻找左边界（左闭右闭）
func findLeft(nums []int, target int) int {
	res := -2
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
			res = high
		}
	}
	return res
}

// 寻找右边界（左闭右闭）
func findRight(nums []int, target int) int {
	res := -2
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
			res = low
		}
	}
	return res
}
