package ArrayList

// 链表版单调栈（递减）
func removeNodes(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0) // 递减单调栈
	for tmp := head; tmp != nil; tmp = tmp.Next {
		for len(stack) > 0 && tmp.Val > stack[len(stack)-1].Val { // 当前元素大于栈顶元素
			stack = stack[:len(stack)-1] // 出栈
		}
		if len(stack) == 0 { // 若是因为栈空退出循环，则该元素是新链表的头节点
			head = tmp
		} else { // 因为不大于栈顶元素退出循环，则该元素是栈顶元素在新链表中的后续
			stack[len(stack)-1].Next = tmp
		}
		stack = append(stack, tmp) // 入栈
	}
	return head
}
