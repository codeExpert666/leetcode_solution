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

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
		return sum
	}
	return sum + sumOfLeftLeaves(root.Left)
}

func main() {
	//n5 := &TreeNode{
	//	Val:   5,
	//	Left:  nil,
	//	Right: nil,
	//}
	n4 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n3 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	n2 := &TreeNode{
		Val:   20,
		Left:  n3,
		Right: n4,
	}
	n1 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:   2,
		Left:  n1,
		Right: n2,
	}
	fmt.Println(sumOfLeftLeaves(root))
}
