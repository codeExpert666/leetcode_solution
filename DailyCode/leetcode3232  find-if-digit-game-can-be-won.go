package DailyCode

func canAliceWin(nums []int) bool {
	single, double := 0, 0
	for _, v := range nums {
		if v <= 9 {
			single += v
		} else {
			double += v
		}
	}
	return single != double
}
