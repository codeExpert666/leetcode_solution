package StackAndQueue

// 单调队列：单调递减或单调递增的队列
// 本题仅考虑窗口中的最大值是不够的，需要对窗口中所有元素进行完整排序，然后再考虑窗口的移动更新
// 考虑这样一种情况：窗口的第一个元素是窗口的最大值，后续移动时，最大值丢失，新进入窗口的元素需要与
// 先前窗口的第二大值进行排序，再往后，还可能与第三大值进行排序。
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	// 双端队列，存储窗口中元素的索引，索引在窗口中对应的元素值单调递减
	// 注意：队列并不是窗口中所有元素的排序，只是最大值到窗口末尾值的排序
	deque := make([]int, 0, len(nums))
	for i, v := range nums {
		// 窗口滑动需移除最左边元素，但队列不一定移除队头元素
		// 主要取决于窗口移除的元素是否为窗口最大值
		if len(deque) > 0 && deque[0] == i-k {
			deque = deque[1:] // 队头出队
		}
		// 窗口滑动需要增加最右边元素，在队列中的处理方式是
		// 移除队尾所有小于该元素的索引（这些元素的保留无意义），然后增添该元素索引
		for len(deque) > 0 && nums[deque[len(deque)-1]] < v {
			deque = deque[:len(deque)-1] // 队尾出队
		}
		deque = append(deque, i) // 队尾新元素入队
		// 记录窗口最大值
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
