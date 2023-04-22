package Greedy

/**
 * 贪心，每收到一位顾客的钱做两个步骤：检查能否给等待列表中的顾客找零，能否给自身找零
 * lc提交失败，原因竟然是我考虑复杂了，必须按顺序结账，不存在什么等待列表，某位顾客当前
 * 找不开，直接就失败，不能等后面的顾客支付零钱后再找零。相当于做了一道更难的题。
 */
func lemonadeChange(bills []int) bool {
	money := make(map[int]int)    // 已经收到的可以用于找零的钞票面值与数量(5与10)
	waitList := make(map[int]int) // 当前未能找零，仍需等待方能完成找零的情况（10与20）
	for _, v := range bills {
		if v == 20 { // 需要找零15元，且该面值不能用于找零
			// 15元可由3个5元或一个10元与一个5元组成
			// 有10元尽量先找10元
			if money[10] > 0 && money[5] > 0 { // 存在10元零钱与5元零钱
				money[10]--
				money[5]--
			} else if money[5] >= 3 { // 存在3张5元零钱
				money[5] -= 3
			} else {
				waitList[v]++
			}
		} else { // 这些面值可以用于找零
			money[v]++
			if v == 10 {
				// 先处理等待列表，且只能处理20元找零的情况
				if waitList[20] > 0 && money[5] > 0 {
					// 有20元待找零，且有5元零钱
					waitList[20]--
					money[5]--
					money[v]--
				}
				// 处理自身的找零
				if money[5] > 0 { // 可以找零
					money[5]--
				} else { // 无法找零
					waitList[v]++
				}
			} else { // 5元无需找零
				// 处理等待列表
				if waitList[10] > 0 { // 先处理10元待找零情况，只需一个5元
					waitList[10]--
					money[v]--
				} else if waitList[20] > 0 { // 处理20元找零情况
					if money[10] > 0 { // 若有10元零钱，可以找零
						waitList[20]--
						money[10]--
						money[v]--
					} else if money[v] >= 3 { // 有3个5元，可以找零
						waitList[20]--
						money[v] -= 3
					}
				}
			}
		}
	}
	// 若能全部找零，则等待列表必为空
	return waitList[10] == 0 && waitList[20] == 0
}

/**
 * 贪心，是上面解法的简化版本，故不做注释说明
 */
func lemonadeChange1(bills []int) bool {
	money := make(map[int]int) // 已经收到的可以用于找零的钞票面值与数量(5与10)
	for _, v := range bills {
		if v == 5 {
			money[v]++
		} else if v == 10 {
			if money[5] == 0 {
				return false
			}
			money[5]--
			money[v]++
		} else {
			if money[10] > 0 && money[5] > 0 {
				money[10]--
				money[5]--
			} else if money[5] >= 3 {
				money[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
