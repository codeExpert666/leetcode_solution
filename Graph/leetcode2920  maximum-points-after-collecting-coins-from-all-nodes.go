package main

// 思路和代码均是正确的，但有一个细节没做好，导致 memo 占用大量空间，从而内存溢出。
// 题给数据 coins[i] 最多在右移 14 位之后就会变成 0。
// 也就是说，对于dfs(node, count)，当 count >= 14 时，该子树所有节点的金币均为0。
// 对于该子树，完全使用方案一获得的积分为 -nk，完全使用方案二获得的积分为 0，交替使用两种方案的积分在二者之间。
// 由于 k >= 0，所以可以直接得到该子树的最大积分为 0，从而不用继续递归。
// 这样既进行了剪枝，同时也可以减小 memo 的记录负担（不用记录 count >= 14 的状态），
// 使得 memo 不用再开辟 nodeNum * nodeNum 的空间去存储数据，极大降低了空间复杂度。
func maximumPoints(edges [][]int, coins []int, k int) int {
	nodeNum := len(edges) + 1
	// 先将 edges 转化为邻接表形式
	adjTable := make([][]int, nodeNum)
	for _, e := range edges {
		n1, n2 := e[0], e[1]
		adjTable[n1] = append(adjTable[n1], n2)
		adjTable[n2] = append(adjTable[n2], n1)
	}
	visited := make([]bool, nodeNum) // 用于遍历

	// 记忆化搜索
	memo := make([][14]int, nodeNum)
	for i := 0; i < nodeNum; i++ {
		for j := 0; j < 14; j++ {
			// 对于任意一棵树，完全按照方案二一定可以获取到非负值积分。
			// 所以 memo 中任意状态的实际值一定是非负值
			// 故而 -1 是无效值
			memo[i][j] = -1
		}
	}

	// 定义递归函数
	// 以 node 为根的子树在金币减少 count 次后，所能获得的最大积分
	var dfs func(int, int) int
	dfs = func(node int, count int) int {
		if count >= 14 { // 剪枝
			return 0
		}
		// 该状态已访问过
		if memo[node][count] != -1 {
			return memo[node][count]
		}
		// 分别计算两种方式所能获得的积分
		var first, second int = coins[node]>>count - k, coins[node] >> (count + 1)
		for _, neighbor := range adjTable[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				first += dfs(neighbor, count)
				second += dfs(neighbor, count+1)
				visited[neighbor] = false
			}
		}
		memo[node][count] = max(first, second)
		return memo[node][count]
	}

	visited[0] = true
	return dfs(0, 0)
}
