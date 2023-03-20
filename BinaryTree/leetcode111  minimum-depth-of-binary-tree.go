package BinaryTree

import (
	"container/list"
	"fmt"
)

/**
 * 方法一：递归，与最大深度一致，稍微改个符号并多考虑一些细节
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	LeftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if LeftDepth >= rightDepth {
		// 注意这里的判0操作，这是因为若较小的子树深度为0，则说明该子树不存在
		// 所以最小深度应为另一子树深度加1
		if rightDepth == 0 {
			return LeftDepth + 1
		} else {
			return rightDepth + 1
		}
	} else {
		if LeftDepth == 0 {
			return rightDepth + 1
		} else {
			return LeftDepth + 1
		}
	}
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
	fmt.Println(minDepth(root))
}
