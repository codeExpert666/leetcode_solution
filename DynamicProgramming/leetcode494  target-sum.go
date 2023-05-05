package DynamicProgramming

/**
 * 01背包问题应用（dp）
 * 添加符号（正负）使得表达式等于目标值等价于找两个不相交子集，使得这两个子集差为target
 * 于是问题转化为找到所有子集，该子集的和等于（sum + target） / 2，这也就是“分割等和子集”问题，然后再转换成01问题
 * 但与传统“分割等和问题”不同的是，这里的dp数组含义有区别（这主要是因为之前的问题是判断能不能，现在的问题是判断有多少），
 * 现在的dp[j]代表用[0,i-1]物品装满容量为j的背包有多少种方案，递推公式变成了dp[j] = dp[j] + dp[j - nums[i]]
 */
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if abs(target) > sum { // 无论怎么添加符号，得到的表达式的值只能在[-sum, sum]之间
		// 所以这种情况没有方案
		return 0
	}
	if (sum+target)%2 != 0 { // 由于转换成了分割等和问题，故sum+target只能是偶数
		return 0
	}

	bagWeight := (sum + target) / 2 // 获取背包重量
	dp := make([]int, bagWeight+1)  // dp数组及初始化
	// 注意这个初始化，常理来说，dp[0]应初始化为0，但这样的初始化会导致dp数组始终为0。
	// 初始化dp[0]为1可以理解为装满容量为0的背包天然有一种方案（也就是什么都不做）
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagWeight; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
		//fmt.Println(dp)
	}
	return dp[bagWeight]
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
