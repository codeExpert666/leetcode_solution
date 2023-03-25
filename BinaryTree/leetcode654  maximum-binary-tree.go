package BinaryTree

/**
 * 题目提到递归地构建，故采用递归法
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// 处理空数组，也即空树
	if len(nums) == 0 {
		return nil
	}
	// 先寻找数组中的最大值，也即根节点
	// 寻找最大值正常遍历就行，双指针可以，但空间复杂度会高一点，没必要
	var index int
	for i, v := range nums {
		// 前提数组中值各不相同，所以不存在相等情况
		if nums[index] < v {
			index = i
		}
	}
	// 构建根节点以及左右子树
	root := &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}

	return root
}
