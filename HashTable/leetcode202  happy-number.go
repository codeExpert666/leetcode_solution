package HashTable

func isHappy(n int) bool {
	m := make(map[int]struct{}) // 哈希表用于记录演变过程中的数
	for {
		// 计算当前数各个位的平方和
		sum := 0
		for n != 0 {
			tmp := n % 10
			sum += tmp * tmp
			n /= 10
		}
		// 判断是否为快乐数
		if _, ok := m[sum]; ok { // hash表存在则会无限循环，不是快乐数
			return false
		} else if sum == 1 { // 快乐数定义
			return true
		} else {
			m[sum] = struct{}{} // 在hash表中存储
			n = sum             // 下一个循环数
		}
	}
}
