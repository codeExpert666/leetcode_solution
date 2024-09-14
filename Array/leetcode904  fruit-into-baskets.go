package Array

// 滑动窗口（左闭右闭）
func totalFruit(fruits []int) int {
	left, right := 0, 0          // 窗口左右边界
	content := make([]int, 0, 2) // 窗口中包含的不重复数字（保持其在原数组中的顺序）
	content = append(content, fruits[0])
	nextLeft := 0 // 左边界下一次的更新值
	maxLen := 0   // 记录最大窗口大小
	for right = 1; right < len(fruits); right++ {
		// 只有当值发生改变，窗口左边界才有可能更新
		if fruits[right] != fruits[right-1] {
			if len(content) == 1 { // 第一次找到满足条件的窗口
				content = append(content, fruits[right])
			} else {
				if fruits[right] != content[0] {
					// 发现新元素，需更新最大窗口的值，并移动左边界
					if right-left > maxLen {
						maxLen = right - left
					}
					left = nextLeft
				}
				// 无论是否发现新元素，都需更新content
				content[0] = content[1]
				content[1] = fruits[right]
			}
			// 更新下一次左边界
			nextLeft = right
		}
	}
	// 由于更新只发生在数字变化时，可能导致遍历结束最后一个窗口无法更新到最终结果中
	// 需再判断一次更新
	if right-left > maxLen {
		maxLen = right - left
	}
	return maxLen
}
