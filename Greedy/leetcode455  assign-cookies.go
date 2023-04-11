package Greedy

import "sort"

/**
 * 贪心（方法一）：先将小饼干尽量分配给胃口小的孩子
 */
func findContentChildren(g []int, s []int) int {
	// 先对两个数组进行排序
	sort.Ints(g)
	sort.Ints(s)
	// 双指针，指向当前两个数组处理到的位置
	var gIndex, sIndex int
	// 遍历饼干尺寸数组，不断按顺序满足孩子
	for sIndex < len(s) && gIndex < len(g) {
		if s[sIndex] >= g[gIndex] {
			// 满足当前孩子最小需求，进行分配
			gIndex++
		}
		sIndex++
	}

	return gIndex
}

/**
 * 贪心（方法二）：先将大饼干尽量分给胃口大的孩子，不至于浪费
 */
func findContentChildren1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var res int
	sIndex, gIndex := len(s)-1, len(g)-1
	for sIndex >= 0 && gIndex >= 0 {
		if s[sIndex] >= g[gIndex] {
			res++
			sIndex--
		}
		gIndex-- // 说明当前饼干无法满足当前孩子，但有可能满足下一孩子
	}

	return res
}
