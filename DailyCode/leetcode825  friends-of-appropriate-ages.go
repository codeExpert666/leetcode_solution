package DailyCode

func NumFriendRequests(ages []int) int {
	// 构造前缀和数组prefixSum，其中prefixSum[i]代表ages中小于等于i的元素个数
	prefixSum := [121]int{}  // age最大值为120
	for _, v := range ages { // 先统计ages中元素出现情况
		prefixSum[v]++
	}
	for i := 1; i < len(prefixSum); i++ { // 滚动累加，形成前缀和
		prefixSum[i] += prefixSum[i-1]
	}
	// 根据题意，假设一个人年龄为x，其只与年龄在(x/2+7, x]范围内的人发送好友请求
	var res int
	for _, age := range ages {
		if age>>1+7 < age {
			res += prefixSum[age] - prefixSum[age>>1+7] - 1 // 去掉自己
		}
	}
	return res
}
