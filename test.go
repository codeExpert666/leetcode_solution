package main

import (
	"container/list"
	"fmt"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil { // 空树特殊处理
		return true
	}
	if root.Left == nil && root.Right != nil {
		return false
	}
	if root.Left != nil && root.Right == nil {
		return false
	}

	leftStack := list.New()
	rightStack := list.New()
	leftPtr, rightPtr := root.Left, root.Right

	for leftPtr != nil || leftStack.Len() > 0 {
		if leftPtr != nil {
			if rightPtr == nil {
				return false
			}
			leftStack.PushBack(leftPtr)
			rightStack.PushBack(rightPtr)
			leftPtr = leftPtr.Left
			rightPtr = rightPtr.Right
		} else {
			if rightPtr != nil {
				return false
			}
			leftPtr = leftStack.Remove(leftStack.Back()).(*TreeNode)
			rightPtr = rightStack.Remove(rightStack.Back()).(*TreeNode)
			if leftPtr.Val != rightPtr.Val {
				return false
			}
			leftPtr = leftPtr.Right
			rightPtr = rightPtr.Left
		}
	}
	return true
}

func main() {
	n5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	n4 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	n3 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	n2 := &TreeNode{
		Val:   3,
		Left:  n5,
		Right: nil,
	}
	n1 := &TreeNode{
		Val:   3,
		Left:  n3,
		Right: n4,
	}
	root := &TreeNode{
		Val:   2,
		Left:  n1,
		Right: n2,
	}
	fmt.Println(isSymmetric(root))
}
