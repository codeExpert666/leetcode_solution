package DailyCode

// 二重循环遍历+剪枝 会超时
// 单重循环非常巧妙，重点在于把表达式看成两个独立的部分（values[i]+i）与（values[j]-j）
// 为什么不将表达式做出其他拆分组合呢？因为其他组合都不可避免在子表达式中同时存在i与j
// 而单重循环一次只能控制一个变量的变化，对于包含两个变量的表达式，无法考虑详尽其所有可能的取值
func MaxScoreSightseeingPair(values []int) int {
	maxScore := 0
	max_i_plus_value := values[0] // 初始化为第一个景点的值加上它的索引

	for j := 1; j < len(values); j++ {
		// 计算当前景点与之前景点的最大得分
		maxScore = Max(maxScore, max_i_plus_value+values[j]-j)
		// 更新 max_i_plus_value
		max_i_plus_value = Max(max_i_plus_value, values[j]+j)
	}

	return maxScore
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
