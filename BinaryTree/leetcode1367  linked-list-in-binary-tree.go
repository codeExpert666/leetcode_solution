package BinaryTree

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归 + KMP
func IsSubPath(head *ListNode, root *TreeNode) bool {
	array, next := getNext(head)
	l := len(array)
	// 递归遍历树，利用 next 数组不断匹配
	// i 表示 next 数组当前待匹配位置，tNode 表示树当前待匹配节点
	var travel func(i int, tNode *TreeNode) bool
	travel = func(i int, tNode *TreeNode) bool {
		if i == l { // 完全匹配
			return true
		}
		if tNode == nil { // 未完全匹配但树的某分支已遍历完毕
			return false
		}
		j := i                     // 数组下一待匹配位置（当前位匹配的情况，不匹配后续修正）
		if tNode.Val != array[i] { // 当前位不匹配
			for j = next[i]; j >= 0 && array[j] != tNode.Val; {
				j = next[j]
			}
		}
		return travel(j+1, tNode.Left) || travel(j+1, tNode.Right)
	}
	return travel(0, root)
}

// 获取 head 链表对应的数组，以及数组对应的 next 数组
func getNext(head *ListNode) ([]int, []int) {
	// 将链表转换为数组
	array := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		array = append(array, cur.Val)
	}
	// 计算 next 数组
	l := len(array)
	next := make([]int, l)
	next[0] = -1 // next[1] 没必要初始化，下面循环会处理，同时如果 l==1，初始化还会导致错误
	for i := 1; i < l; i++ {
		j := next[i-1]
		for j >= 0 && array[j] != array[i-1] {
			j = next[j]
		}
		next[i] = j + 1
	}
	return array, next
}

// // 树中是否存在以根节点为起点的向下路径与 head 链表对应
// func isSubPathWithRoot(head *ListNode, root *TreeNode) bool {
// 	if (head != nil && root == nil) || (head == nil && root != nil) {
// 		return false
// 	}
// 	if head == nil && root == nil {
// 		return true
// 	}
// 	if head.Val != root.Val {
// 		return false
// 	}
// 	return isSubPathWithRoot(head.Next, root.Left) || isSubPathWithRoot(head.Next, root.Right)
// }
