package DailyCode

type NeighborSum struct {
	scale    int      // 方便判断相邻元素是否出界
	NumToPos [][2]int // 方便从元素定位到位置
	grid     [][]int  // 方便从位置定位到元素
}

func Constructor(grid [][]int) NeighborSum {
	n := len(grid)
	ns := NeighborSum{n, make([][2]int, n*n), grid}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ns.NumToPos[grid[i][j]] = [2]int{i, j}
		}
	}
	return ns
}

// 可以写一个方法以当前元素和邻居列表为输入，获取邻居的和。这样可以减少AdjacentSum与DiagonalSum中的代码重复
func (ns *NeighborSum) AdjacentSum(value int) int {
	var res int
	neighbors := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, neighbor := range neighbors {
		pos := ns.NumToPos[value]
		newx, newy := pos[0]+neighbor[0], pos[1]+neighbor[1]
		if newx < 0 || newx >= ns.scale || newy < 0 || newy >= ns.scale {
			continue
		}
		res += ns.grid[newx][newy]
	}
	return res
}

func (ns *NeighborSum) DiagonalSum(value int) int {
	var res int
	neighbors := [4][2]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
	for _, neighbor := range neighbors {
		pos := ns.NumToPos[value]
		newx, newy := pos[0]+neighbor[0], pos[1]+neighbor[1]
		if newx < 0 || newx >= ns.scale || newy < 0 || newy >= ns.scale {
			continue
		}
		res += ns.grid[newx][newy]
	}
	return res
}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
