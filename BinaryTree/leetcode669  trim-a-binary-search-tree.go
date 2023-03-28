package BinaryTree

/**
 * 递归法，比较简单
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

/**
 * 迭代法，稍微复杂些，主要分为三步：1、寻找结点值在给定区间的根节点
 * 2、修剪该根节点的左子树 3、修剪该根节点的右子树
 */
func trimBST1(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 处理 root，让 root 移动到[low, high] 范围内，注意是左闭右闭
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	cur := root
	// 此时 root 已经在[low, high] 范围内，处理左孩子元素小于 low 的情况（左节点是一定小于 root.Val，因此天然小于 high）
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	cur = root
	// 此时 root 已经在[low, high] 范围内，处理右孩子大于 high 的情况
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}
