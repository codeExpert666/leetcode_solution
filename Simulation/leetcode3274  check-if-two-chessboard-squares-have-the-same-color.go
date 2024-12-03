package Simulation

// 判断给定坐标是否为黑色
func isBlack(coordinate string) bool {
	x, y := coordinate[0]-'a'+1, coordinate[1]-'0'
	return x%2 == y%2
}

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	return isBlack(coordinate1) == isBlack(coordinate2)
}
