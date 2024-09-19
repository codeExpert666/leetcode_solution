package String

import "strings"

// 方法一：KMP法，利用最长公共前后缀进行判断：对于一个字符串，若最长公共前后缀的长度小于总长度的一半，则一定不由重复子串构成
// 若等于，则一定是；若大于，则取决于前后缀起始位置之间的子串能否重复构成完整串
func repeatedSubstringPattern(s string) bool {
	if len(s) < 2 { // 单独处理特殊情况，以防len(s)=1时，next[len(s)-1] >= len(s)/2失效
		return false
	}
	next := getNext(s)
	if next[len(s)-1] >= len(s)/2 {
		// 有重复串的可能，且重复的模式串只可能是前后缀起始位置之间的子串
		// 从子串的二倍长位置处开始遍历，每增加一个子串长度，最长公共前后缀的长度也增加一个子串
		reStrLen := len(s) - next[len(s)-1] // 重复串长度
		if len(s)%reStrLen == 0 {
			for i := 2*reStrLen - 1; i < len(s); i += reStrLen {
				if next[i] != i-reStrLen+1 {
					return false
				}
			}
			return true
		}
	}
	return false
}

// 方法一的优化版本，可以证明：
// 一个字符串由重复子串构成等价于该字符串除最大公共前后缀的剩余部分长度能够整除总字符串的长度
// 该原则不适用于字符串没有公共前后缀的情况，也即next[len(s)-1] = 0，当然，
// 这种情况很容易看出s一定不由重复子串构成
func repeatedSubstringPattern1(s string) bool {
	next := getNext(s)
	return next[len(s)-1] != 0 && len(s)%(len(s)-next[len(s)-1]) == 0
}

// 方法二：将字符串s复制扩展一倍，s由重复子串构成等价于s+s内部包含s（不考虑一头一尾两个s）
func repeatedSubstringPattern2(s string) bool {
	ss := s + s
	return strings.Contains(ss[1:len(ss)-1], s)
}
