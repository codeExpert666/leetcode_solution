package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	res := 1
	var flag int // 用于判断差值是否交替，记录的是前一组有效数据的差值
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if d == 0 {
			// 遇到差值为0，直接跳过
			continue
		}
		tmp := d * flag
		if tmp <= 0 {
			// 两种情况：1、差值交替 2、第一次遇到差值不为零的一组数据
			res++
			flag = d
		}
	}
	return res
}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
}
