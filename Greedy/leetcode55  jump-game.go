package Greedy

/**
 * 回溯法，每次跳最大步数，如果无法超过终点，就回溯
 * 这种方法本质上还是暴力解法，leetcode会超时
 */
func canJump(nums []int) bool {
	// 回溯函数
	var backtrack func(int) bool
	backtrack = func(index int) bool {
		// 终止条件
		if index >= len(nums)-1 {
			return true
		}
		// 按照当前位置所能跳跃的最大长度递减遍历
		for i := nums[index]; i >= 1; i-- {
			index += i
			if backtrack(index) {
				return true
			}
			index -= i // 回溯
		}
		// 当跳跃长度为0时，永远不可能到达
		return false
	}

	return backtrack(0)
}

/**
 * 贪心法：一次遍历，不关注每次跳几步，只关注当前可以跳跃的最大范围
 * 在遍历中不断更新该最大范围，看其能够覆盖终点
 */
func canJump1(nums []int) bool {
	var cover int

	for i := 0; i <= cover; i++ {
		tmp := nums[i] + i
		if tmp > cover {
			// 更新覆盖范围
			cover = tmp
		}
		if cover >= len(nums)-1 {
			// 当覆盖范围包括终点，则能到达终点
			return true
		}
	}
	// 若跳出循环，则一定只能到达终点前的0，也即无法到达终点
	return false
}
