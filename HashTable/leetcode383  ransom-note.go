package HashTable

func canConstruct(ransomNote string, magazine string) bool {
	// 记录magazine中字母出现情况
	charCount := [26]int{} // 手工hash
	for _, c := range magazine {
		charCount[c-rune('a')]++
	}
	// 遍历ransomNote，检查是否能被hash表覆盖
	for _, c := range ransomNote {
		if charCount[c-rune('a')] == 0 { // 字母出现次数使用完毕
			return false
		}
		charCount[c-rune('a')]--
	}
	return true
}
