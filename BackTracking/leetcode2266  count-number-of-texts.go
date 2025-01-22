package BackTracking

const mod = 1000000007
const mx = 100001

var f = [mx]int{1, 1, 2, 4}
var g = f

func init() {
	for i := 4; i < mx; i++ {
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
		g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	}
}

// 分组动态规划+乘法原理，需要注意4个'2'的情况并不会解析为字符'a'
// 连续解析的字符个数最多只能为3个或4个
// 以下代码来自灵神题解
func countTexts(pressedKeys string) int {
	n := len(pressedKeys)
	res, count := 1, 0

	for i := 0; i < n; i++ {
		count++
		if i == n-1 || pressedKeys[i] != pressedKeys[i+1] {
			if pressedKeys[i] == '7' || pressedKeys[i] == '9' {
				res = (res * g[count]) % mod
			} else {
				res = (res * f[count]) % mod
			}
			count = 0
		}
	}

	return res
}
