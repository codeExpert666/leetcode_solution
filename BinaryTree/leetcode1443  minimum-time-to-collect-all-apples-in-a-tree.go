package BinaryTree

func MinTime(n int, edges [][]int, hasApple []bool) int {
	// 将 edges 转换为邻接表形式，注意题目的输入是无向边
	adjTable := make([][]int, n)
	for _, e := range edges {
		n1, n2 := e[0], e[1]
		adjTable[n1] = append(adjTable[n1], n2)
		adjTable[n2] = append(adjTable[n2], n1)
	}
	// visited 标记访问过的节点
	visited := make([]bool, n)
	// travel 返回以指定节点为根的树中是否存在苹果，以及获取所有苹果所需的最短时间
	var travel func(int) (bool, int)
	travel = func(node int) (apple bool, time int) {
		if len(adjTable[node]) == 0 { // 叶子节点
			apple = hasApple[node]
			return
		}
		// 访问根节点
		visited[node] = true
		apple = hasApple[node]
		// 访问子树
		for _, next := range adjTable[node] {
			if !visited[next] {
				if a, t := travel(next); a {
					time += t + 2 // 子树获取苹果时间、前往子树根节点时间、子树根节点返回时间
					apple = a
				}
			}
		}
		return
	}
	_, t := travel(0)
	return t
}
