package Greedy

// PredictPartyVictory
// 每个参议员都尽可能让后续不同阵营的参议员丧失权利，这样可以让自己阵营的人数最大化
// 对于靠后的参议员，其后续没有另一阵营参议员，那就让前面的参议员丧失权利，
// 在算法中，体现为：虽然靠前的另一阵营参议员加入了队列，但其是失权的，遍历到他时会直接跳过
func PredictPartyVictory(senate string) string {
	senateSlice := []byte(senate)
	inValidR, inValidD := 0, 0 // 当前累计丧失权利参议员数量，放在外层是考虑到靠后参议员的情况
	for len(senateSlice) > 0 {
		nextRound := make([]byte, 0) // 用于存储下一轮可能拥有权利的参议员
		countR, countD := 0, 0       // 统计nextRound中R与D的数量
		for _, c := range senateSlice {
			if c == 'R' {
				if inValidR > 0 { // R不具备权利
					inValidR--
				} else { // R具备权利
					nextRound = append(nextRound, c)
					countR++
					inValidD++ // 让一个D失效
				}
			} else {
				if inValidD > 0 { // D不具备权利
					inValidD--
				} else { // D具备权利
					nextRound = append(nextRound, c)
					countD++
					inValidR++ // 让一个R失效
				}
			}
		}
		if countD == 0 { // 下一轮参议员中全是R
			return "Radiant" // R获胜
		}
		if countR == 0 { // 下一轮参议员中全是D
			return "Dire" // D获胜
		}
		senateSlice = nextRound // 更新拥有权利的参议员
	}
	return "" // 不可能走到这一步
}

// 代码优化思路：
// 1、不用新建nextRound切片，直接在给定senate切片上进行操作，失权的参议员赋一个别的值即可
// 2、inValidR与inValidD可以合并成一个变量，用正负区分是谁无效
// 3、无需用countR和countD统计数量，可以改成bool值，记录有无即可
