package BackTracking

/**
 * 回溯
 */
var (
	res2  []string
	path2 string
	m     = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
)

func letterCombinations(digits string) []string {
	res2 = make([]string, 0)
	backtracking1(digits, 0)
	return res2
}

func backtracking1(digits string, pos int) {
	// 终止条件
	if pos == len(digits) {
		if len(digits) > 0 {
			// 将空串情况区分开
			res2 = append(res2, path2)
		}
		return
	}
	for _, c := range m[digits[pos]] {
		path2 += string(c)
		backtracking1(digits, pos+1)
		path2 = path2[:len(path2)-1] //回溯
	}
}
