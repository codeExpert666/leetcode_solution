package BinaryTree

import "container/list"

/**
 * 方法一：递归法，本题的关键点是判断根节点的左右子树是否是对称的，对称只有两种情况：
 * 1、左右子树同时为空；2、左右子树均不为空，左右子树根节点值相等，且左右子树的左右子树分别对称
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return compare(root.Left, root.Right)
	}
}

func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	} else {
		return compare(left.Left, right.Right) && compare(left.Right, right.Left)
	}
}

/**
 * 方法二：迭代法，通过队列来判断根节点的左子树和右子树的内侧和外侧是否相等
 */
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 队列初始化，左右孩子判空在后面处理
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)

	// 类似与层序遍历的操作，但与层序遍历的出入队数量与顺序不同
	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)  // 左子树待比较结点
		right := queue.Remove(queue.Front()).(*TreeNode) // 右子树待比较结点
		// 判空在这里进行
		if left == nil && right == nil {
			continue
		} else if left != nil && right != nil {
			// 当左右待比较结点均不为空
			if left.Val != right.Val {
				return false
			}
			// 入队，注意入队顺序
			queue.PushBack(left.Left)
			queue.PushBack(right.Right)
			queue.PushBack(left.Right)
			queue.PushBack(right.Left)
		} else {
			return false
		}
	}
	return true
}

/**
 * 方法三：思路与方法二一致，用栈也可以达到同样效果
 */
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 栈初始化，左右孩子判空在后面处理
	stack := list.New()
	stack.PushBack(root.Right)
	stack.PushBack(root.Left)

	for stack.Len() > 0 {
		left := stack.Remove(stack.Back()).(*TreeNode)  // 左子树待比较结点
		right := stack.Remove(stack.Back()).(*TreeNode) // 右子树待比较结点
		// 判空在这里进行
		if left == nil && right == nil {
			continue
		} else if left != nil && right != nil {
			// 当左右待比较结点均不为空
			if left.Val != right.Val {
				return false
			}
			// 入栈，注意入栈顺序
			stack.PushBack(right.Left)
			stack.PushBack(left.Right)
			stack.PushBack(right.Right)
			stack.PushBack(left.Left)
		} else {
			return false
		}
	}
	return true
}
