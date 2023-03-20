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

/**
 * 方法二：层序遍历模板，主要目标是找到第一个叶子结点
 */
func minDepth2(root *TreeNode) int {
	var res int

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		res++
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			if cur.Left == nil && cur.Right == nil {
				return res
			}
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
	}
	return res // 实际不会走到这一步
}

func main() {
	n3 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	n2 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: n3,
	}
	n1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: n2,
	}
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: n1,
	}
	fmt.Println(minDepth2(root))
}
