package StringFunc

// 滑动窗口，本题属于有思路但不会写，一眼就看出是滑窗，但有两点不清晰：
// 1、本题窗口的滑动依据显然是子串的覆盖，但怎么判断覆盖没想清楚
// 2、窗口在滑动过程中如何累积符合条件的子串个数也没想清楚
// 以下解答参考的灵神的题解
func validSubstringCount(word1 string, word2 string) int64 {
	if len(word1) < len(word2) {
		return 0
	}

	var res int64
	// diff 与 less 共同判断覆盖
	var diff [26]int // word2 中字母出现次数与窗口中字母出现次数之差
	for i := 0; i < len(word2); i++ {
		diff[word2[i]-'a']++
	}
	var less int // 窗口中还有多少字母的出现次数比 word2 少
	for _, v := range diff {
		if v != 0 {
			less++
		}
	}

	left := 0 // 窗口左边界
	for i := 0; i < len(word1); i++ {
		diff[word1[i]-'a']--
		if diff[word1[i]-'a'] == 0 {
			less--
		}
		for less == 0 {
			if diff[word1[left]-'a'] == 0 { // 注意这两个顺序
				less++
			}
			diff[word1[left]-'a']++ // 注意这两个顺序，这个放在前会导致 diff[word1[left]-'a'] > 0 的情况出错
			left++
		}
		// 这里累加符合条件的子串个数
		res += int64(left) // 此时窗口恰好是不符合条件的子串，所以以小于 left 的下标为左边界的窗口均符合条件，共有 left 个
	}

	return res
}
