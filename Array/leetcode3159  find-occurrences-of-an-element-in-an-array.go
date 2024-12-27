package Array

func OccurrencesOfElement(nums []int, queries []int, x int) []int {
	// 先找到所有 x 的位置
	xPos := make([]int, 0)
	for i, v := range nums {
		if v == x {
			xPos = append(xPos, i)
		}
	}
	// 响应查询
	l := len(xPos)
	ans := make([]int, len(queries))
	for i, q := range queries {
		if q > l {
			ans[i] = -1
		} else {
			ans[i] = xPos[q-1]
		}
	}
	return ans
}
