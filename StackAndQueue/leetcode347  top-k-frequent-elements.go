package StackAndQueue

import "container/heap"

// Item 表示优先队列中的一个元素
type Item struct {
	value    int
	priority int
}

// PriorityQueue 实现了heap.Interface接口
type PriorityQueue []*Item

// Len 返回队列长度
func (pq PriorityQueue) Len() int { return len(pq) }

// Less 比较优先级
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

// Swap 交换元素
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 添加新元素
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop 移除并返回最高优先级的元素
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // 避免内存泄漏
	*pq = old[0 : n-1]
	return item
}

// // 更新修改一个Item的优先级
// func (pq *PriorityQueue) update(item *Item, value string, priority int) {
//     item.value = value
//     item.priority = priority
//     heap.Fix(pq, item.index)
// }

// 先使用hash表获取每个元素的出现频数
// 再利用优先队列找到出现频率前k高的元素
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	// 定义hash表记录nums中每个数的出现频数
	numCount := make(map[int]int)
	for _, n := range nums {
		numCount[n]++
	}
	// 遍历hash表，加入到优先队列中
	// key作为元素，value作为优先级
	pq := make(PriorityQueue, 0, len(nums))
	heap.Init(&pq)
	for num, count := range numCount {
		heap.Push(&pq, &Item{value: num, priority: count})
	}
	// 从优先队列中取出前k个数
	for k > 0 {
		item := heap.Pop(&pq).(*Item)
		res = append(res, item.value)
		k--
	}
	return res
}

// 还有一种更优的方法，即利用基于小根堆的优先队列，且维持堆的大小为k
// 遍历将数字映射到频数的hashmap，将数据不断添加到小根堆中，当堆的大小大于k时，
// 弹出堆顶元素，这样，堆中最终留下的便是频数前k大的元素
// 该方法的实践复杂度为nlogk，空间复杂度为大小为k的堆，均小于上一种方法

// 总结：前K大，小根堆；前K小，大根堆
