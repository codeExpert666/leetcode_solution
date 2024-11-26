package StringFunc

// 同时遍历两个字符串，保证同等位置处输入的字符数大于名字中的字符
func IsLongPressedName(name string, typed string) bool {
	// i, j分别指向name和typed待比较字符
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] { // 待比较字符不同
			if j == 0 || typed[j] != typed[j-1] {
				// 1、typed的第一个字符就与name不同，一定不匹配
				// 2、typed到达下一字符，但name上一字符还在延续
				// 3、typed与name均到达下一字符，但不相等
				return false
			}
			// 走到这里说明name到达下一字符，typed还在延续上一字符
			// 此时让typed遍历到下一字符即可
			for j < len(typed) && typed[j] == typed[j-1] {
				j++
			}
		} else { // 待比较字符相同，进行下一对字符的比较
			i++
			j++
		}
	}
	// 此时可能存在两种情况：1、name未遍历完 2、typed未遍历完
	if i != len(name) { // name未遍历完，说明typed已经结束，也即未完全匹配
		return false
	}
	for j < len(typed) { // typed未遍历完，只要后续剩余元素均为最后比较元素，则代表成功匹配，反之失败
		if typed[j] != typed[j-1] {
			return false
		}
		j++
	}
	return true
}
