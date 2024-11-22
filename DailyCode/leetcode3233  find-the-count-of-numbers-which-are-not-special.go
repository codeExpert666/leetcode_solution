package DailyCode

import "math"

// NonSpecialCount 通过数学推理，“特殊数字”一定是质数的平方
// 问题转化为1,2...,int(sqrt(r))中共有多少个质数的平方在[l,r]范围内
// 除去这些数字，[l,r]剩余数的个数即为所求
func NonSpecialCount(l int, r int) int {
	bound := int(math.Sqrt(float64(r)))
	// 利用埃拉托色尼筛法获取1,2...,int(sqrt(r))的所有质数，并逐个判断平方是否在[l,r]中
	isPrime := make([]bool, bound+1) // 标记任意数字是否为质数
	for i := 2; i <= bound; i++ {    // 初始化
		isPrime[i] = true
	}
	count := 0                    // 记录平方在[l,r]范围内的质数个数
	for i := 2; i <= bound; i++ { // 从2开始，遍历数组。如果isPrime[i]为true，则将(i)的所有倍数(从(i^2)开始）标记为false
		tmp := i * i
		if isPrime[i] {
			if tmp >= l { // 统计“特殊数字”个数
				count++
			}
			for ; tmp <= bound; tmp += i { // 埃筛法标记非质数
				isPrime[tmp] = false
			}
		}
	}
	return r - l + 1 - count // 非“特殊数字”个数
}
