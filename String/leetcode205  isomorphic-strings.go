package String

func IsIsomorphic(s string, t string) bool {
	// m1是字符串s到字符串t的映射；m2是字符串t到字符串s的映射
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if c, exist := m1[s[i]]; exist && c != t[i] { // s中的相同字符在t中映射到不同字符
			return false
		}
		if c, exist := m2[t[i]]; exist && c != s[i] { // s中的不同字符在t中映射到相同字符
			return false
		}
		// 添加新映射规则
		m1[s[i]] = t[i]
		m2[t[i]] = s[i]
	}
	return true
}
