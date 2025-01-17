package Array

import "container/heap"

// MinOperations1 每次取出最小的两个元素，且元素可能重复，故采用小根堆的方式实现
func MinOperations1(nums []int, k int) int {
	// 将 nums 转换为小根堆
	// mh := &mHeap(nums) 注意这种写法是错误的，因为go编译器会将其解释为试图获取一个临时值的地址
	// 这在go中是不允许的。
	mh := (*mHeap)(&nums)
	heap.Init(mh)

	// 计算操作次数
	var res int
	for nums[0] < k { // 最小的元素小于 k 时继续操作
		// 取出最小的两个元素
		v1 := heap.Pop(mh).(int)
		v2 := heap.Pop(mh).(int)
		// 将新值加入小根堆
		heap.Push(mh, 2*min(v1, v2)+max(v1, v2))
		res++
	}

	return res
}

type mHeap []int

func (mh *mHeap) Len() int {
	return len(*mh)
}

func (mh *mHeap) Less(i, j int) bool {
	return (*mh)[i] < (*mh)[j]
}

func (mh *mHeap) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *mHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *mHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}
