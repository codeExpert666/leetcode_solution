package ArrayList

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// 分别统计链表AB的长度
	countA, countB := 0, 0
	curA, curB := headA, headB
	for ; curA != nil; curA = curA.Next {
		countA++
	}
	for ; curB != nil; curB = curB.Next {
		countB++
	}
	// 长的链表先遍历，以保证链表尾部对齐
	curA, curB = headA, headB
	for countA > countB {
		curA = curA.Next
		countA--
	}
	for countB > countA {
		curB = curB.Next
		countB--
	}
	// 同时遍历，直到遇到相同指针
	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}
	return curA
}
