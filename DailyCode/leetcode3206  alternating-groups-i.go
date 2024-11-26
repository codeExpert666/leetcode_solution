package DailyCode

func numberOfAlternatingGroups(colors []int) int {
	var res int
	n := len(colors)
	for i := 0; i < n; i++ {
		if colors[i] == colors[(i+2)%n] && colors[i] != colors[(i+1)%n] {
			res++
		}
	}
	return res
}
