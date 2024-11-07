package ArrayList

// 由于该链表的实现是不带空节点的链表，故需要对头结点单独处理
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val { // 删除链表开头的所有符合条件的值
		head = head.Next
	}
	if head == nil { // 空链表无法删除
		return head
	}
	// 头结点不是待删除值时，才方便定义父节点
	pre, cur := head, head.Next
	for cur != nil {
		if cur.Val == val { // 删除当前节点
			pre.Next = cur.Next
		} else { // 更新父节点
			pre = cur
		}
		cur = cur.Next
	}
	return head
}
