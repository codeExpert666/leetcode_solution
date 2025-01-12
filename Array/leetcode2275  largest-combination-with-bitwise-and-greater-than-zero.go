package Array

func LargestCombination(candidates []int) int {
	// 统计每个数的二进制中 1 的出现次数，最大数值为 10**7
	biOneCounts := make([]int, 24)

	for _, c := range candidates {
		// 获取 c 的二进制数，并记录 1 的出现
		for i := 0; c != 0; i++ {
			// if c&1 == 1 {
			// 	biOneCounts[i]++
			// }
			biOneCounts[i] += c & 1 // 上面注释的部分可以简写成一行语句，这种写法不仅简洁，还减少了一次 if 判断，运行效率大大提高！
			c >>= 1
		}
	}

	// 1 的最大出现次数即为最长组合长度
	var res int
	for _, v := range biOneCounts {
		if v > res {
			res = v
		}
	}

	return res
}
