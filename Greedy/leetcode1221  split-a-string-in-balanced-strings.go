package Greedy

// 利用栈匹配L与R，当栈为空时，表明切割出一个平衡子串。
func balancedStringSplit(s string) int {
	var res int
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == stack[len(stack)-1] {
			// 栈为空或栈顶元素与当前元素相同，直接入栈
			stack = append(stack, s[i])
		} else {
			// 栈顶元素不同，则弹出栈顶元素，表明一对LR匹配
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 出栈导致栈为空，表明分割出一个平衡子串
				res++
			}
		}
	}
	return res
}

// 优化思路：可以用一个int变量替代栈，具体方法是，遇到L就减1，遇到R就加1。当该变量为0时，说明分割出一个平衡子串
