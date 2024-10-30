package BinaryTree

// 递归法
func PostorderTraversal(root *TreeNode) []int {
	var res []int
	PostorderTraversalHelper(root, &res)
	return res
}

func PostorderTraversalHelper(root *TreeNode, res *[]int) {
	if root != nil {
		PostorderTraversalHelper(root.Left, res)
		PostorderTraversalHelper(root.Right, res)
		*res = append(*res, root.Val)
	}
}

// 迭代法（有难度，容易忘）
func PostorderTraversal1(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	// cur指向当前访问结点，pre指向上一访问结点
	// 由于是后序遍历，栈中元素需要访问两次，第二次才能出栈
	// pre的设置可以辅助判断是否为第二次出栈
	var pre, cur *TreeNode = nil, root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			// 注意这里cur并没有被访问，不应更新pre
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			if cur.Right == nil || cur.Right == pre {
				// 关键一步，通过观察cur是否从右子树退出来判断是否为第二次出栈
				// 或者是右子树为空，这种情况pre不可能到右侧，但此时cur应该被访问
				res = append(res, cur.Val)
				stack = stack[:len(stack)-1]
				pre = cur // pre指向上一个访问结点，只有这种情况结点才算是被访问
				cur = nil // 为了下一次继续从栈中获取结点
			} else {
				// 注意这里cur并没有被访问，不应更新pre
				cur = cur.Right
			}
		}
	}
	return res
}
