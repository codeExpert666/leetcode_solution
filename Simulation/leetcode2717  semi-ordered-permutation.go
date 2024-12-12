package Simulation

func SemiOrderedPermutation(nums []int) int {
	n := len(nums)
	// 首先找到 1 和 n 的位置
	var pos1, posN, count int // count 表示当前找到位置的元素个数
	for i, v := range nums {
		if v == 1 || v == n {
			if v == 1 {
				pos1 = i
			} else {
				posN = i
			}
			count++
			if count == 2 { // 1 和 n 都已找到
				break
			}
		}
	}
	// 直接计算所需步数
	if pos1 < posN {
		return pos1 + (n - 1 - posN)
	}
	return pos1 + (n - 1 - posN) - 1
}
