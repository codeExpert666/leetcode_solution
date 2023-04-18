package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	var sum int // 记录总体油量得失情况
	for i := 0; i < len(gas); i++ {
		gas[i] = gas[i] - cost[i] // 将每一站的油量得失存在gas中
		sum += gas[i]
	}
	if sum < 0 {
		// 若总体油量得失为负，则无解
		return -1
	}
	// 若有解
	var res int
	for ; res < len(gas); res++ {
		if gas[res] >= 0 && isSolution(gas, res) {
			// 起始站的油量得失一定为正，利用此结论进行剪枝
			break
		}
	}
	return res
}

// 判断给定站能否成为起始站
func isSolution(gasAdded []int, start int) bool {
	// 注意加油站站都在环上，处理方式类似于循环队列
	var curGas int // 记录当前油量
	index := start // 当前所在站
	l := len(gasAdded)
	for ; index != (start-1)%l; index = (index + 1) % l {
		curGas += gasAdded[index]
		if curGas < 0 {
			// 一旦经过该站油量为负，则该站不能成为起始站
			return false
		}
	}
	// 上面循环剩余终点站未处理
	curGas += gasAdded[index]
	return curGas >= 0
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
