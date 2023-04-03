package main

import (
	"fmt"
)

var (
	res9  [][]int
	path9 []int
)

func findSubsequences(nums []int) [][]int {
	res9, path9 = make([][]int, 0), make([]int, 0, len(nums))
	backtracking8(nums, 0, -150)
	return res9
}

func backtracking8(nums []int, startIndex int, preValue int) {
	// 终止条件（存在一种情况，原数组连续递增，那么终止条件就应是遍历到数组尾部）
	if startIndex == len(nums) || nums[startIndex] < preValue {
		if len(path9) >= 2 { // 注意题目条件，子序列长度至少为2
			tmp := make([]int, len(path9))
			copy(tmp, path9)
			res9 = append(res9, tmp)
		}
		return
	}

	for i := startIndex; i < len(nums); i++ {
		path9 = append(path9, nums[i])
		backtracking8(nums, i+1, nums[i])
		path9 = path9[:len(path9)-1]
	}
}

func main() {
	fmt.Println(findSubsequences([]int{4, 4, 3}))
}
