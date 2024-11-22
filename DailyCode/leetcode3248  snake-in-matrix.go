package DailyCode

func finalPositionOfSnake(n int, commands []string) int {
	// 统计水平和垂直方向的位移，有正负
	var horizontal, vertical int
	for _, c := range commands {
		switch c {
		case "LEFT":
			horizontal--
		case "RIGHT":
			horizontal++
		case "UP":
			vertical--
		case "DOWN":
			vertical++
		}
	}
	// 将坐标转换为位置标签
	return vertical*n + horizontal
}
