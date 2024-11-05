package DailyCode

// 价值115只能由4枚10硬币和一枚75硬币组成
func LosingPlayer(x int, y int) string {
	if min(x, y/4)%2 == 0 { // 最后一次是Bob拿
		return "Bob"
	}
	return "Alice"
}
