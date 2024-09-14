package Array

// 解法一：二分法（左闭右闭）
func isPerfectSquare(num int) bool {
	if num <= 1 { // 处理特殊情况
		return true
	}
	low, high := 1, num
	for low <= high {
		mid := low + (high-low)>>1
		if mid < num/mid {
			low = mid + 1
		} else if mid > num/mid {
			high = mid - 1
		} else if mid*mid == num { // 因为num / mid会取整，所以当mid == num / mid时不一定是真相等
			return true // 真相等
		} else {
			return false // 假相等
		}
	}
	// 没出现过相等的情况，一定不是完全平方数
	return false
}
