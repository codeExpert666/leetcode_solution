package StringFunc

// 先整体反转，然后遍历，边处理空格，边识别并反转每个单词
func reverseWords(s string) string {
	ss := []byte(s)
	reverseString(ss)
	// 去除空格并反转单词
	res := make([]byte, 0, len(ss)) // 存储结果串
	var start, end int              // 记录每个单词的起始结束位置
	for i, c := range ss {
		if c == ' ' {
			if i > 0 && ss[i-1] != ' ' { // 一个单词的末尾
				end = i - 1
				// 已经找到一个单词，进行反转
				reverseString(ss[start : end+1])
				res = append(res, ss[start:end+1]...)
				res = append(res, ' ') // 一个单词结束后添加空格
			}
			if i < len(ss)-1 && ss[i+1] != ' ' { // 一个单词的开始
				start = i + 1
			}
		}
	}
	if end == 0 || end < start { // 存在一种情况，ss不存在尾随空格，最后一个字符串无法录入
		// 具体来说：1、只有一个单词，且没有尾随  2、大于1个单词，且没有尾随
		reverseString(ss[start:])
		res = append(res, ss[start:]...)
	} else { // 若ss存在尾随空格，会导致res中多出一个尾随空格
		res = res[:len(res)-1]
	}
	return string(res)
}

// 思路一致，但无需额外创建数组，直接在原数组上原地操作，节省空间
func reverseWords1(s string) string {
	ss := []byte(s)
	// 去除字符串多余空格（pre与cur指针实现原地删除的方式需要好好体会）
	pre := 0 // 待插入位置
	for cur := 0; cur < len(ss); cur++ {
		if ss[cur] != ' ' || (ss[cur] == ' ' && cur-1 >= 0 && ss[cur-1] != ' ') {
			// 两种情况需要保留：一种是遇到非空格，一种是遇到空格，但前一字符是非空格
			ss[pre] = ss[cur]
			pre++
		}
	}
	if ss[pre-1] == ' ' { // 当字符串有尾随空格时，上述做法得到的字符串的末尾会多一个空格
		pre--
	}
	// 反转得到的字符串
	reverseString(ss[:pre])
	// 再反转每个单词
	wordStart := 0
	for i := 0; i < pre; i++ { // 反转除最后一个单词的所有单词
		if ss[i] == ' ' {
			reverseString(ss[wordStart:i])
			wordStart = i + 1
		}
	}
	reverseString(ss[wordStart:pre]) // 反转最后一个单词
	return string(ss[:pre])
}
