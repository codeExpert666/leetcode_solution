package MonotonicStack

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	var res = make([]int, len(nums1))
	var stack []int // 存元素本身
	// toNext记录nums2中所有元素的下一元素，方便nums1的查找
	var toNext = make(map[int]int)
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			toNext[stack[len(stack)-1]] = v // 记录
			stack = stack[:len(stack)-1]    // 出栈
		}
		stack = append(stack, v) // 入栈
	}
	// 遍历nums1从toNext中获取对应结果
	for i, v := range nums1 {
		if nextValue, exist := toNext[v]; exist {
			res[i] = nextValue
		} else {
			res[i] = -1
		}
	}
	return res
}

// 区别在于map的用法，空间复杂度稍低
func NextGreaterElement1(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := 0; i < len(res); i++ {
		res[i] = -1 // 初始化
	}
	var stack []int                 // 存元素本身
	numToIndex := make(map[int]int) // nums1的元素映射到下标
	for i, v := range nums1 {
		numToIndex[v] = i
	}
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			// 记录
			if index, exist := numToIndex[stack[len(stack)-1]]; exist {
				res[index] = v
			}
			stack = stack[:len(stack)-1] // 出栈
		}
		stack = append(stack, v) // 入栈
	}
	return res
}
