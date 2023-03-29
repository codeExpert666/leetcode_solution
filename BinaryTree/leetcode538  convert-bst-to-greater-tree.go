package BinaryTree

import "container/list"

/**
 * 递归法：比某一结点大的值都在中序遍历结果的右方，故采用逆中序遍历（右中左），不断计算累加值
 */
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Right) // 注意先遍历右子树
		sum += node.Val
		node.Val = sum
		travel(node.Left)
	}
	travel(root)
	return root
}

/**
 * 迭代法，与递归思路一致
 */
func convertBST1(root *TreeNode) *TreeNode {
	var sum int
	stack := list.New()
	cur := root

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Right
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			sum += cur.Val
			cur.Val = sum
			cur = cur.Left
		}
	}
	return root
}
