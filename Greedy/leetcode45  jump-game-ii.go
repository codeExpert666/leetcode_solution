package Greedy

/**
 * 贪心法：一次遍历，不关注每次跳几步，只关注当前可以跳跃的最大范围
 * 在遍历中不断更新该最大范围，关注第一次覆盖终点时的跳跃次数
 * 本题的难点在于记录跳跃次数，实际上这与每层结点隔开的层序遍历相似
 * 需要记录两个变量：1、当前最大覆盖范围 2、下一步的最大覆盖范围
 * 当遍历到当前最大覆盖时更新跳跃次数
 */
func jump(nums []int) int {
	if len(nums) == 1 {
		// 长度为1的数组特殊处理
		return 0
	}
	var cover, jumpCount int // 下一步最大覆盖范围与跳跃次数
	for i := 0; i < len(nums); {
		coverCopy := cover   // 当前最大覆盖范围
		jumpCount++          // 先更新跳跃次数
		for i <= coverCopy { // 在当前最大覆盖范围内遍历
			if nums[i]+i > cover {
				// 更新下一步最大覆盖范围
				cover = nums[i] + i
			}
			if cover >= len(nums)-1 {
				// 当下一步最大覆盖范围覆盖顶点时，直接返回，跳跃次数已提前更新
				return jumpCount
			}
			i++
		}
	}
	return jumpCount // 这一步执行不到
}

/**
 * 贪心优化版本：无需特殊处理长度为1的数组
 */
func jump1(nums []int) int {
	n := len(nums)
	cur, next := 0, 0
	step := 0
	for i := 0; i < n; i++ {
		next = max(nums[i]+i, next)
		if i == cur {
			if cur != n-1 {
				step++
				cur = next
				if cur >= n-1 {
					return step
				}
			} else {
				return step
			}
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * 贪心再优化版本：优点是代码更加简洁，缺点是必须完整遍历一遍数组，无法提前跳出
 */
func jump2(nums []int) int {
	n := len(nums)
	cur, next := 0, 0
	step := 0
	for i := 0; i < n-1; i++ {
		next = max(nums[i]+i, next)
		if i == cur {
			cur = next
			step++
		}
	}
	return step
}
