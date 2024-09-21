package StackAndQueue

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, s := range tokens {
		switch s {
		case "+", "-", "*", "/": // 操作符
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch s {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		default: // 操作数
			num, _ := strconv.Atoi(s)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
