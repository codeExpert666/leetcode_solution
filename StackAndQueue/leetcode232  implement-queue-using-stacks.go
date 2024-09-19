package StackAndQueue

type MyQueue struct {
	InputStack  []int // 模拟队列的入队
	OutputStack []int // 模拟队列的出队
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.InputStack = append(this.InputStack, x)
}

// Pop 这里并未对非法操作进行处理，例如对空队列进行出队操作
func (this *MyQueue) Pop() int {
	if len(this.OutputStack) == 0 {
		//模拟出队的栈为空，需要将模拟入队的栈中元素转移过来
		for len(this.InputStack) > 0 {
			v := this.InputStack[len(this.InputStack)-1]
			this.InputStack = this.InputStack[:len(this.InputStack)-1] // 出栈
			this.OutputStack = append(this.OutputStack, v)             // 入栈
		}
	}
	// 模拟出队的栈不为空，出栈即为出队
	v := this.OutputStack[len(this.OutputStack)-1]
	this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
	return v
}

// Peek 这里并未对非法操作进行处理，例如查看空队列的队头元素操作
func (this *MyQueue) Peek() int {
	if len(this.OutputStack) == 0 {
		//模拟出队的栈为空，需要将模拟入队的栈中元素转移过来
		for len(this.InputStack) > 0 {
			v := this.InputStack[len(this.InputStack)-1]
			this.InputStack = this.InputStack[:len(this.InputStack)-1] // 出栈
			this.OutputStack = append(this.OutputStack, v)             // 入栈
		}
	}
	return this.OutputStack[len(this.OutputStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.InputStack) == 0 && len(this.OutputStack) == 0
}
