package main

import (
	"fmt"
)

var (
	res7  [][]int
	path7 []int
)

func subsets(nums []int) [][]int {
	res7, path7 = make([][]int, 0), make([]int, 0, len(nums))
	backtracking6(nums, 0)
	return res7
}

func backtracking6(nums []int, index int) {
	// 终止条件
	if index == len(nums) {
		tmp := make([]int, len(path7))
		copy(tmp, path7)
		res7 = append(res7, tmp)
		return
	}
	for i := 0; i < 2; i++ { // 每个位置的元素都有两种可能，在或不在子集中
		if i == 0 {
			// 在子集中
			path7 = append(path7, nums[index])
			backtracking6(nums, index+1)
			path7 = path7[:len(path7)-1]
		} else {
			// 不在子集中
			backtracking6(nums, index+1)
		}
	}
}

func main() {
	fmt.Println(subsets([]int{1}))
}
