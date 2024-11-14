package DailyCode

// 常规做法：时间复杂度为O(n*Q)，很高
func CountKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	res := make([]int64, len(queries))
	for idx, q := range queries {
		l := q[0]
		r := q[1]
		// 计算满足条件的子字符串数量
		res[idx] = countValidSubstrings(s, l, r, k)
	}
	return res
}

// 计算满足 k 约束的子字符串数量
func countValidSubstrings(s string, l, r, k int) int64 {
	var count int64 = 0
	zerosCount := 0 // 滑动窗口内 0 的数量
	onesCount := 0  // 滑动窗口内 1 的数量
	left := l       // 滑动窗口的左边界

	for right := l; right <= r; right++ {
		// 更新窗口内 0 和 1 的数量
		if s[right] == '0' {
			zerosCount++
		} else {
			onesCount++
		}

		// 当窗口内的 0 和 1 数量都超过 k 时，移动左边界
		for zerosCount > k && onesCount > k && left <= right {
			if s[left] == '0' {
				zerosCount--
			} else {
				onesCount--
			}
			left++
		}
		// 累加满足条件的子字符串数量
		count += int64(right - left + 1)
	}
	return count
}

// 方法二：可以通过预先记录一些数据，实现快速查询
// 所需知识点都很基础，但组合起来却相当复杂，见灵神题解：https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-ii/solutions/2884463/hua-dong-chuang-kou-qian-zhui-he-er-fen-jzo25
func CountKConstraintSubstrings1(s string, k int, queries [][]int) []int64 {
	n := len(s)
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l
		sum[i+1] = sum[i] + i - l + 1
	}

	right := make([]int, n)
	l = 0
	for i := range right {
		for l < n && left[l] < i {
			l++
		}
		right[i] = l
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := min(right[l], r+1)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}
