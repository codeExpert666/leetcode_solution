package StackAndQueue

// 用栈匹配括号：遇到左括号入栈，遇到右括号尝试与栈顶左括号匹配
func isValid(s string) bool {
	var stack []byte
	for _, bracket := range s {
		if bracket == '(' || bracket == '[' || bracket == '{' {
			// 左括号直接入栈
			stack = append(stack, byte(bracket))
		} else if len(stack) > 0 && byte(bracket)-stack[len(stack)-1] <= 2 {
			// 考虑到右括号的ascii码值与其对应的左括号差距小于等于2
			// 故采用简便写法，若字符串中包含除括号外的字符，该写法失效
			stack = stack[:len(stack)-1] // 左括号出栈
		} else {
			// 栈空或栈顶括号未能与右括号匹配
			return false
		}
	}
	// 存在一种可能，遍历字符串匹配完左右括号后，仍有剩余左括号
	return len(stack) == 0
}
