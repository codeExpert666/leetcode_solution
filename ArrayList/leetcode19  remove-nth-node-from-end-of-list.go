package ArrayList

// 固定大小滑动窗口，大小为n+1（这是删除，寻找是n）
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head} // 创建虚拟头结点
	// 寻找倒数第n+1个结点
	slow, fast := newHead, newHead
	count := 1
	for ; fast.Next != nil; fast = fast.Next {
		if count < n+1 { // 创建初始窗口
			count++
		} else { // 移动窗口
			slow = slow.Next
		}
	}
	// 此时slow指向倒数第n+1结点，利用该结点删除指定结点
	slow.Next = slow.Next.Next
	// 返回真正头结点
	return newHead.Next
}
