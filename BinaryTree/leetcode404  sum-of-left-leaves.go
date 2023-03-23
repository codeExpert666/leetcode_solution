package BinaryTree

import "container/list"

/**
 * 方法一：递归法，本题的难点在于当前结点只能判断他的左节点是不是左叶子结点，无法判断自己，右子树正常递归即可
 * 左子树递归时只需要特殊处理左叶子结点即可
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
		return sum
	}
	return sum + sumOfLeftLeaves(root.Left)
}

/**
 * 方法二：迭代，中序遍历模板。思路：中序遍历会遍历树中所有结点，那么在遍历每个结点时判断该节点是否为左叶子结点
 * 然后累加，但问题是当前遍历结点无法判断自身是否为左叶子结点，它只可以判断它的左孩子是否为左叶子结点，故这样行不通，那么，
 * 我们可以把这个访问操作理解成“判断左孩子是否为左叶子结点”，这样的话，最终会漏下根节点没有被判断，但根节点一定不可能是做叶子结点
 * 故这种处理是没有问题的。
 */
func sumOfLeftLeaves1(root *TreeNode) int {
	var res int
	stack := list.New()
	cur := root

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			if cur.Left != nil && cur.Left.Left == nil && cur.Left.Right == nil {
				res += cur.Left.Val
			}
			cur = cur.Right
		}
	}
	return res
}
