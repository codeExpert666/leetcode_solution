package Array

// 螺旋数组层次排布，故按层确定数字位置
func generateMatrix(n int) [][]int {
	// n为奇数时，有n>>1 + 1层；n为偶数时，有n>>1层
	// 故先考虑前n>>1层，奇数的最后一层单独处理
	res := make([][]int, n) // 初始化二维切片
	for i := range res {    // range返回索引和值，只关心索引，值可以省略
		res[i] = make([]int, n)
	}
	start := 0                              // 记录每一层开始的数字，减一处理是为了后续方便
	for layer := 0; layer < n>>1; layer++ { // 层数从0开始计数，方便处理
		for i := 1; i <= 4*n-8*layer-4; i++ { // 第layer+1层共有4*(n-2*layer)-4个元素
			if i <= n-2*layer { // 数学方法推导位置
				res[layer][layer+i-1] = start + i
			} else if i <= 2*n-4*layer-1 {
				res[3*layer+i-n][n-layer-1] = start + i
			} else if i <= 3*n-6*layer-2 {
				res[n-layer-1][3*n-5*layer-i-2] = start + i
			} else {
				res[4*n-7*layer-i-3][layer] = start + i
			}
		}
		start += 4*n - 8*layer - 4 // 更新下一层的起始数字
	}
	// n为奇数单独处理最中间一层（一个数字n^2）
	if n%2 != 0 {
		res[n>>1][n>>1] = n * n
	}
	return res
}

// 与前一解法思路一致，但写法更为简单优美，逻辑也更清晰
func generateMatrix1(n int) [][]int {
	loop, mid := n/2, n/2
	startx, starty, offset := 0, 0, 1
	count := 1

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for loop > 0 {
		i, j := starty, startx
		for ; i < n-offset; i++ {
			res[startx][i] = count
			count++
		}
		for ; j < n-offset; j++ {
			res[j][i] = count
			count++
		}
		for ; i > starty; i-- {
			res[j][i] = count
			count++
		}
		for ; j > startx; j-- {
			res[j][i] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 == 1 {
		res[mid][mid] = count
	}
	return res
}
