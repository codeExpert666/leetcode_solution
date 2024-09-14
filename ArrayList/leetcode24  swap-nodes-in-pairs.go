package ArrayList

// 迭代法，不适用虚拟头结点（实际上使用虚拟头结点更方便），本题递归法也很方便。
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 处理空链表和只有一个结点的链表
		return head
	}
	// 处理第一组的两个结点
	cur := head.Next
	head.Next = cur.Next
	cur.Next = head
	head = cur
	// 两两处理后续结点
	pre := head.Next // 记录前一组的最后一个结点
	for cur = pre.Next; cur != nil && cur.Next != nil; cur = cur.Next {
		// 交换当前组的两个结点
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = cur
		pre.Next = tmp
		// 进入下一组的两个结点
		pre = cur
	}
	return head
}
