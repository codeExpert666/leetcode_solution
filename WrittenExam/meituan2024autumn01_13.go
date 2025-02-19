package main

import (
	"fmt"
)

var memo [][]int
var colors []int
var n int

func main() {
	var q int
	_, _ = fmt.Scan(&n, &q)

	colors = make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&colors[i])
	}

	initMemo()

	left, right := 0, n-1 // 彩带两端的相对位置（在长度为 n 的彩带上的位置）
	for i := 0; i < q; i++ {
		var direction byte
		var cutL int
		_, _ = fmt.Scanf("%c %d", &direction, &cutL)
		if cutL >= n { // 裁剪长度大于单条彩带长度，则颜色数量等于单条彩条
			fmt.Println(memo[0][n-1])
		} else if direction == 'L' { // 裁剪左侧
			fmt.Println(cutLeft(left, cutL))
			left = (left + cutL) % n
		} else {
			fmt.Println(cutRight(right, cutL))
			right = (right - cutL + n) % n
		}
	}
}

func initMemo() {
	// memo[i][j]： 起点为 i，终点为 (j+n)%n 的彩带共有多少种颜色
	memo = make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}
	// 初始化起点为 0 的部分
	colorSet := make(map[int]struct{})
	for i, c := range colors {
		colorSet[c] = struct{}{}
		memo[0][i] = len(colorSet)
	}
}

func cutLeft(left, length int) int {
	p := &memo[left][(left+length-1)%n]
	if *p == 0 {
		colorSet := make(map[int]struct{})
		for i := 0; i < length; i++ {
			index := (left + i) % n
			color := colors[index]
			colorSet[color] = struct{}{}
			memo[left][index] = len(colorSet)
		}
	}

	return *p
}

func cutRight(right, length int) int {
	left := (right - length + 1 + n) % n
	return cutLeft(left, length)
}
