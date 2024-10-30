package MonotonicStack

// 单调栈：适用于解决寻找数组中每个元素的下一个更大元素或者更小元素。
// 找更大元素：单调栈从栈头到栈底递增
// 找更小元素：单调栈从栈头到栈底递减
func DailyTemperatures(temperatures []int) []int {
	var res = make([]int, len(temperatures))
	var stack []int // 存储下标，方便记录距离
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			// 记录栈顶元素的下一更大元素距离
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			// 栈顶出栈
			stack = stack[:len(stack)-1]
		}
		// 入栈
		stack = append(stack, i)
	}
	return res
}
