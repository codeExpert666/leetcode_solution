package BinaryTree

/**
 * 迭代法：若待删除值小于根节点，则往左子树遍历；若待删除值大于根节点，则往右子树遍历
 * 若相等，则删除该结点。删除方法是：若该节点是叶子节点，则直接删除，若不是叶子结点，
 * 则用其左子树的最右结点或右子树的最左结点替代它，再删除叶子结点
 * 此方法是我自己想到的放，寻找待删除结点的过程使用了迭代法，删除结点使用了递归法
 */
func deleteNode5(root *TreeNode, key int) *TreeNode {
	var parent *TreeNode
	cur := root

	for cur != nil {
		if key < cur.Val {
			parent = cur
			cur = cur.Left
		} else if key > cur.Val {
			parent = cur
			cur = cur.Right
		} else {
			// 找到待删除结点，执行删除操作
			if cur == root {
				// 待删除结点恰好是根节点
				root = deleteRoot2(cur)
			} else if cur == parent.Left {
				parent.Left = deleteRoot2(cur)
			} else {
				parent.Right = deleteRoot2(cur)
			}
			break
		}
	}
	return root
}

// 删除一颗二叉搜索树的根节点
func deleteRoot(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		// 根节点是叶子结点，直接删除
		return nil
	}
	var parent, cur *TreeNode // 指向当前结点与其父节点
	if root.Left != nil {
		// 寻找左子树的最右的结点
		parent = root
		cur = root.Left
		for cur.Right != nil {
			parent = cur
			cur = cur.Right
		}
		// cur指向最右结点，将cur结点值与根节点值互换
		//root.Val, cur.Val = cur.Val, root.Val
		// 不用互换，只需保存好cur结点的值即可
		root.Val = cur.Val
		// 递归删除以cur为根节点的树
		if parent == root {
			// 最右结点恰好是根结点的左孩子
			parent.Left = deleteRoot(cur)
		} else {
			parent.Right = deleteRoot(cur)
		}
		return root
	} else {
		// 该树只有右子树，直接返回右子树根节点即可
		return root.Right
	}
}

// 删除一颗二叉搜索树的根节点(第二种写法，无需迭代，更好的写法)
func deleteRoot2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Right == nil {
		// 右子树为空，直接返回左子树的根节点
		return root.Left
	}
	// 寻找右子树的最左结点
	cur := root.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	// !!!非常重要的一步，直接将左子树接在最左结点的左侧
	cur.Left = root.Left
	// 完美删除根节点
	return root.Right
}

/**
 * 完全的递归法，与迭代法的思路一致，只是处理右子树最左结点的方式不同
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	minNode := root.Right
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	root.Val = minNode.Val
	root.Right = deleteNode1(root.Right)
	return root
}

// 递归删除二叉搜索树的最左结点
func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}
