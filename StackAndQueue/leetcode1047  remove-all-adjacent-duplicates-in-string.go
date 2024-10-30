package StackAndQueue

func removeDuplicates(s string) string {
	var stack []byte
	for _, c := range s {
		if len(stack) > 0 && byte(c) == stack[len(stack)-1] {
			// 遍历元素与栈顶元素相同，表明为重复项
			stack = stack[:len(stack)-1]
		} else {
			// 非重复项则加入
			stack = append(stack, byte(c))
		}
	}
	return string(stack)
}
