package BinaryTree

import "container/list"

/**
 * 递归法：一棵平衡二叉树的左右子树也一定是平衡二叉树，且左右子树的高度差小于1
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		flag := isBalanced(root.Left) && isBalanced(root.Right)
		if flag {
			d := maxDepth(root.Left) - maxDepth(root.Right)
			return d <= 1 && d >= -1
		}
		return false
	}
}

/**
 * 迭代法，利用后序遍历模板，访问结点操作改为判断左右子树高度差
 */
func isBalanced1(root *TreeNode) bool {
	stack := list.New()
	cur := root
	var pre *TreeNode

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Back().Value.(*TreeNode)
			if cur.Right == nil || cur.Right == pre {
				// 判断左右子树高度差
				d := maxDepth(cur.Left) - maxDepth(cur.Right)
				if d > 1 || d < -1 {
					return false
				}
				stack.Remove(stack.Back())
				pre = cur
				cur = nil
			} else {
				cur = cur.Right
			}
		}
	}
	return true
}

/**
 * 递归法：另一种更简单的递归法，在获取深度的递归方法基础上加上平衡的判断，当有子树不平衡时，则整棵树不平衡
 * 这种递归比第一种更好，因为全程只牵涉到一个函数的递归
 */
func isBalanced2(root *TreeNode) bool {
	h := getHeight(root)
	if h == -1 {
		return false
	}
	return true
}

// 返回以该节点为根节点的二叉树的高度，如果不是平衡二叉树了则返回-1
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := getHeight(root.Left), getHeight(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	return max(l, r) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
