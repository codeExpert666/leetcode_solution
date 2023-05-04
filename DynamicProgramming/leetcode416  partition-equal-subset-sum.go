package DynamicProgramming

/**
 * 01背包问题的应用（dp）
 * 本题的关键是如何套用01背包，所谓找到两个子集和相等，也即找到一个子集的和为sum/2
 * 于是，可进行如下转化：
 * 1、背包的体积为sum / 2
 * 2、背包要放入的商品（集合里的元素）重量为 元素的数值，价值也为元素的数值
 * 3、背包如果正好装满，说明找到了总和为 sum / 2 的子集。
 * 4、背包中每一个元素是不可重复放入。
 */
func canPartition(nums []int) bool {
	// 获取所有元素总和
	var sum int
	for _, v := range nums {
		sum += v
	}
	// 先对总和进行初步判断（总和是奇数，则一定不存在）
	if sum%2 != 0 {
		return false
	}
	// dp过程
	bagWeight := sum / 2 // 背包总重
	dp := make([]int, bagWeight+1)

	for i := 0; i < len(nums); i++ {
		for j := bagWeight; j >= nums[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			//// 仅关注递推表最后一列的值，一旦最大价值与背包总重相同时，便可以划分
			//if j == bagWeight && dp[j] == j {
			//	return true
			//}
		}
	}
	//// 找不到与背包总重相同的最大价值，则不可以划分
	//return false
	// 也可以直接在最后判断，代码更简洁
	// 这是因为：1、如果前面几个物品就已经能够达到最大价值（背包总重），则一定有dp[bagWeight]=bagWeight
	// 2、若dp[bagWeight]=bagWeight，则说明一定有几个物品能够使得价值达到背包总重
	// 所以这是充要条件，再最后判断没有问题
	return dp[bagWeight] == bagWeight
}
