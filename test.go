package main

import "fmt"

var m = make(map[int]int) // 记录已经计算过的斐波那契值，避免重复计算

func fib(n int) int {
	if n <= 1 {
		m[n] = n
		return n
	}
	// 斐波那契数列前两项值
	tmp1, e1 := m[n-2]
	tmp2, e2 := m[n-1]
	if !e1 {
		tmp1 = fib(n - 2)
		m[n-2] = tmp1
	}
	if !e2 {
		tmp2 = fib(n - 1)
		m[n-1] = tmp2
	}
	// 当前项等于前两项之和
	return tmp1 + tmp2
}

func main() {
	fmt.Print(fib(2))
}
