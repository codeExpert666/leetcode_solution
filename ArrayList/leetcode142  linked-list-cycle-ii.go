package ArrayList

// 检测第一个被二次遍历的结点，即为入环结点
// hash表法，空间复杂度高
func detectCycle(head *ListNode) *ListNode {
	// 定义hash表用于存储遍历过的结点
	pointerMap := make(map[*ListNode]struct{})
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := pointerMap[cur]; ok {
			// hash表中已存在该结点，说明该结点是入环结点
			return cur
		} else {
			// 不存在则添加以记录
			pointerMap[cur] = struct{}{}
		}
	}
	// 链表遍历出现空结点，说明不存在环
	return nil
}

// 快慢指针法，先通过快慢指针的相遇判断是否有环，再根据相遇结点与头结点到入环结点的距离相等寻找入环结点
// 依赖一定的数学推理
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			tmp := head
			for tmp != fast {
				tmp = tmp.Next
				fast = fast.Next
			}
			return fast
		}
	}
	return nil
}
