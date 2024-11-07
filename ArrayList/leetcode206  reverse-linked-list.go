package ArrayList

// 双指针遍历法
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 单独处理不包含元素和只包含一个元素的链表
		return head
	}
	pre, cur := head, head.Next
	for cur != nil {
		tmp := cur.Next // 提前保存后续元素的开始指针，不然会被覆盖
		cur.Next = pre  // 更新Next指针
		// 更新循环用到的指针
		pre = cur
		cur = tmp
	}
	head.Next = nil // head指针变为尾指针，Next需赋空值
	return pre      // 最后pre指向头结点
}

// 双指针遍历精简版
func ReverseList1(head *ListNode) *ListNode {
	var pre *ListNode // 通过将pre的初始值设置为nil，实现遍历过程的统一
	cur := head
	for cur != nil {
		tmp := cur.Next // 提前保存后续元素的开始指针，不然会被覆盖
		cur.Next = pre  // 更新Next指针
		// 更新循环用到的指针
		pre = cur
		cur = tmp
	}
	return pre // 最后pre指向头结点
}

// 递归
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	subList := ReverseList2(head.Next) // 反转后续子串
	head.Next.Next = head              // 将原头结点添加到子串末尾
	head.Next = nil                    // 原头结点变成末尾结点，其Next应为空
	return subList
}
