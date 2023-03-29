package BinaryTree

import "container/list"

/**
 * 递归；取中间数作为根节点，左半部分递归构造左子树，右半部分递归构造右子树
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

/**
 * 迭代法：通过三个队列来模拟，一个队列放遍历的节点，一个队列放左区间下标，一个队列放右区间下标
 * 三个队列相同位置相互对应，处理的是某一个区间的根节点选择问题
 */
func sortedArrayToBST1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{}            // 初始根节点
	nodeQueue := list.New()        // 放遍历的节点
	leftBorderQueue := list.New()  // 保存左区间下标
	rightBorderQueue := list.New() // 保存右区间下标
	// 初始化
	nodeQueue.PushBack(root)                 // 根节点入队列
	leftBorderQueue.PushBack(0)              // 0为左区间下标初始位置
	rightBorderQueue.PushBack(len(nums) - 1) // len(nums) - 1为右区间下标初始位置

	for nodeQueue.Len() > 0 {
		// 处理根节点
		curNode := nodeQueue.Remove(nodeQueue.Front()).(*TreeNode)
		leftBorder := leftBorderQueue.Remove(leftBorderQueue.Front()).(int)
		rightBorder := rightBorderQueue.Remove(rightBorderQueue.Front()).(int)
		mid := leftBorder + (rightBorder-leftBorder)/2 // 中间值作为根节点
		curNode.Val = nums[mid]
		// 处理左区间
		if leftBorder <= mid-1 {
			curNode.Left = &TreeNode{}
			nodeQueue.PushBack(curNode.Left)
			leftBorderQueue.PushBack(leftBorder)
			rightBorderQueue.PushBack(mid - 1)
		}
		// 处理右区间
		if rightBorder >= mid+1 {
			curNode.Right = &TreeNode{}
			nodeQueue.PushBack(curNode.Right)
			leftBorderQueue.PushBack(mid + 1)
			rightBorderQueue.PushBack(rightBorder)
		}
	}
	return root
}
