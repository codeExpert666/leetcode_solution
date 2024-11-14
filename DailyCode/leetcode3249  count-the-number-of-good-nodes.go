package DailyCode

func CountGoodNodes(edges [][]int) int {
	// 重构树的存储结构，改为邻接表形式
	// 配以visited数组，标记已访问节点，毕竟是树，防止回头遍历
	adjTable := make([][]int, len(edges)+1)
	visited := make([]bool, len(edges)+1)
	for _, e := range edges {
		adjTable[e[0]] = append(adjTable[e[0]], e[1])
		adjTable[e[1]] = append(adjTable[e[1]], e[0])
	}
	// 通过后续遍历模板确定每个节点是否为好节点，过程中需要计算树中的节点个数
	var res int
	var dfs func(int) int // 返回值是以输入节点为根的树的节点个数
	dfs = func(node int) int {
		visited[node] = true          // 一定先访问当前节点，不然后面遍历子节点会回头
		if len(adjTable[node]) == 0 { // 叶子节点
			res++
			return 1
		}
		var total = 1 // 树节点总个数
		var (         // 用于每个分支的节点数是否相等
			firstCount int
			equal      = true
		)
		for _, next := range adjTable[node] { // 计算每个分支节点个数
			if !visited[next] {
				branchCount := dfs(next)
				if firstCount == 0 { // 第一个分支
					firstCount = branchCount
				} else if branchCount != firstCount { // 后续分支
					equal = false
				}
				total += branchCount
			}
		}
		if equal { // 所有子树节点数相等
			res++
		}
		return total
	}
	dfs(0)
	return res
}

// 同样的思路，官方题解代码更简洁，体现在两点：
// 1、通过向dfs中传入parent节点有效防止回头，省去了visited。毕竟是树，当前节点若回头，只有一个父节点可以走，所以将其记录下来，能够有效防止回头
// 2、叶子节点无需单独处理，由于标记处for循环的限制，能进入循环体，当前节点的孩子节点一定存在，空孩子节点是进不去循环的，从而有效保证递归停止在叶子节点处
func CountGoodNodes1(edges [][]int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	res := 0
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {
		valid := true
		treeSize := 0
		subTreeSize := 0

		for _, child := range g[node] { // 这个for循环
			if child != parent {
				size := dfs(child, node)
				if subTreeSize == 0 {
					subTreeSize = size
				} else if size != subTreeSize {
					valid = false
				}
				treeSize += size
			}
		}
		if valid {
			res++
		}
		return treeSize + 1
	}

	dfs(0, -1)
	return res
}
