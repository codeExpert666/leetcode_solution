package DailyCode

// 滑动窗口（每次统计以左边界为首的子串）
func CountKConstraintSubstrings(s string, k int) int {
	var res int
	var count0, count1 int // 统计窗口中0与1的个数
	var left, right int    // 窗口左右边界
	for ; right < len(s); right++ {
		// 更新0与1个数
		if s[right] == '0' {
			count0++
		} else {
			count1++
		}
		// 窗口右边界的移动会破坏k约束，此时移动左边界，并同时记录符合条件的子串个数
		for count0 > k && count1 > k {
			res += right - left // 统计以s[left]为首，s[right-1]为尾的字符串个数
			if s[left] == '0' { // 0移出窗口
				count0--
			} else { // 1移出窗口
				count1--
			}
			left++ // 移动左边界
		}
	}
	// 右边界移出窗口后，左边界若仍在字符串中，更新以左边界为首的子串
	for ; left < len(s); left++ { // 此时right=len(s)
		res += right - left
	}
	return res
}

// 方法二：滑动窗口（每次统计以右边界为尾的子串），在这个角度下，for循环后不需要再单独写一段针对left的for循环
