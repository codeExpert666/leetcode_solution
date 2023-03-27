package BinaryTree

/**
 * 递归法，注意本题的前提，p和q一定存在于二叉树中，如果不存在，这种方法就不适用
 * 仔细体会本题的递归思想，值得学习
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// check，这一部分可以理解为先序遍历，目的是及时返回
	if root == nil {
		return root
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// Conquer，这一部分可以理解为后序遍历，目的是回溯
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	// 包含了p，q其中一个是另一个的祖先的情况
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	// 左右都为空，说明p，q不在root的子树中
	return nil
}
