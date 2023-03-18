package BinaryTree

/**
 * 本题虽然也是层序遍历，但并不适合递归，因为二叉树的深度未知，且在递归过程中也不好获取
 * 所以本体只好采用队列的层序遍历方法，最后再将结果反转
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var queue []*TreeNode // 队列
	queue = append(queue, root)

	for len(queue) > 0 {
		var layer []int
		// 注意：一定要先获取长度，不能在下面的for循环中用Len()获取长度
		// 因为在循环的入队出队过程中，长度是动态变化的
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			layer = append(layer, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, layer)
	}
	// 反转结果切片
	for j, k := 0, len(res)-1; j < k; j, k = j+1, k-1 {
		res[j], res[k] = res[k], res[j]
	}
	// 反转结果切片另一种写法
	//for i:=0; i<len(res)/2; i++ {
	//	res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	//}

	return res
}
