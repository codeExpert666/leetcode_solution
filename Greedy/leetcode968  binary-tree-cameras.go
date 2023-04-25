package Greedy

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 贪心法
 * 从下往上看，局部最优：让叶子节点的父节点安摄像头，所用摄像头最少，整体最优：全部摄像头数量所用最少！
 * 从下往上看的原因是：头结点放不放摄像头也就省下一个摄像头，叶子节点放不放摄像头省下了的摄像头数量是指数阶别的。
 * 两个关键点：1、后序遍历（从下往上） 2、每隔两个结点放一个摄像头
 * 每个节点有3个状态，用三个数字表示：0：该节点无覆盖；1：本节点有摄像头；2：本节点有覆盖
 */
func minCameraCover(root *TreeNode) int {
	var res int // 记录摄像头数量
	var travel func(*TreeNode) int
	// 后序遍历，返回值是当前树的根节点的状态
	travel = func(node *TreeNode) int {
		// 空节点的状态是已覆盖
		if node == nil {
			return 2
		}
		left := travel(node.Left)   // 获取左子树根节点状态
		right := travel(node.Right) // 获取右子树根节点状态
		// 左右根节点至少一个未覆盖
		if left == 0 || right == 0 {
			res++
			return 1 // 当前节点需要有摄像头
		}
		// 左右节点均已覆盖
		if left == 2 && right == 2 {
			return 0 // 当前节点的覆盖交给父节点（贪心，尽量减少摄像头数量）
		}
		// 左右节点至少一个有摄像头
		if left == 1 || right == 1 {
			return 2 // 当前节点已覆盖
		}
		return -1 // 实际不会走到这一步，上面三个分支实际已经分析了所有情况，只是没用else
	}
	// 可能存在根节点没有被覆盖的情况（例如左右节点均为已覆盖）
	if rootState := travel(root); rootState == 0 {
		res++ // 根节点需要补充摄像头
	}
	return res
}
