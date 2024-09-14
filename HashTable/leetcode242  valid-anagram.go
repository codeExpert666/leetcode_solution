package HashTable

// 常规hash表
func isAnagram(s string, t string) bool {
	// 先将字符串s中的内容存在hash表中
	// key为字符，value为该字符出现的次数
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	// 遍历t，每出现一个字符，判断是否在hash表中，不在则不满足条件
	// 在则减少出现次数，若最终所有字符出现次数为0，则符合条件
	for _, c := range t {
		if _, ok := m[c]; ok { // 存在
			m[c]--
			if m[c] < 0 { // 次数为负说明两个字符串对应字符个数不等
				return false
			}
		} else { // 不存在
			return false
		}
	}
	// 检查hash表中字符次数是否都降为0
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// 字符均为小写字母，故可将hash表简化为长度为26的数组
func isAnagram1(s string, t string) bool {
	var charCount [26]int // 把数组当作简单的hash表（主要是因为字符串里都是小写字母）

	for _, c := range s {
		charCount[c-rune('a')]++
	}
	for _, c := range t {
		charCount[c-rune('a')]--
	}

	return charCount == [26]int{}
}
