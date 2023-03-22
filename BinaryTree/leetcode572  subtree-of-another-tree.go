package BinaryTree

import "container/list"

/**
 * （Me）基于队列的迭代法，对大树做层序遍历，当遇到某一结点的值与小树的根节点值相同时，调用判断
 * 两树相等的方法letcode100
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil { // 处理大树为空树的情况
		return subRoot == nil
	}
	if subRoot == nil { // 处理小树为空的情况
		return false
	}

	queue := list.New() // 用于大树的层序遍历
	queue.PushBack(root)

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(*TreeNode)
		// 调用判断两树相等的方法
		if cur.Val == subRoot.Val && isSameTree(cur, subRoot) {
			return true
		}
		if cur.Left != nil {
			queue.PushBack(cur.Left)
		}
		if cur.Right != nil {
			queue.PushBack(cur.Right)
		}
	}
	return false
}

/**
 * 递归法，当大树与小树均不为空时，小树是大树的子树的充要条件是 1、大树与小树相同 或 2、小树是大树左子树的子树 或 3、小树是大树右子树的子树
 */
func isSubtree1(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root != nil && subRoot != nil {
		return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	} else {
		return false
	}
}
