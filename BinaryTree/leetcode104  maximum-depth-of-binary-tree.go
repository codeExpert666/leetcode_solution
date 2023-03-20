package BinaryTree

import "container/list"

/**
 * 方法一：递归法，思路较简单
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

/**
 * 方法二：层序遍历模板，问题等价于求二叉树的层数
 */
func maxDepth2(root *TreeNode) int {
	var res int

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		res++
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
	}
	return res
}
