package main

import (
	"fmt"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const INT_MAX = int(^uint(0) >> 1)

var pre *TreeNode

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return INT_MAX
	}
	// 根据中序遍历的顺序，先把左子树最小值当作全局最小值
	min := getMinimumDifference(root.Left)
	// 更新最小值
	if pre != nil {
		tmp := root.Val - pre.Val
		if tmp < min {
			min = tmp
		}
	}
	pre = root
	// 右子树最小值
	rightMin := getMinimumDifference(root.Right)
	// 更新最小值
	if rightMin < min {
		return rightMin
	} else {
		return min
	}
}

func main() {
	//n5 := &TreeNode{
	//	Val:   5,
	//	Left:  nil,
	//	Right: nil,
	//}
	n4 := &TreeNode{
		Val:   49,
		Left:  nil,
		Right: nil,
	}
	n3 := &TreeNode{
		Val:   12,
		Left:  nil,
		Right: nil,
	}
	n2 := &TreeNode{
		Val:   48,
		Left:  n3,
		Right: n4,
	}
	n1 := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  n1,
		Right: n2,
	}
	fmt.Println(getMinimumDifference(root))
}
