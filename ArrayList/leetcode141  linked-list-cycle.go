package ArrayList

// 快慢指针，若两指针能够相遇，则有环
// 一个指针从相遇位置出发，一个从链表头部出发，两个指针的相遇位置即为环的起始位置
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 放在后面是排除slow,fast初始值相等的情况
			return true
		}
	}
	return false
}
