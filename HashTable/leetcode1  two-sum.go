package HashTable

// 本题有一个关键点（加法交换律）：若nums[i] + nums[j] = target，且i < j
// 可以考虑先遍历到j位置，因为此时nums[i]一定在hash表中
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // 记录每个元素出现的位置
	for i := 0; i < len(nums); i++ {
		if index, ok := m[target-nums[i]]; ok {
			// 另一个数在hash表中存在，则找到两个加数
			return []int{index, i}
		} else {
			// 不存在则记录
			m[nums[i]] = i
		}
	}
	return []int{-1, -1} // 找不到和为target的两个加数
}
