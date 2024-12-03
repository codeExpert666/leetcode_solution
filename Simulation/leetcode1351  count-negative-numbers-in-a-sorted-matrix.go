package Simulation

// 抓住一点：每一行第一个负数所在的列一定是非严格递减的。
// 所以可以倒序遍历，逐个找到每一行的第一个负数，累加每一行负数的个数
func countNegatives(grid [][]int) int {
	var res int
	m, n := len(grid), len(grid[0])
	pos := n // 记录每一行第一个负数的位置
	row := 0
	for ; row < m && pos > 0; row++ {
		for pos > 0 && grid[row][pos-1] < 0 { // 找到每一行的第一个负数
			pos--
		}
		res += n - pos
	}
	// 退出循环的原因可能是 pos==0，此时剩余行一整行都是负数
	return res + (m-row)*n
}
