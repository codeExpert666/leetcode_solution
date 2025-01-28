package DynamicProgramming

import "math"

// 动态规划：刷表法（根据当前状态去更新其他状态）
// dp[i] 表示到达 nums[i] 的最小跳跃次数
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}

	for i, v := range nums {
		for j := 1; j <= v && i+j < n; j++ {
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}

	return dp[n-1]
}

// 灵神题解：区间贪心 O(n)。
// 详细解析见链接：https://leetcode.cn/problems/jump-game-ii/solutions/2926993/tu-jie-yi-zhang-tu-miao-dong-tiao-yue-yo-h2d4/?envType=daily-question&envId=2025-01-27
// 适合区间跳跃题目
func jump1(nums []int) (ans int) {
	curRight := 0  // 已建造的桥的右端点
	nextRight := 0 // 下一座桥的右端点的最大值
	for i, num := range nums[:len(nums)-1] {
		nextRight = max(nextRight, i+num)
		if i == curRight { // 到达已建造的桥的右端点
			curRight = nextRight // 造一座桥
			ans++
		}
	}
	return
}
