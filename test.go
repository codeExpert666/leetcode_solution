package main

import (
	"fmt"
)

// 利用hash表记录一个数组的元素出现情况，进而用于判断另一数组元素在前一数组中的出现情况
func intersection(nums1 []int, nums2 []int) []int {
	// hash表记录长度较小的数组中各数字出现的次数，以节省空间
	if len(nums1) <= len(nums2) {
		return restrictIntersection(nums1, nums2)
	}
	return restrictIntersection(nums2, nums1)
}

func restrictIntersection(minNums []int, maxNums []int) []int {
	res := make([]int, 0, len(minNums))
	// 记录较小元素出现情况
	m := make(map[int]int)
	for _, n := range minNums {
		m[n]++
	}
	// 判断较大数组的元素出现情况
	for _, n := range maxNums {
		if count, ok := m[n]; ok && count > 0 { // 存在且还有剩余出现次数
			res = append(res, n)
			m[n]--
		}
	}
	return res
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
