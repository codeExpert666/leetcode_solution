package BinaryTree

/**
 * 递归法，和leetcode105思路一致
 */
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}
	var leftNodeNum int
	for i, v := range inorder {
		if v == root.Val {
			leftNodeNum = i
		}
	}
	left := buildTree1(inorder[:leftNodeNum], postorder[:leftNodeNum])
	right := buildTree1(inorder[leftNodeNum+1:], postorder[leftNodeNum:len(postorder)-1])
	root.Left, root.Right = left, right
	return root
}
