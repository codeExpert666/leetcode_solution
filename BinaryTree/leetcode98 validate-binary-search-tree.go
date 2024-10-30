package BinaryTree

import "container/list"

/**
 * 递归法一，一棵树是二叉搜索树等价于其左右子树都是二叉搜索树，并且左子树的最右结点小于根节点，右子树的的最左结点大于根节点
 */
func IsValidBST(root *TreeNode) bool {
	if root == nil {
		// 空树是二叉搜索树
		return true
	}
	if !IsValidBST(root.Left) || !IsValidBST(root.Right) {
		// 当左右子树有非搜索二叉树时，整个二叉树也是非搜索二叉树
		return false
	}
	// 寻找左右子树的最右最左结点
	left, right := root.Left, root.Right
	if left != nil { // 寻找左子树最右结点
		for left.Right != nil {
			left = left.Right
		}
		// 与根节点比较
		if root.Val <= left.Val {
			return false
		}
	}
	if right != nil { // 寻找右子树最左结点
		for right.Left != nil {
			right = right.Left
		}
		// 与根结点比较
		if root.Val >= right.Val {
			return false
		}
	}
	// 所有条件均满足，则该树是二叉搜索树
	return true
}

/**
 * 递归法二：利用中序遍历递归思想，若二叉树是搜索树，则遍历到的值应不断增大。这种递归法比较简洁
 * 以后再遇到需要利用几个递归思路的题目时，可以之间套用递归模板，在此基础上进行修改，这样的代码会比较简洁
 */
var pre *TreeNode // 记录前一个结点，因为要不断与前一结点进行比较
func IsValidBST1(root *TreeNode) bool {
	if root == nil {
		// 空树是二叉搜索树
		return true
	}
	left := IsValidBST1(root.Left)
	if pre != nil && root.Val <= pre.Val {
		// 当前结点与前一结点比较
		return false
	}
	pre = root // 更新前一结点
	right := IsValidBST1(root.Right)

	return left && right // 左右子树都是二叉搜索树
}

/**
 * 迭代法，也是中序遍历思想，利用栈来实现
 */
func IsValidBST2(root *TreeNode) bool {
	stack := list.New()
	cur := root
	var pre *TreeNode // // 记录前一个结点，因为要不断与前一结点进行比较

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			// 与前一结点比较
			if pre != nil && cur.Val <= pre.Val {
				return false
			}
			pre = cur
			cur = cur.Right
		}
	}
	return true
}
