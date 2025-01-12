package Simulation

// 按照数位从高到低遍历模拟
func GenerateKey(num1 int, num2 int, num3 int) int {
	var res int
	base := []int{1000, 100, 10, 1}

	for _, b := range base {
		n1, n2, n3 := num1/b, num2/b, num3/b
		res += b * min(n1, n2, n3)
		num1, num2, num3 = num1%b, num2%b, num3%b
	}

	return res
}

// 灵神解答。按照数位从低到高遍历模拟
func GenerateKey1(x, y, z int) (ans int) {
	for pow10 := 1; x > 0 && y > 0 && z > 0; pow10 *= 10 {
		ans += min(x%10, y%10, z%10) * pow10
		x /= 10
		y /= 10
		z /= 10
	}
	return
}
