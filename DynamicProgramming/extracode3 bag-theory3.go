package DynamicProgramming

import "log"

/**
 * 多重背包问题，具体可以描述如下：
 *   有N种物品和一个容量为V的背包。第i种物品最多有Mi件可用，每件耗费的空间是Ci，价值是Wi。
 *   求解将哪些物品装入背包可使这些物品的耗费的空间 总和不超过背包容量，且价值总和最大。
 * 多重背包的解决方式是将其转化为 01 背包，每件物品最多有Mi件可用，把Mi件摊开，其实就是一个01背包问题了。
 */
func multiplePack(weight, value, nums []int, bagWeight int) int {
	// 将多重背包转换为 01 背包
	for i := 0; i < len(nums); i++ {
		for nums[i] > 1 {
			weight = append(weight, weight[i])
			value = append(value, value[i])
			nums[i]--
		}
	}

	log.Println(weight)
	log.Println(value)

	res := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			res[j] = max(res[j], res[j-weight[i]]+value[i])
		}
		log.Println(res)
	}

	return res[bagWeight]
}
