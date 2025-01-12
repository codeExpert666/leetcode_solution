package weeklycompetition

import "strings"

func hasMatch(s string, p string) bool {
	if p == "*" { // 特殊情况
		return true
	}

	// 以 * 为基线将 p 分为前后缀
	pos := strings.Index(p, "*")
	pre, suf := p[:pos], p[pos+1:]

	// s 必须包含 pre 与 suf，且 pre 在前
	// 先处理前后缀为空的情况
	if len(pre) == 0 {
		return strings.Contains(s, suf)
	}
	if len(suf) == 0 {
		return strings.Contains(s, pre)
	}

	// 再处理前后缀均不为空的情况
	preFirstPos := strings.Index(s, pre)
	if preFirstPos == -1 {
		return false
	}
	sufLastPos := strings.LastIndex(s, suf)
	if sufLastPos == -1 {
		return false
	}

	return preFirstPos+len(pre) <= sufLastPos
}
