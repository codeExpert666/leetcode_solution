package Array

import "sort"

// 方法一：逐位判断以计算数字的二进制表示中1的个数
// 复杂度为数字的二进制长度
func countOnes(num int) int {
	var res int
	for i := 1; i <= num; i = i << 1 {
		if num&i == i {
			res++
		}
	}
	return res
}

// 方法二：逐个消除数字的二进制表示中的最低位1
// 复杂度为数字的二进制表示中1的个数
func countOnes1(num int) int {
	var res int
	for num > 0 {
		res++
		num &= num - 1 // 消除最低位的1
	}
	return res
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		iOnes, jOnes := countOnes(arr[i]), countOnes(arr[j])
		if iOnes == jOnes {
			return arr[i] < arr[j]
		}
		return iOnes < jOnes
	})
	return arr
}
