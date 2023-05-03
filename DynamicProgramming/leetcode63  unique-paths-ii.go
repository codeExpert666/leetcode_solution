package DynamicProgramming

/**
 * 动态规划，利用一维数组记录，节省空间
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 { // 出发点就有障碍物
		return 0
	}
	dp := make([]int, len(obstacleGrid[0]))
	for i := 0; i < len(dp); i++ { // 初始化，遇到障碍物将剩余位置的路径数置为零
		if obstacleGrid[0][i] == 1 {
			dp[i] = 0
		} else if i == 0 {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	for j := 1; j < len(obstacleGrid); j++ {
		for i := 0; i < len(dp); i++ { // 注意，从0开始遍历，因为第一列可能存在障碍物，dp[0]需要改变
			if obstacleGrid[j][i] == 1 { // 当前位置有障碍物
				dp[i] = 0
			} else if i != 0 { // 当前位置没有障碍物，0处不做处理可以保证第一列当前位置的路径数与该列前一位置的路径数一致
				dp[i] += dp[i-1]
			}
		}
	}

	return dp[len(dp)-1]
}
