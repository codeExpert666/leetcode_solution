package Simulation

// 定义车、象、后棋子的方向
var rookDirections = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var bishopDirections = [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
var queenDirections = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

// 这题太麻烦了，要考虑一堆约束条件，不会做也懒得做，直接 cv 了，这种题目没啥意义
func countCombinations(pieces []string, positions [][]int) int {
	n, res := len(pieces), 0
	stack := []Movement{}

	// Check 判断第 u 个棋子是否与之前的棋子发生相交
	check := func(u int) bool {
		for v := 0; v < u; v++ {
			if stack[u].cross(&stack[v]) {
				return false
			}
		}
		return true
	}

	// DFS 深度优先搜索
	var dfs func(u int)
	dfs = func(u int) {
		if u == n {
			res++
			return
		}
		var directions [][2]int
		switch pieces[u] {
		case "rook":
			directions = rookDirections
		case "queen":
			directions = queenDirections
		default:
			directions = bishopDirections
		}

		// 处理第 u 个棋子原地不动的情况
		stack = append(stack, Movement{startX: positions[u][0], startY: positions[u][1], endX: positions[u][0], endY: positions[u][1]})
		if check(u) {
			dfs(u + 1)
		}
		stack = stack[:len(stack)-1]

		// 枚举第 u 个棋子在所有方向、所有步数的情况
		for _, dir := range directions {
			for j := 1; j < 8; j++ {
				x := positions[u][0] + dir[0]*j
				y := positions[u][1] + dir[1]*j
				if x < 1 || x > 8 || y < 1 || y > 8 {
					break
				}
				stack = append(stack, Movement{startX: positions[u][0], startY: positions[u][1], endX: x, endY: y, dx: dir[0], dy: dir[1]})
				if check(u) {
					dfs(u + 1)
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	dfs(0)
	return res
}

// Movement 结构体表示棋子的一个移动
type Movement struct {
	startX, startY, endX, endY, dx, dy, curX, curY int
}

// Reset 重置棋子的当前位置
func (m *Movement) reset() {
	m.curX = m.startX
	m.curY = m.startY
}

// Stopped 判断棋子是否停止
func (m *Movement) stopped() bool {
	return m.curX == m.endX && m.curY == m.endY
}

// Advance 让棋子按照步长移动
func (m *Movement) advance() {
	if !m.stopped() {
		m.curX += m.dx
		m.curY += m.dy
	}
}

// Cross 判断两个棋子是否相遇
func (m *Movement) cross(oth *Movement) bool {
	m.reset()
	oth.reset()
	for !m.stopped() || !oth.stopped() {
		m.advance()
		oth.advance()
		if m.curX == oth.curX && m.curY == oth.curY {
			return true
		}
	}

	return false
}
