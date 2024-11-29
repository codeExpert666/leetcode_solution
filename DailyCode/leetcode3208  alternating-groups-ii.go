package DailyCode

// NumberOfAlternatingGroups1 两层循环，复杂度有些高
func NumberOfAlternatingGroups1(colors []int, k int) int {
	var res int
	l := len(colors)
Outer:
	for i := 0; i < l; i++ {
		for j := 0; j < k-1; j++ {
			cur, next := (i+j)%l, (i+j+1)%l
			if colors[cur] == colors[next] { // 颜色不交替
				// 剪枝，跳过所有包含colors[cur]与colors[next]的子串
				if cur < i {
					break Outer // 利用 Go 特性，跳出外层循环
				}
				i = cur
				continue Outer // 利用 Go 特性，外层进入下一循环
			}
		}
		res++
	}
	return res
}

// NumberOfAlternatingGroups2 滑动窗口，一次遍历
func NumberOfAlternatingGroups2(colors []int, k int) int {
	var res int
	l := len(colors)
	for left, right := 0, 1; left < l; right = (right + 1) % l {
		if colors[right] == colors[(right-1+l)%l] { // 颜色不交替
			// 剔除包含colors[right-1]和colors[right]的子串
			if right < left {
				break
			}
			left = right
		} else if (right+l-left)%l+1 == k { // 找到一个交替组（注意子串长度的计算方式，加1放在括号外面）
			res++
			left++ // 交替组元素个数上限
		}
	}
	return res
}

// 优化思路：不需要 left 记录左边界，只需要用一个 int 变量记录当前连续交替瓷砖数量即可
// 连续两个元素不同，变量加 1；相同则重置。在这种思路下，遍历的数组元素一定是交替组的右边界，
// 为此，需要从第-(k-2)个元素开始遍历。
func numberOfAlternatingGroups3(colors []int, k int) int {
	n := len(colors)
	res, cnt := 0, 1
	for i := -k + 2; i < n; i++ {
		if colors[(i+n)%n] != colors[(i-1+n)%n] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			res++
		}
	}
	return res
}
