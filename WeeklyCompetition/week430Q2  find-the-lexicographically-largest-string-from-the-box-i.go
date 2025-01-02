package weeklycompetition

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	// 先找到所有字典序最大字母对应的下标
	var maxChar byte
	pos := make([]int, 0)
	for i := 0; i < len(word); i++ {
		if word[i] == maxChar {
			pos = append(pos, i)
		} else if word[i] > maxChar {
			maxChar = word[i]
			clear(pos)
			pos = append(pos, i)
		}
	}
	// 最大子串只可能以最大字母为首，且长度受到 numFriends 的限制
	var res string
	maxLen := len(word) - numFriends + 1 // 子串理论最大长度
	for _, p := range pos {
		end := min(len(word), p+maxLen) // 子串尾部不能超出字符串
		if s := word[p:end]; s > res {
			res = s
		}
	}
	return res
}
