package Simulation

type roomStatus struct {
	isVisited bool // 房间是否访问
	isOdd     bool // 房间访问次数是否为奇数
}

// 模拟推导
// 漏看题目条件：0 <= nextVisit[i] <= i
// 我的解法可以处理 0 <= nextVisit[i] < n 的情况，但当限制在题目条件下，我的解法就稍显复杂了
// 针对题目条件有更特殊的解法
func FirstDayBeenInAllRooms(nextVisit []int) int {
	l, mod := len(nextVisit), int(1e9+7)
	roomsLeft, day := l, -1     // 剩余未访问的房间数，当前日期
	rs := make([]roomStatus, l) // 每个房间访问次数是否为奇数

	// 不断根据规则模拟，直至 roomsLeft 为 0
	roomId := 0 // 从 0 号房间开始
	for roomsLeft > 0 {
		if !rs[roomId].isVisited { // 房间之间未被访问
			roomsLeft--
		}
		rs[roomId].isVisited = true          // 房间标记为已访问
		rs[roomId].isOdd = !rs[roomId].isOdd // 改变访问次数
		day = (day + 1) % mod                // 日期更新到当天
		// 更新下一次访问的 roomId
		if rs[roomId].isOdd { // roomId 的访问次数为奇数
			roomId = nextVisit[roomId]
		} else {
			roomId = (roomId + 1) % l
		}
	}
	return day
}

// 动态规划
// 考虑到题目条件：0 <= nextVisit[i] <= i，可有：
// 1、访问完所有房间的日期等价于第一次访问房间 n-1 的日期
// 2、第一次访问 i 房间时，上一次访问的房间一定是 i-1，且是第二次访问该房间
// 3、第一次访问 i 房间时，其前面的房间一定都被访问了偶数次
// 于是，dp[i] 可定义为第一次访问房间 i 的日期，那么有：
// dp[i] = dp[i-1] + 1 + dp[i-1] - dp[nextVisit[i-1]] + 1
// 其中，dp[i-1] + 1 是再一次到达 nextVisit[i-1] 房间的日期，可以推得，
// 此时 nextVisit[i-1] 的访问次数为奇数次，i 前的其余房间访问次数均为偶数次，
// 这种状态与第一次到达 nextVisit[i-1] 是一样的。所以，nextVisit[i-1] 再次到达
// i-1 所花费的天数与从第一次到达 nextVisit[i-1] 再第一次到达 i-1 所花费的天数一致，
// 也即，dp[i-1] - dp[nextVisit[i-1]]。此时已经第二次到达 i-1，下一天第一次到达 i。
func FirstDayBeenInAllRooms1(nextVisit []int) int {
	l, mod := len(nextVisit), int(1e9+7)
	dp := make([]int, l)
	dp[0] = 0
	for i := 1; i < l; i++ {
		dp[i] = (dp[i-1]<<1 + 2 - dp[nextVisit[i-1]] + mod) % mod
	}
	return dp[l-1]
}
