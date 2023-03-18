package BinaryTree

import "container/list"

/**
 * 层序遍历模板
 */
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var sum int
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
			sum += cur.Val
		}
		res = append(res, float64(sum)/float64(l))
	}
	return res
}
