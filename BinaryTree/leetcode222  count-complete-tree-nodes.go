package BinaryTree

import "container/list"

/**
 * 方法一：递归，一棵完全二叉树的左右子树必定也是完全二叉树
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return countNodes(root.Left) + countNodes(root.Right) + 1
	}
}

/**
 * 方法二：迭代，以层序遍历为模板，当发现有结点的左右孩子不完整时，直接计算出最终结果
 */
func countNodes2(root *TreeNode) int {
	var res int

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(*TreeNode)
		res++
		// 左右孩子不完整时，由于完全二叉树的性质，队列中的元素就是所有还未统计的二叉树结点
		if cur.Left == nil {
			return res + queue.Len()
		}
		queue.PushBack(cur.Left)
		if cur.Right == nil {
			return res + queue.Len()
		}
		queue.PushBack(cur.Right)
	}
	return res
}

/**
 * 方法三：另一种更优秀的递归。原理：1、满二叉树的结点个数可以通过深度直接计算
 * 2、完全二叉树不断递归过程中一定会遇到满二叉树的子树，计算该子树的结点数量，向上层返回
 * 此方法的时间复杂度为O(logN * logN), 空间复杂度为O(logN)
 */
func countNodes3(root *TreeNode) int {
	if root == nil { // 空树特殊处理
		return 0
	}
	// 判断该树是否为满二叉树
	var leftDepth, rightDepth int
	left := root.Left
	right := root.Right
	// 获取左子树深度
	for left != nil {
		left = left.Left
		leftDepth++
	}
	// 获取右子树深度（并不是真正意义上的深度，主要用来判断root树是否为满二叉树）
	for right != nil {
		right = right.Right
		rightDepth++
	}
	if leftDepth == rightDepth { // 是满二叉树
		// 返回结点个数
		return (2 << leftDepth) - 1
	}
	// 不是满二叉树
	return countNodes3(root.Left) + countNodes3(root.Right) + 1
}
