package BinaryTree

import "container/list"

/**
 * 方法一：递归，思路比较简单，相同树的左右子树也应相同
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

/**
 * 方法二：基于队列的迭代，利用层序遍历模板
 */
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	queue := list.New()
	queue.PushBack(p)
	queue.PushBack(q)

	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)
		if left == nil && right == nil {
			continue
		} else if left != nil && right != nil {
			if left.Val != right.Val {
				return false
			}
			// 注意入队顺序
			queue.PushBack(left.Left)
			queue.PushBack(right.Left)
			queue.PushBack(left.Right)
			queue.PushBack(right.Right)
		} else {
			return false
		}
	}
	return true
}
