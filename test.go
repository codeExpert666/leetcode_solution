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

var (
	res1  [][]int
	path1 []int
	sum   int
)

func combinationSum3(k int, n int) [][]int {
	path1, res1 = make([]int, 0, k), make([][]int, 0)
	backtracking(k, n, 1)
	return res1
}

func backtracking(k int, n int, startNum int) {
	// 终止条件
	if len(path1) == k && sum == n {
		tmp := make([]int, 0, k)
		tmp = append(tmp, path1...)
		res1 = append(res1, tmp)
		return
	}
	for i := startNum; i <= 9-(k-len(path1))+1; i++ {
		// 添加
		sum += i
		path1 = append(path1, i)
		backtracking(k, n, i+1)
		// 回溯
		sum -= i
		path1 = path1[:len(path1)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 6))
}
