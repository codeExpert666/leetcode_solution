package BinaryTree

import "container/list"

/**
 * 方法一：递归法
 */
func findBottomLeftValue(root *TreeNode) int {
	res, _ := findValue(root, 0)
	return res
}

func findValue(root *TreeNode, depth int) (int, int) {
	// 空树没有最下层左节点，故将该结点层数设置为-1
	if root == nil {
		return -1, -1
	}
	// 获取左右子树的最下层左节点
	leftVal, leftDepth := findValue(root.Left, depth+1)
	rightVal, rightDepth := findValue(root.Right, depth+1)
	// 若左右子树均为空，则根节点是最下层左节点
	if leftDepth == -1 && rightDepth == -1 {
		return root.Val, depth
	}
	// 左右子树至少有一个不空时,比较最左结点深度，选取较深的结点
	// 若深度相同，则选取较左结点
	if leftDepth < rightDepth {
		return rightVal, rightDepth
	} else {
		// 包含深度相同情况
		return leftVal, leftDepth
	}
}

/**
 * 方法二：层序遍历迭代
 */
func findBottomLeftValue1(root *TreeNode) int {
	var res int

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				// 处理每一层的最左结点
				res = cur.Val
			}
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
	}
	return res
}
