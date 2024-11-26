package StringFunc

func NumSplits(s string) int {
	var res int
	// 先统计整个字符串中不同字符的种类与数量
	category, count := 0, [26]int{}
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if count[index] == 0 {
			category++
		}
		count[index]++
	}
	// 遍历分割位置，逐步检查种类是否相同
	leftCate, leftCount := 0, [26]bool{} // 记录左侧字符的种类
	for i := 0; i < len(s)-1; i++ {
		index := s[i] - 'a'
		// 左侧
		if !leftCount[index] {
			leftCate++
		}
		leftCount[index] = true
		// 右侧
		count[index]--
		if count[index] == 0 {
			category--
		}
		// 判断左右两侧
		if leftCate == category {
			res++
		}
	}
	return res
}
