package Array

import (
	"container/heap"
	"sort"
)

type minHeap struct {
	sort.IntSlice
	nums []int
}

func (hp *minHeap) Less(i, j int) bool {
	idxI, idxJ := hp.IntSlice[i], hp.IntSlice[j]
	if hp.nums[idxI] == hp.nums[idxJ] {
		return idxI < idxJ
	}
	return hp.nums[idxI] < hp.nums[idxJ]
}

func (hp *minHeap) Push(any) {}

func (hp *minHeap) Pop() (_ any) { return }

func GetFinalState(nums []int, k int, multiplier int) []int {
	l := len(nums)
	idx := make([]int, l)
	for i := 0; i < l; i++ {
		idx[i] = i
	}
	numsHeap := &minHeap{idx, nums}
	heap.Init(numsHeap)
	for range k {
		nums[idx[0]] *= multiplier
		heap.Fix(numsHeap, 0)
	}
	return nums
}
