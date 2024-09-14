package Array

// 解法一：二分法（左闭右闭）
func mySqrt(x int) int {
	// 先处理特殊情况
	if x <= 1 {
		return x
	}
	if x == 2 || x == 3 {
		return 1
	}
	// 后续使用二分法找到2,3...,x/2中平方数小于等于x的最大整数
	low, high := 2, x>>1
	for low <= high {
		mid := low + (high-low)>>1
		if mid <= x/mid { // 防止溢出
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	// 返回high或low-1
	return high
}
