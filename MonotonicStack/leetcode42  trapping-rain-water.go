package MonotonicStack

// 雨水总量有两种计算方式：按行计算与按列计算
// 单调栈适用于按行计算；双指针法适用于按列计算
// 单调栈从栈头到栈底单调递增，从而保证形成凹槽容纳雨水
func Trap(height []int) int {
	var res int
	var stack []int
	for i, h := range height {
		// 注意这里for中对待插入元素与栈顶元素相同的情况也进行了处理
		// 因为在求凹槽宽度的时候，如果遇到相同高度的柱子，需要使用最右边的柱子来计算宽度
		// 所以栈中对于相同高度的柱子需要始终保持最新的位置
		for len(stack) > 0 && height[stack[len(stack)-1]] <= h {
			midHeight := height[stack[len(stack)-1]] // 凹槽底部高度
			stack = stack[:len(stack)-1]
			if len(stack) > 0 { // 如果此时栈中没有元素了，则代表此凹槽没有左壁，无需计算雨水
				res += (min(height[stack[len(stack)-1]], h) - midHeight) *
					(i - stack[len(stack)-1] - 1) // 凹槽中雨水体积（按行计算）
			}
		}
		stack = append(stack, i)
	}
	return res
}
