package BinaryTree

/**
 * 递归法，确定根节点，再确定左右子树在两个输入数组中的范围，递归构造左右子树，最后与根节点相连
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		// 空树
		return nil
	}
	// 构建根节点，前序遍历的第一个结点
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	// 确定根节点在中序遍历的位置，进而分割左右子树
	var leftNodeNum int // 记录左子树结点数目
	for i, v := range inorder {
		if v == preorder[0] {
			leftNodeNum = i
			break
		}
	}
	// 递归构建左右子树
	left := buildTree(preorder[1:1+leftNodeNum], inorder[:leftNodeNum])
	right := buildTree(preorder[1+leftNodeNum:], inorder[leftNodeNum+1:])
	// 将左右子树与根节点相连
	root.Left = left
	root.Right = right

	return root
}
