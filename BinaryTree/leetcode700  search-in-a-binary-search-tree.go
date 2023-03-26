package BinaryTree

/**
 * 本题的迭代法与递归法思路完全一致：小于根节点值，搜索左子树，大于根节点值，搜索右子树
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val > val {
			// 由于是二叉搜索树，所以如果当前节点的值大于目标值，那么目标值一定在当前节点的左子树中
			root = root.Left
		} else {
			// 由于是二叉搜索树，所以如果当前节点的值小于目标值，那么目标值一定在当前节点的右子树中
			root = root.Right
		}
	}
	// 两种情况：1、root == nil，说明没有找到目标值，返回nil
	// 2、root.Val == val，说明找到了目标值，返回root
	return root
}

/**
 * 递归法
 */
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val { // 注意顺序
		return root
	}
	var res *TreeNode
	if root.Val > val {
		res = searchBST1(root.Left, val)
	} else {
		res = searchBST1(root.Right, val)
	}
	return res
}
