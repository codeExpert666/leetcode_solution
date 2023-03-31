package BackTracking

/**
 * 回溯
 */
var (
	res5  [][]string
	path5 []string
)

func partition(s string) [][]string {
	res5, path5 = make([][]string, 0), make([]string, 0)
	backtracking4(s, 0)
	return res5
}

func backtracking4(s string, startIndex int) {
	// 终止条件
	if startIndex == len(s) {
		tmp := make([]string, len(path5))
		copy(tmp, path5)
		res5 = append(res5, tmp)
		return
	}
	for i := startIndex; i < len(s); i++ {
		if s[i] == s[startIndex] && isPalindrome(s[startIndex:i+1]) { // 剪枝
			path5 = append(path5, s[startIndex:i+1])
			backtracking4(s, i+1)
			path5 = path5[:len(path5)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
		if s[low] != s[high] {
			return false
		}
	}
	return true
}
