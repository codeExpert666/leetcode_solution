package BinaryTree

import "container/list"

// INT_MIN 32位int类型的最小值（补码）
const INT_MIN = ^int(^uint(0) >> 1)

/**
 * 层序遍历模板
 */
func largestValues(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		max := INT_MIN
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
			if cur.Val > max {
				max = cur.Val
			}
		}
		res = append(res, max)
	}
	return res
}
