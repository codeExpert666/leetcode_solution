package weeklycompetition

func CalculateScore(s string) int64 {
	var res int64
	traveled := make([][]int, 26) // 记录遍历过的字母出现情况
	traveled[s[0]-'a'] = append(traveled[s[0]-'a'], 0)

	for i := 1; i < len(s); i++ {
		mirrorIndex := 'z' - s[i]
		if l := len(traveled[mirrorIndex]); l > 0 {
			res += int64(i - traveled[mirrorIndex][l-1])
			traveled[mirrorIndex] = traveled[mirrorIndex][:l-1] // 标记
		} else { // 没找到镜像才加入，也是为了标记
			traveled[s[i]-'a'] = append(traveled[s[i]-'a'], i)
		}
	}

	return res
}
