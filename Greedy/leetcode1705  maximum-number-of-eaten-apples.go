package Greedy

// 贪心：每天只要有苹果剩余，就吃一个，而且尽量吃最快腐烂的苹果
func EatenApples(apples []int, days []int) int {
	var res int

	// 获取苹果全部腐烂的日期
	var maxDay int
	for i, d := range days {
		if d != 0 && i+d > maxDay {
			maxDay = i + d
		}
	}
	if maxDay == 0 { // 完全没有苹果产出
		return res
	}

	// 初始化
	l := maxDay + 1
	corrupted := make([]int, l)  // 记录第 i 天会腐烂的苹果数
	nearestDay := len(corrupted) // 最快腐烂的苹果日期
	appleLeft := 0               // 苹果库存

	// 考虑每天获取和腐烂的苹果，再结合每天尽量吃一个最快腐烂苹果的原则，更新每日库存
	// 当库存为 0 时，达到能够吃到的苹果的最大数量
	for i := 0; i < l; i++ {
		// 更新库存
		if i < len(apples) {
			appleLeft += apples[i] - corrupted[i]
		} else {
			appleLeft -= corrupted[i]
			if appleLeft <= 0 { // 一点优化：此时已经超出 apples 数组，不可能再长出苹果
				break
			}
		}
		// 从 corrupted 移除已经腐烂的苹果
		if corrupted[i] > 0 { // 这种情况可以轻易推得 nearestDay == i
			nearestDay++
			for nearestDay < l && corrupted[nearestDay] == 0 { // 更新下一批快要腐烂苹果
				nearestDay++
			}
		}
		// 向 corrupted 中更新刚长出的苹果的腐烂日期
		if i < len(apples) && apples[i] > 0 {
			d := i + days[i]
			if corrupted[d] == 0 && d < nearestDay {
				nearestDay = d
			}
			corrupted[d] += apples[i]
		}
		// 当天吃掉一个苹果(没有库存就不吃)
		if appleLeft > 0 {
			res++
			appleLeft--
			// 且该吃掉的苹果是最近快要腐烂的苹果
			corrupted[nearestDay]--
			for nearestDay < l && corrupted[nearestDay] == 0 { // 更新下一批快要腐烂苹果
				nearestDay++
			}
		}
	}
	return res
}

// 此外，也可以用堆、优先队列、红黑树这样的数据结构去维护上一方法中的 corrupted 切片，也是很优秀的解法，而且代码量与可读性可能更好
