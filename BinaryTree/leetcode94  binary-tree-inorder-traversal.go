package BinaryTree

// 递归法
func InorderTraversal(root *TreeNode) []int {
	var res []int
	InorderTraversalHelper(root, &res)
	return res
}

func InorderTraversalHelper(root *TreeNode, res *[]int) {
	if root != nil {
		InorderTraversalHelper(root.Left, res)
		*res = append(*res, root.Val)
		InorderTraversalHelper(root.Right, res)
	}
}

func InorderTraversal1(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1] // 中序遍历，出栈再访问
			res = append(res, p.Val)
			stack = stack[:len(stack)-1]
			p = p.Right // 访问右子树
		}
	}
	return res
}
