package Simulation

// 看似模拟，实则本质上还是 dp
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := rowIndex; j >= 1; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}
