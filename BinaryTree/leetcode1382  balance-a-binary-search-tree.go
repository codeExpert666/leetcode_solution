package BinaryTree

func BalanceBST(root *TreeNode) *TreeNode {
	// 先对搜索树中序遍历获取有序数组
	var inOrderList []int
	var travel func(*TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		inOrderList = append(inOrderList, node.Val)
		travel(node.Right)
	}
	travel(root)
	// 根据有序数组建立平衡二叉搜索树
	var build func(int, int) *TreeNode
	build = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := start + (end-start)>>1
		return &TreeNode{
			Val:   inOrderList[mid],
			Left:  build(start, mid-1),
			Right: build(mid+1, end),
		}
	}
	return build(0, len(inOrderList)-1)
}
