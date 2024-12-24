package Array

// 倒序遍历，一旦发现重复元素，就将其前面所有元素删除
// 同时考虑到删除元素必须是 3 的整数倍，后续可能也要删除个别元素
func MinimumOperations(nums []int) int {
	frequency := [100]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		frequency[nums[i]]++
		if frequency[nums[i]] > 1 { // 出现重复
			group := (i + 1) / 3 // 删除以 3 个为一组
			if (i+1)%3 != 0 {    // 不满 3 个按照 3 个算
				group++
			}
			return group
		}
	}
	return 0
}
