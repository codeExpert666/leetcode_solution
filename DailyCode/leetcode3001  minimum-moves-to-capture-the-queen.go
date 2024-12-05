package DailyCode

// 无论车还是象，在没有遮挡的情况下，捕获后最多只需要2步
// 对于车，如果其与后在一条直线上，且被象遮挡，那么也是 2 步：象移开，车移动
// 对于车，如果其与后不在一条直线上，那么最短路径有两条，象无论怎么遮挡只能挡一条，所以还是 2 步
// 所以问题转化为，找出能一步捕获后的情况，除此之外，都是 2 步
func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	// 若车一步捕获后，则二者在一条直线，且无遮挡
	if a == e && (c != e || d > max(b, f) || d < min(b, f)) { // 车后在同一行，且象不遮挡
		return 1
	}
	if b == f && (d != f || c > max(a, e) || c < min(a, e)) { // 车后在同一列，且象不遮挡
		return 1
	}
	// 若象一步捕获后，则二者在一条斜线，且无遮挡
	ac, bd, ec, fd, ae, bf := a-c, b-d, e-c, f-d, a-e, b-f
	if abs(ec) == abs(fd) && (abs(ac) != abs(bd) || (ae)*(-ec) <= 0 || (bf)*(-fd) <= 0 || (ac)*(ec) <= 0 || (bd)*(fd) <= 0) {
		return 1
	}
	return 2
}
