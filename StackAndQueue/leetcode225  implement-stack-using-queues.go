package StackAndQueue

type MyStack struct {
	Q1 []int
	Q2 []int
}

func Constructor1() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	// 在不为空的队列中执行入队操作，即为入栈
	// 注意：两个队列都为空时规定在Q1中入队
	if len(this.Q2) > 0 {
		this.Q2 = append(this.Q2, x)
	} else {
		this.Q1 = append(this.Q1, x)
	}
}

// Pop 并未对非法操作进行处理
func (this *MyStack) Pop() int {
	// 出栈需要将先入队的元素转移到另一队列中，只留下一个元素
	// 然后对该元素出队、即为出栈
	var val int
	if len(this.Q1) > 0 {
		for len(this.Q1) > 1 {
			v := this.Q1[0]
			// 转移元素
			this.Q1 = this.Q1[1:]
			this.Q2 = append(this.Q2, v)
		}
		val = this.Q1[0]
		this.Q1 = this.Q1[1:]
	} else {
		for len(this.Q2) > 1 {
			v := this.Q2[0]
			// 转移元素
			this.Q2 = this.Q2[1:]
			this.Q1 = append(this.Q1, v)
		}
		val = this.Q2[0]
		this.Q2 = this.Q2[1:]
	}
	return val
}

// Top 并未对非法操作进行处理
func (this *MyStack) Top() int {
	if len(this.Q1) > 0 {
		return this.Q1[len(this.Q1)-1]
	} else {
		return this.Q2[len(this.Q2)-1]
	}
}

func (this *MyStack) Empty() bool {
	return len(this.Q1) == 0 && len(this.Q2) == 0
}

// 单个队列也可实现栈，重点在于Pop()操作的实现。具体来说，当队列不空时，
// 将队头的n-1个元素（假设此时队列的长度为n）出队，然后再入队。
// 这样原先队尾的元素抵达队头，再出队，实现了栈的弹出（因为该元素就是对应栈的栈顶元素）
// 除Pop()操作外，其余的操作与双队列的实现方式基本一致
