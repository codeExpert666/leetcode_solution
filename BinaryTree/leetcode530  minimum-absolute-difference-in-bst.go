package BinaryTree

import (
	"container/list"
	"math"
)

const INT_MAX = int(^uint(0) >> 1)

/**
 * 递归法：利用二叉搜索树的性质即可，中序遍历思路，与leetcode98相似
 */
var pre1 *TreeNode

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return INT_MAX
	}
	// 根据中序遍历的顺序，先把左子树最小值当作全局最小值
	min := getMinimumDifference(root.Left)
	// 更新最小值
	if pre1 != nil {
		tmp := root.Val - pre1.Val
		if tmp < min {
			min = tmp
		}
	}
	pre1 = root
	// 右子树最小值
	rightMin := getMinimumDifference(root.Right)
	// 更新最小值
	if rightMin < min {
		return rightMin
	} else {
		return min
	}
}

/**
 * 递归法，和上一种方法完全一致，但写法更好，更具go语言特色（go语言的函数也是一种变量），值得学习
 */
func getMinimumDifference2(root *TreeNode) int {
	// 保留前一个节点的指针
	var prev *TreeNode
	// 定义一个比较大的值
	min := math.MaxInt64
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return min
}

/**
 * 迭代法，中序遍历模板修改
 */
func getMinimumDifference1(root *TreeNode) int {
	stack := list.New()
	cur := root
	var pre *TreeNode // 记录前一个结点，因为要不断与前一结点进行比较
	res := INT_MAX    // 初始化为int类型最大值

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			// 与前一结点比较
			if pre != nil {
				tmp := cur.Val - pre.Val
				// 若差值更小，则更新最小值
				if tmp < res {
					res = tmp
				}
			}
			pre = cur
			cur = cur.Right
		}
	}
	return res
}
