package String

// 这里将map改为[26]int能减少时间复杂度
func CommonChars(words []string) []string {
	mapList := [2]map[byte]int{} // 两个hash表轮换使用，目的是记录当前为止重复使用的字符
	// 注意nil map无法直接使用，必须先初始化
	mapList[0] = make(map[byte]int)
	// 首先记录第一个字符串的所有字符
	for i := 0; i < len(words[0]); i++ {
		mapList[0][words[0][i]]++
	}
	// 遍历后续字符串，每遍历一个字符串，更新当前的重复字符
	for i := 1; i < len(words); i++ {
		deleteMap, useMap := (i-1)%2, i%2    // 新的重复字符更新到useMap中，deleteMap舍弃
		mapList[useMap] = make(map[byte]int) // 清空无用数据
		for j := 0; j < len(words[i]); j++ {
			if v := mapList[deleteMap][words[i][j]]; v > 0 { // 重复字符在新单词中也出现
				mapList[useMap][words[i][j]]++
				mapList[deleteMap][words[i][j]] = v - 1
			}
		}
	}
	// 遍历最后使用的hash表，提取所有重复字符
	var res []string
	for k, v := range mapList[(len(words)-1)%2] {
		for v > 0 {
			res = append(res, string(k))
			v--
		}
	}
	return res
}

// 还有一个思路：统计出搜索字符串里26个字符的出现的频率，然后取每个字符频率最小值，最后转成输出格式就可以了
// 这种思路也需要两个hash表同时工作
