package binarysearch

import "sort"

// 若一个数 x 在数组 arr 中出现次数超过四分之一，且该数组有序。
// 记 m = x / 4 则 arr[m], arr[2m+1], arr[3m+2] 中必包含该数
// 具体解释可参考灵神题解：https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array/solutions/3067559/olog-n-er-fen-cha-zhao-zheng-que-xing-zh-5mu9/
func FindSpecialInteger(arr []int) int {
	m := len(arr) / 4

	// 检验 arr[m], arr[2m+1], arr[3m+2] 中的哪个数的出现次数超过了四分之一
	for _, i := range []int{m, 2*m + 1} {
		// 二分查找第一个大于等于 arr[i] 的元素下标
		// 也即 arr[i] 第一次出现的下标
		j := sort.SearchInts(arr, arr[i])
		if arr[j+m] == arr[j] {
			return arr[i]
		}
	}

	return arr[3*m+2]
}
