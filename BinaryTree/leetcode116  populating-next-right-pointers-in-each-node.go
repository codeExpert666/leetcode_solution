package BinaryTree

import "container/list"

/**
 * 层序遍历模板，实际上这种方法不符合题目要求的O(1)空间复杂度，方法二的递归符合
 */
func connect1(root *Node) *Node {
	// 特殊处理空树
	if root == nil {
		return root
	}

	// 初始化队列
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*Node)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
			if i < l-1 {
				// cur在本层有下一结点
				cur.Next = queue.Front().Value.(*Node)
			}
		}
	}
	return root
}

/**
 * 递归解法，思路就是对于树中的每个结点，都考虑两件事：1、连接本结点的左右孩子。2、将本结点的右孩子与兄弟结点的左孩子连接
 */
func connect2(root *Node) *Node {
	if root != nil && root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		connect2(root.Left)
		connect2(root.Right)
	}
	return root
}
