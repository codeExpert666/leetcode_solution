package BinaryTree

import "container/list"

/**
 * 方法一：递归，若左右子树能找到和为sum值减根节点值的路径，则该树存在要求路径
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 空树遇到目标值，一定没有路径
	if root == nil {
		return false
	}
	// 根节点当前值等于目标值，不一定找到了路径，还得看该节点是否为叶子结点
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	// 左右有一个路径存在即可
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

/**
 * 方法二：迭代法，利用后序遍历模板，与leetcode257思路一致
 */
func hasPathSum1(root *TreeNode, targetSum int) bool {
	var sum int

	stack := list.New()
	cur := root
	var pre *TreeNode

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			sum += cur.Val // 更新此时的总和
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Back().Value.(*TreeNode)
			if cur.Right == nil || cur.Right == pre {
				if cur.Left == nil && cur.Right == nil && sum == targetSum {
					// 遍历到叶子结点,且此时和等于目标值
					return true
				}
				stack.Remove(stack.Back())
				sum -= cur.Val // 更新此时的总和
				pre = cur
				cur = nil
			} else {
				cur = cur.Right
			}
		}
	}
	return false
}
