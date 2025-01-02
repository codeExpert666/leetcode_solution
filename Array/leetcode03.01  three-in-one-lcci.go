package Array

type TripleInOne struct {
	stack []int // 三合一栈
	size  int   // 每个栈的容量
	top   []int // 每个栈的栈顶指针
}

func ConstructorTIO(stackSize int) TripleInOne {
	return TripleInOne{
		stack: make([]int, stackSize*3),
		size:  stackSize,
		top:   []int{0, stackSize, 2 * stackSize},
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.top[stackNum] < this.size*(stackNum+1) {
		this.stack[this.top[stackNum]] = value
		this.top[stackNum]++
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	if !this.IsEmpty(stackNum) {
		this.top[stackNum]--
		return this.stack[this.top[stackNum]]
	}
	return -1
}

func (this *TripleInOne) Peek(stackNum int) int {
	if !this.IsEmpty(stackNum) {
		return this.stack[this.top[stackNum]-1]
	}
	return -1
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.top[stackNum] == this.size*stackNum
}
