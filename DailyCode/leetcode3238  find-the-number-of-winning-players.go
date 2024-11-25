package DailyCode

func winningPlayerCount(n int, pick [][]int) int {
	var res int
	winners := make([]bool, n)  // 标记获胜玩家，防止重复统计
	count := make([][11]int, n) // 最多11种球
	for _, p := range pick {
		player, ball := p[0], p[1]
		if !winners[player] { // 已获胜玩家不再统计
			count[player][ball]++
			if count[player][ball] > player {
				res++
				winners[player] = true
			}
		}
	}
	return res
}
