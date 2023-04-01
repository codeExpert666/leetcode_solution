package BackTracking

import "strconv"

/**
 * 回溯，注意一个坑：这里的给出的串不一定具有有效的ip地址
 */
var (
	res6  []string
	path6 []string
)

func restoreIpAddresses(s string) []string {
	res6, path6 = make([]string, 0), make([]string, 0, 4)
	backtracking5(s, 0)
	return res6
}

func backtracking5(s string, startIndex int) {
	// 终止条件
	if startIndex == len(s) {
		if len(path6) == 4 {
			tmp := getString(path6)
			res6 = append(res6, tmp)
		}
		return
	}
	for i := startIndex; i < startIndex+3 && i < len(s); i++ { // ip地址每一部分最多只有三位
		// 待处理串
		part := s[startIndex : i+1]
		partNum, _ := strconv.Atoi(part)
		// ！！！ip地址有效性判断
		if partNum > 255 {
			break
		}
		path6 = append(path6, part)
		backtracking5(s, i+1)
		path6 = path6[:len(path6)-1]
		if len(part) == 1 && partNum == 0 { // 剪枝，若首位遇到0，则该部分只能是0
			break
		}
		if len(part) == 2 && partNum > 25 { // 剪枝，若这一部分前两位已经大于25，则该部分不可能有三位
			break
		}
	}
}

func getString(path []string) string {
	var res string
	for _, v := range path {
		res += v + "."
	}
	return res[:len(res)-1]
}
