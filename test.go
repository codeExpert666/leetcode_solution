package main

import (
	"fmt"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func combine(n int, k int) [][]int {
	// 终止条件
	if k == 1 {
		var res [][]int
		for i := 1; i <= n; i++ {
			res = append(res, []int{i})
		}
		return res
	}
	if n == k {
		var tmp []int
		for i := 1; i <= k; i++ {
			tmp = append(tmp, i)
		}
		return [][]int{tmp}
	}
	// 从前n-1个数中找k个数
	res := combine(n-1, k)
	// 从前n-1个数中找k-1个数
	tmp := combine(n-1, k-1)
	// 与最后一个数拼接
	for i := 0; i < len(tmp); i++ {
		tmp[i] = append(tmp[i], n)
	}
	// 整合结果
	res = append(res, tmp...)
	return res
}

func main() {
	fmt.Println(combine(2, 1))
}
