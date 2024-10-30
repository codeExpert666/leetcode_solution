package DailyCode

// 设置了simpleStr，目的是用空间换时间
func MaximumSubsequenceCount(text string, pattern string) int64 {
	var simpleStr []byte     // 用于存储text删减掉无关字母后剩余串
	var count1, count2 int64 // 记录pattern串的两个字符在text中的出现次数
	// 删减字符串并统计出现次数
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[0] || text[i] == pattern[1] {
			simpleStr = append(simpleStr, text[i])
			if text[i] == pattern[0] {
				count1++
			}
			if text[i] == pattern[1] {
				count2++
			}
		}
	}
	// 统计text中pattern的出现次数
	var res int64
	for i, c2 := 0, count2; c2 > 0 && i < len(simpleStr); i++ {
		if simpleStr[i] == pattern[1] { // 为了应对pattern[0]==pattern[1]的情况
			c2--
		}
		if simpleStr[i] == pattern[0] {
			res += c2
		}
	}
	// pattern[0]与pattern[1]分别放在text的开头和末尾能取得模式串出现次数的最大值
	// pattern[0]的加入取决于原串中pattern[1]的数量，反之亦然
	if count1 > count2 {
		return res + count1
	} else {
		return res + count2
	}
}

// 同样思路，一次遍历即可
func MaximumSubsequenceCount1(text string, pattern string) int64 {
	var res int64
	var base int64           // 用于模式串出现次数的累加基数
	var count1, count2 int64 // 记录pattern串的两个字符在text中的出现次数
	// 统计pattern、pattern[0]、pattern[1]在text中的出现次数
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[1] { // 先判断pattern[1]为了应对pattern[0]==pattern[1]的情况
			res += base
			count2++
		}
		if text[i] == pattern[0] {
			base++
			count1++
		}
	}
	// pattern[0]与pattern[1]分别放在text的开头和末尾能取得模式串出现次数的最大值
	// pattern[0]的加入取决于原串中pattern[1]的数量，反之亦然
	if count1 > count2 {
		return res + count1
	} else {
		return res + count2
	}
}
