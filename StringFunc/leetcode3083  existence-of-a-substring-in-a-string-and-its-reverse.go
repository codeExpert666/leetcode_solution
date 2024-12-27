package StringFunc

// 问题等价于能否找到一个长度为 2 的子串，其反转串也在 s 中
// 采用两数之和的思路
func IsSubstringPresent(s string) bool {
	subString := [26][26]bool{} // 用二维数组代替 hash 表
	for i := 0; i < len(s)-1; i++ {
		first, second := s[i]-'a', s[i+1]-'a' // 连续两个字符在数组中的下标
		// 顺序很重要，先标记后寻找，这是为了应对子串两个字符相等的情况
		subString[first][second] = true
		if subString[second][first] {
			return true
		}
	}
	return false
}
