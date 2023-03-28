package BinaryTree

/**
 * 迭代法：待插入值比根节点值小，则插入左子树；反之，则插入右子树，直到找到空结点，即为插入位置
 */
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	// 待插入值最终一定成为搜索二叉树的叶子结点，故左右指针不用赋值
	insertNode := &TreeNode{
		Val: val,
	}
	cur := root // 以防根节点指针丢失
	// 迭代过程
	for cur != nil {
		if val < cur.Val {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				// 插入
				cur.Left = insertNode
				break
			}
		} else {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				// 插入
				cur.Right = insertNode
				break
			}
		}
	}
	// 单独处理根节点为空的情况
	if cur == nil {
		root = insertNode
	}
	return root
}

/**
 * 递归法：思路与迭代法一致
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
