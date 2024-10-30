package BinaryTree

// 递归法
func PreorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	PreorderTraversalHelper(root, &res)
	return res
}

func PreorderTraversalHelper(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		PreorderTraversalHelper(root.Left, res)
		PreorderTraversalHelper(root.Right, res)
	}
}

// 迭代法，利用栈
func PreorderTraversal1(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	p := root // 指向当前遍历树结点
	for p != nil || len(stack) > 0 {
		if p != nil {
			res = append(res, p.Val) // 先序遍历，入栈前访问
			stack = append(stack, p)
			p = p.Left // 不断访问左子树
		} else { // 出栈访问右子树
			p = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
