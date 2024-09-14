package Array

func backspaceCompare(s string, t string) bool {
	sAfter := removeBackspace(s)
	tAfter := removeBackspace(t)
	return sAfter == tAfter
}

// 快慢双指针法（还原退格后的字符串）
func removeBackspace(s string) string {
	b := []byte(s)
	slow := 0
	for fast := 0; fast < len(b); fast++ {
		if b[fast] == '#' && slow > 0 {
			slow--
		} else if b[fast] != '#' {
			b[slow] = b[fast]
			slow++
		}
	}
	return string(b[:slow])
}
