package main

import (
	"fmt"
	"sort"
)

var (
	res11  [][]int
	path11 []int
)

func permuteUnique1(nums []int) [][]int {
	sort.Ints(nums) // 先排序
	res11, path11 = make([][]int, 0), make([]int, 0, len(nums))
	bp(nums, make([]bool, len(nums)))
	return res11
}

func bp(nums []int, used []bool) {
	if len(path11) == len(nums) {
		tmp := make([]int, len(path11))
		copy(tmp, path11)
		res11 = append(res11, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] { // 注意这个去重操作
			continue
		}
		if !used[i] {
			path11 = append(path11, nums[i])
			used[i] = true
			bp(nums, used)
			used[i] = false
			path11 = path11[:len(path11)-1]
		}
	}
}

func main() {
	fmt.Println(permuteUnique1([]int{1, 1, 2}))
}
