package DailyCode

func DistinctNames(ideas []string) int64 {
	// 创建一个map来存储每个首字母对应的后缀集合
	groups := make(map[byte]map[string]struct{})
	for _, idea := range ideas {
		first, suffix := idea[0], idea[1:]
		if groups[first] == nil {
			groups[first] = make(map[string]struct{})
		}
		groups[first][suffix] = struct{}{}
	}

	var result int64
	// 遍历所有可能的首字母对
	for firstA, suffixSetA := range groups {
		for firstB, suffixSetB := range groups {
			if firstA == firstB {
				continue
			}
			// 计算两个组之间不重复的后缀数量
			common := 0
			for suffix := range suffixSetA {
				if _, ok := suffixSetB[suffix]; ok {
					common++
				}
			}
			// 计算有效组合数量
			result += int64(len(suffixSetA)-common) * int64(len(suffixSetB)-common)
		}
	}

	return result
}
