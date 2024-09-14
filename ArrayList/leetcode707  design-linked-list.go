package ArrayList

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	for cur, i := this.Head, 0; cur != nil; cur, i = cur.Next, i+1 {
		if i == index {
			return cur.Val
		}
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &Node{
		Val:  val,
		Next: this.Head,
	}
	this.Head = newHead
}

func (this *MyLinkedList) AddAtTail(val int) {
	tail := &Node{
		Val:  val,
		Next: nil,
	}
	cur := this.Head
	if cur == nil {
		this.Head = tail
	} else {
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = tail
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else {
		for cur, i := this.Head, 0; cur != nil; cur, i = cur.Next, i+1 {
			if i == index-1 {
				insNode := &Node{
					Val:  val,
					Next: cur.Next,
				}
				cur.Next = insNode
				break
			}
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Head == nil {
		return
	}
	if index == 0 {
		this.Head = this.Head.Next
	} else {
		for cur, i := this.Head, 0; cur != nil && cur.Next != nil; cur, i = cur.Next, i+1 {
			if i == index-1 {
				cur.Next = cur.Next.Next
				break
			}
		}
	}
}
