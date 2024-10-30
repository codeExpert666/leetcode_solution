package MonotonicStack

// 个人对单调栈的一些理解：
// 1、单调栈的主角是栈顶元素，虽然说待插入元素的插入过程大费周章，但本质是在更新栈顶元素的相关性质
// 2、对于栈中的每个元素，其左侧元素一定是其左侧第一小（大）的元素
// 3、当待插入元素使得栈结构更新时，待插入元素便成为了待删除元素的右侧第一小（大）元素
// 4、所以对于待删除元素，其两侧第一小（大）元素均为已知

// 本题还有更优解法：在heights首尾各加一个0元素，以避免算法运行过程中需要对以下几种情况的特殊处理
// 1、栈顶元素右侧没有更小元素导致无法从栈中删除
// 2、栈顶元素左侧没有没有更小元素导致面积计算公式发生改变
// 3、待插入元素与栈顶元素相等导致面积比较时需特殊处理（首尾加0后可保证两个相等元素一定都能从栈中弹出以计算各自面积）
func LargestRectangleArea(heights []int) int {
	res := -1
	stack := make([]int, 0)
	for i, h := range heights {
		// 注意这里包含等于会多处一些无意义的max比较，但如果不加，后续出for循环后需要进行复杂的判断（栈中连续两数相等）
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			// 对于栈顶的部分元素，找到了两侧的较小元素，则栈顶柱子所能围成的最大面积随之确定
			hei := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				res = max(res, hei*(i-stack[len(stack)-1]-1))
			} else {
				res = max(res, hei*i)
			}
		}
		stack = append(stack, i)
	}
	// 仍需处理栈中剩余元素(每个元素右侧没有更小元素)
	for len(stack) > 1 {
		hei := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		res = max(res, hei*(len(heights)-1-stack[len(stack)-1]))
	}
	// 栈中最后一个元素特殊处理
	if len(stack) > 0 {
		res = max(res, heights[stack[len(stack)-1]]*len(heights))
	}
	return res
}
