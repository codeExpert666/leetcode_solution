package Array

import "sort"

// NextPermutation 获取一个整数数组的下一个字典序更大的排列可采用如下两个步骤：
// 1、找到最后一个严格升序对，交换升序对中两个元素值
// 2、值交换完成后，将较大元素位置后的所有元素升序排列，即为所求
func NextPermutation(nums []int) {
	// 步骤1
	i := len(nums) - 2 // 记录步骤2排序位置
Outer:
	for ; i >= 0; i-- { // 倒序遍历，找到第一不满足升序（从后往前）的元素
		if nums[i] < nums[i+1] { // nums[i]需要与其后的最小的更大元素交换
			for j := len(nums) - 1; j > i; j-- { // 倒序遍历
				if nums[j] > nums[i] { // 交换
					nums[i], nums[j] = nums[j], nums[i]
					break Outer // 采用golang标签语法跳出外层循环
				}
			}
		}
	}
	// 步骤2
	sort.Ints(nums[i+1:])
}
