package Array

import "sort"

// 当求解“最小值的最大值”或“最大值的最小值”时，可以使用二分查找
// 注意：这里的二分查找是在该“最小值”或“最大值”可能的取值范围内进行的
func MaxDistance(position []int, m int) int {
	sort.Ints(position)
	// sort.Search(n, f) 在 [0, n) 内二分查找出满足 f(i)==true 的最小 i
	return sort.Search(position[len(position)-1], func(f int) bool {
		prev := position[0]
		cnt := 1
		for _, curr := range position {
			if curr-prev >= f {
				// 如果当前位置与上一个位置的距离大于等于 f，则可以放置一个球
				cnt++
				prev = curr
			}
		}
		// 如果放置的球数小于 m，则说明 f 取值过大，需要减小 f
		return cnt < m
	}) - 1
}
