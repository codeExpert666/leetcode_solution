package StringFunc

// 获取KMP算法中待匹配串的Next数组
func getNext(s string) []int {
	next := make([]int, len(s))
	for i := 1; i < len(next); i++ {
		// next数组的后续值可由先前值导出
		// 这里next[j]的值指的是s[:j+1]的最大匹配前后缀的长度
		// 由于数组从0编后续字符串匹配时待比较的字符下标
		for j := i - 1; j >= 0; j = next[j] - 1 {
			if s[i] == s[next[j]] {
				next[i] = next[j] + 1
				break
			}
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	// 获取next数组
	next := getNext(needle)
	// i,j分别指向haystack和needle的待比较字符
	for i, j := 0, 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			// j根据next不断回退，考虑到next[0]=0会造成无限循环，故回退到j=0时会跳出循环
			// j=0时，haystack[i]与needle[j]的关系并未判断
			j = next[j]
		}
		// 故而这里的判断包含了j=0的情况
		if haystack[i] == needle[j] {
			j++
			if j == len(needle) {
				return i - len(needle) + 1
			}
		}
	}
	return -1
}
