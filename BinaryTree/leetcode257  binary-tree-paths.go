package BinaryTree

import (
	"container/list"
	"strconv"
)

/**
 * 方法一：递归法，原理是一棵树的叶子结点同时也是其左右子树的叶子结点
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)
	// 拼接过程
	left = append(left, right...) // 左右拼接
	// 路径头部拼接根节点
	v := strconv.Itoa(root.Val)
	if len(left) == 0 {
		left = append(left, v)
	} else {
		for i := 0; i < len(left); i++ {
			left[i] = strconv.Itoa(root.Val) + "->" + left[i]
		}
	}
	return left
}

/**
 * 方法二：迭代法，利用后序遍历模板
 */
func binaryTreePaths1(root *TreeNode) []string {
	var res []string
	var path []int // 记录此时的路径，先用结点的Val值组成的切片表示

	stack := list.New()
	cur := root
	var pre *TreeNode

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			// 补充当前路径，添加根节点
			path = append(path, cur.Val)
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Back().Value.(*TreeNode)
			if cur.Right == nil || cur.Right == pre {
				if cur.Left == nil && cur.Right == nil {
					// 遍历到叶子结点
					res = append(res, pathString(path))
				}
				// 删减当前路径，剔除根节点
				path = path[:len(path)-1]
				stack.Remove(stack.Back())
				pre = cur
				cur = nil
			} else {
				cur = cur.Right
			}
		}
	}
	return res
}

/**
 * 将int切片路径转换成字符串
 */
func pathString(p []int) string {
	res := strconv.Itoa(p[0])
	for i := 1; i < len(p); i++ {
		res += "->" + strconv.Itoa(p[i])
	}
	return res
}
