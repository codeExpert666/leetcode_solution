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
					res = append(res, path)
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
	//n3 := &TreeNode{
	//	Val:   15,
	//	Left:  nil,
	//	Right: nil,
	//}
	n2 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	n1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  n1,
		Right: n2,
	}
	fmt.Println(pathSum1(root, 3))
}
