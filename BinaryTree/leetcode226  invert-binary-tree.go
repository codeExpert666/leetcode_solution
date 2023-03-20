package BinaryTree

import "container/list"

/**
 * 方法一：递归解法：思路比较简单，先反转左右子树，再交换左右孩子结点指针
 */
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		invertTree(root.Left)
		invertTree(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
	return root
}

/**
 * 方法二：后序遍历模板，将原先的访问操作改为左右孩子互换的操作
 */
func invertTree2(root *TreeNode) *TreeNode {
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
				cur.Left, cur.Right = cur.Right, cur.Left
				stack.Remove(stack.Back())
				pre = cur
				cur = nil
			} else {
				cur = cur.Right
			}
		}
	}
	return root
}
