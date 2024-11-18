package DailyCode

// 获取以（x,y）为中心的九宫格的左上角顶点
func getLeftTop(x, y int) (int, int) {
	if x > 0 {
		x--
	}
	if y > 0 {
		y--
	}
	return x, y
}

// 获取以（x,y）为中心的九宫格的右下角顶点
func getRightBottom(x, y, row, col int) (int, int) {
	if x < row-1 {
		x++
	}
	if y < col-1 {
		y++
	}
	return x, y
}

// ImageSmoother 本题中存在大量重复计算，为尽可能减少重复计算，采用前缀和
func ImageSmoother(img [][]int) [][]int {
	row, col := len(img), len(img[0])
	// 将原数组更改为前缀和数组，其中img[i][j]表示左上角元素之和
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i > 0 && j > 0 {
				img[i][j] += img[i-1][j] + img[i][j-1] - img[i-1][j-1] // 注意剔除重复计算
			} else if i > 0 { // j == 0
				img[i][j] += img[i-1][j]
			} else if j > 0 { // i == 0
				img[i][j] += img[i][j-1]
			}
		}
	}
	// 利用前缀和数组计算平滑结果
	res := make([][]int, row)
	for i := 0; i < row; i++ {
		res[i] = make([]int, col)
		for j := 0; j < col; j++ {
			ltx, lty := getLeftTop(i, j)               // 九宫格左上角顶点
			rbx, rby := getRightBottom(i, j, row, col) // 九宫格右下角顶点
			n := (rbx - ltx + 1) * (rby - lty + 1)     // 九宫格像素个数
			if ltx == 0 && lty == 0 {                  // 左上角顶点是第一个像素点
				res[i][j] = img[rbx][rby] / n
			} else if ltx == 0 { // 九宫格紧贴图片上边界
				res[i][j] = (img[rbx][rby] - img[rbx][lty-1]) / n
			} else if lty == 0 { // 九宫格紧贴图片左边界
				res[i][j] = (img[rbx][rby] - img[ltx-1][rby]) / n
			} else { // 九宫格在图片中
				res[i][j] = (img[rbx][rby] - img[rbx][lty-1] - img[ltx-1][rby] + img[ltx-1][lty-1]) / n
			}
		}
	}
	return res
}

// 优化思路：为了减少边界判断，可以分别在img的左侧和上方再加两个0层，构建一个更大的前缀和数组
// 缺点是，前缀和数组中的下标与img中不对应，且消耗更多的空间。
