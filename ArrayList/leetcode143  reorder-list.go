package ArrayList

func ReorderList(head *ListNode) {
	if head.Next == nil { // 特殊情况
		return
	}
	// 先利用快慢指针找到后半段链表的起点
	var slow, fast, pre *ListNode = head, head, nil
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil // 将前后半段截断
	// 此时slow指向后半段起点，对后半段进行反转
	pre = nil
	for slow != nil {
		tmp := slow.Next
		slow.Next = pre
		pre, slow = slow, tmp
	}
	// 此时pre指向反转后的后半段链表，与head指向的前半段链表交叉
	// 前半段链表更短，以前半段链表为循环主体
	h1, h2 := head, pre
	for h1 != nil {
		tmp1, tmp2 := h1.Next, h2.Next
		h1.Next, h1 = h2, tmp1
		if h1 != nil { // 当h1结束时，h2无需进行处理
			h2.Next, h2 = h1, tmp2
		}
	}
}
