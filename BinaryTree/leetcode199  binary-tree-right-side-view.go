package BinaryTree

import "container/list"

/**
 * 层序遍历模板，只记录每层最后一个结点
 */
func rightSideView(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	queue := list.New() // 队列，利用go自带链表
	queue.PushBack(root)
	for queue.Len() > 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
			if i == l-1 {
				// 本层最后一个结点
				res = append(res, cur.Val)
			}
		}
	}
	return res
}
