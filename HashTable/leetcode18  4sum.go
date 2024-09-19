package HashTable

import "sort"

// 同三数之和，哈希很麻烦，利用双指针
// 里面细节很多，很容易错
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	// 先对数组排序
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ { // i < len(nums)-3边界关键，以防在i超出正常边界时以下面的return结束
		//if nums[i] > target {  // 这一步不对，因为target可能小于0，例如：target=-5， 四个数分别是-4，-1，0，0
		//	return res  // 剪枝错误
		//}
		// 正确剪枝
		if nums[i] > target && (target >= 0 || nums[i] >= 0) {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ { // j < len(nums)-2边界关键，以防在i超出正常边界时以下面的return结束
			//if nums[i]+nums[j] > target {  // 更不对，例如：target=-6， 四个数分别为-4， -1， -1， 0
			//	return res  // 剪枝错误
			//}
			// 正确剪枝
			if nums[i]+nums[j] > target && (target >= 0 || nums[i]+nums[j] >= 0) {
				return res
			}
			if j > i+1 && nums[j] == nums[j-1] { // 注意j > i+1这个边界，不是j > 1
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					right--
					left++
				}
			}
		}
	}
	return res
}
