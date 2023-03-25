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

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	// 新建结点，不利用已有结点
	root := &TreeNode{}
	if root1 == nil {
		// 第一棵树空，第二棵树不空
		root.Val = root2.Val
		root.Left = mergeTrees1(nil, root2.Left)
		root.Right = mergeTrees1(nil, root2.Right)
	} else if root2 == nil {
		// 第一棵树不空，第二棵树空
		root.Val = root1.Val
		root.Left = mergeTrees1(root1.Left, nil)
		root.Right = mergeTrees1(root1.Right, nil)
	} else {
		root.Val = root1.Val + root2.Val
		root.Left = mergeTrees1(root1.Left, root2.Left)
		root.Left = mergeTrees1(root1.Right, root2.Right)
	}
	return root
}

func main() {
	//n5 := &TreeNode{
	//	Val:   5,
	//	Left:  nil,
	//	Right: nil,
	//}
	//n4 := &TreeNode{
	//	Val:   7,
	//	Left:  nil,
	//	Right: nil,
	//}
	n2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root2 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: n2,
	}
	n1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root1 := &TreeNode{
		Val:   1,
		Left:  n1,
		Right: nil,
	}
	fmt.Println(mergeTrees1(root1, root2))
}
