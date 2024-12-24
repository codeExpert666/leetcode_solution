package StringFunc

// 抓住 s 的长度一定是同位子串长度的整数倍，由于 s 的长度最大为 10**5，可以枚举
func MinAnagramLength(s string) int {
	n := len(s)
	// 先统计 s 中字符的频数
	totalFq := [26]int{}
	for i := 0; i < n; i++ {
		totalFq[s[i]-'a']++
	}
	// 枚举子串长度
Outer:
	for k := 1; k <= n>>1; k++ {
		if n%k == 0 { // s 的长度必须是 k 的整数倍
			subNum := n / k // 子串个数
			// 获取子串的字符频数
			idealSubFq := totalFq
			for i := 0; i < 26; i++ {
				if idealSubFq[i]%subNum != 0 { // 某个字符的个数无法均分到每个子串中
					continue Outer
				}
				idealSubFq[i] /= subNum
			}
			// 判断每个子串是否按照预计字符频数组成
			for i := 0; i < subNum; i++ {
				realSubFq := [26]int{}
				for j := 0; j < k; j++ { // 统计子串真实字符组成
					realSubFq[s[i*k+j]-'a']++
				}
				if realSubFq != idealSubFq {
					continue Outer
				}
			}
			return k
		}
	}
	return n
}
