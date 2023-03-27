package BinaryTree

/**
 * 递归法，注意本题的前提，p和q一定存在于二叉树中，如果不存在，这种方法就不适用
 * 由于本题的二叉树是二叉搜索树，所以可以根据结点值来判断最近公共祖先的位置
 * 这种递归的写法考虑了p，q不在二叉树中的情况
 */
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 若p和q的值都小于root的值，则最近公共祖先在左子树中
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor1(root.Left, p, q)
	}
	// 若p和q的值都大于root的值，则最近公共祖先在右子树中
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor1(root.Right, p, q)
	}
	// 若 1、p和q的值分别大于和小于root的值；2、p的值或q的值与root相等，则最近公共祖先为root
	return root
}

/**
 * 迭代法，思路与递归法一致
 * 这种迭代的写法也考虑了p，q不在二叉树中的情况
 */
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
