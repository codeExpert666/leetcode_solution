package ArrayList

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	// 第一遍遍历，获取链表长度
	var length int
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	if length == 1 { // 处理特殊情况
		return true
	}
	// 第二遍遍历，将前半段链表反转
	pre, cur, count := head, head.Next, 1
	for count < length>>1 {
		tmp := cur.Next
		cur.Next = pre
		pre, cur = cur, tmp
		count++
	}
	// 此时pre指向反转后的前半段起始点，cur指向后半段的起始点（length为奇数时cur需右移一次）
	if length%2 == 1 {
		cur = cur.Next
	}
	// 前半段反向遍历，后半段正向遍历，判断是否相同
	for cur != nil {
		if pre.Val != cur.Val {
			return false
		}
		pre, cur = pre.Next, cur.Next
	}
	return true
}

// 两点值得优化：1、第一次遍历与第二次遍历可以合并，这两次遍历的目的是将链表拆分成前后两部分，
// 那么完全可以通过快慢双指针的方式拆分，快指针一次移动两个单位，慢指针移动一个单位，这样慢指针能到达中间位置
// 2、反转链表时，count从0开始即可，此时可令pre为nil，cur为head，这样就不用对头节点特殊处理，可统一反转步骤
func IsPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//慢指针，找到链表中间分位置，作为分割
	slow := head
	fast := head
	//记录慢指针的前一个节点，用来分割链表
	pre := head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	//分割链表
	pre.Next = nil
	//前半部分
	cur1 := head
	//反转后半部分，总链表长度如果是奇数，cur2比cur1多一个节点
	cur2 := ReverseList(slow)

	//开始两个链表的比较
	for cur1 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return true
}
