package StringFunc

func reverseStr(s string, k int) string {
	res := []byte(s)
	// 每2k个数进行一次循环，先确定有多少个循环
	loop := len(s) / (2 * k) // 会丢失最后不足2k个数的循环
	if len(s)%(2*k) != 0 {
		loop++ // 补充不足2k个数的循环
	}
	for i := 0; i < loop; i++ {
		end := i*2*k + k - 1 // 记录每2k个数中的反转边界
		if i == loop-1 && len(s)-1 < end {
			end = len(s) - 1 // 考虑最后一轮循环数字不足k个
		}
		// 反转
		for j := 0; j < (end-i*2*k+1)/2; j++ {
			res[i*2*k+j], res[end-j] = res[end-j], res[i*2*k+j]
		}
	}
	return string(res)
}

// 更简洁的写法
func reverseStr1(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(ss); i += 2 * k {
		low := i
		high := i + k - 1
		if high > len(ss)-1 { // 防止最后不满k个
			high = len(ss) - 1
		}
		for low < high {
			ss[low], ss[high] = ss[high], ss[low]
			low++
			high--
		}
	}
	return string(ss)
}
