package DailyCode

import (
	"container/heap"
	"math"
)

var (
	next  []int
	numsC [][]int
)

// leetcode官解，贪心 + 小根堆
// https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/solutions/355881/zui-xiao-qu-jian-by-leetcode-solution/?envType=daily-question&envId=2024-11-25
func smallestRange(nums [][]int) []int {
	numsC = nums
	rangeLeft, rangeRight := 0, math.MaxInt32
	minRange := rangeRight - rangeLeft
	maxValue := math.MinInt32
	size := len(nums)
	next = make([]int, size)
	h := &IHeap{}
	heap.Init(h)

	for i := 0; i < size; i++ {
		heap.Push(h, i)
		maxValue = max(maxValue, nums[i][0])
	}

	for {
		minIndex := heap.Pop(h).(int)
		curRange := maxValue - nums[minIndex][next[minIndex]]
		if curRange < minRange {
			minRange = curRange
			rangeLeft, rangeRight = nums[minIndex][next[minIndex]], maxValue
		}
		next[minIndex]++
		if next[minIndex] == len(nums[minIndex]) {
			break
		}
		heap.Push(h, minIndex)
		maxValue = max(maxValue, nums[minIndex][next[minIndex]])
	}
	return []int{rangeLeft, rangeRight}
}

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return numsC[h[i]][next[h[i]]] < numsC[h[j]][next[h[j]]] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
