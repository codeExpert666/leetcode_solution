package BinaryTree

import "container/list"

/**
 * 方法一：递归，思路和leetcode112一致，从找到路径，变成记录所有路径
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return [][]int{[]int{root.Val}}
	}
	left := pathSum(root.Left, targetSum-root.Val)
	right := pathSum(root.Right, targetSum-root.Val)

	left = append(left, right...) // 左右路径合并
	if len(left) != 0 {
		// 存在路径，在所有路径的头部添加当前结点值
		for i := 0; i < len(left); i++ {
			left[i] = append([]int{root.Val}, left[i]...)
		}
	}
	// 左右子树不存在路径，且当前也不是叶子节点或是叶子结点但总合不相等，直接返回
	return left
}

/**
 * 方法二，迭代，思路与leetcode112完全一致，从找到路径，变成记录所有路径，代码变化不大
 */
func pathSum1(root *TreeNode, targetSum int) [][]int {
	var sum int
	var path []int
	var res [][]int

	stack := list.New()
	cur := root
	var pre *TreeNode

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			sum += cur.Val               // 更新此时的总和
			path = append(path, cur.Val) // 更新当前路径
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Back().Value.(*TreeNode)
			if cur.Right == nil || cur.Right == pre {
				if cur.Left == nil && cur.Right == nil && sum == targetSum {
					// 遍历到叶子结点,且此时和等于目标值,也即找到路径
					res = append(res, append([]int{}, path...))
					// ！！！一定注意：这样写是不对的，因为go语言的切片本质上是一个结构体
					// 结构体中包含指针，指针指向底层数组，这种写法只是且切片结构体（指针）复制
					// 一份，当底层path数组改变时，已经存在res中的切片内容也可能因此改变，而在本题中
					// 底层数组是需要实时更新的
					// res = append(res, path)
				}
				stack.Remove(stack.Back())
				sum -= cur.Val            // 更新此时的总和
				path = path[:len(path)-1] // 更新此时路径
				pre = cur
				cur = nil
			} else {
				cur = cur.Right
			}
		}
	}
	return res
}
