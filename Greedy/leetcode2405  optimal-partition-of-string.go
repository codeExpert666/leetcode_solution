package Greedy

// 每个子串在满足限制条件的情况下尽可能长
func partitionString(s string) int {
	var res int
	eMap := [26]bool{} // 标记子串中字符出现情况
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if eMap[index] { // 当前子串中已存在该字符，需要截断
			res++
			eMap = [26]bool{} // hash表重置
		}
		eMap[index] = true // 统计字符
	}
	return res + 1 // 最后一个子串未统计
}
