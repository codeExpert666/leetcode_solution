package StringFunc

func CountKeyChanges(s string) int {
	var res int
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		ss[i] = upperToLower(ss[i])
		if i > 0 && ss[i] != ss[i-1] {
			res++
		}
	}
	return res
}

// 将大写字母转换为小写。如果已经是小写字母，则不做任何改变
func upperToLower(c byte) byte {
	if c < 'a' { // 大写字母
		c += 'a' - 'A'
	}
	return c
}
