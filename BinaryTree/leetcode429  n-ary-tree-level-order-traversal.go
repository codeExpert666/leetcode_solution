package BinaryTree

import "container/list"

// Node Definition for a node.
type Node struct {
	Val      int
	Children []*Node
}

/**
 * N叉树的层序遍历，与二叉树的思路一致，仍然可以采用基于队列的迭代方法
 */
func levelOrder3(root *Node) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var layer []int
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*Node)
			for _, node := range cur.Children {
				queue.PushBack(node)
			}
			layer = append(layer, cur.Val)
		}
		res = append(res, layer)
	}
	return res
}

/**
 * 递归解法，也与二叉树的层序遍历递归法一致
 */
func levelOrder4(root *Node) [][]int {
	var res [][]int
	nBranchLevelOrder(root, &res, 0)
	return res
}

func nBranchLevelOrder(root *Node, res *[][]int, depth int) { // depth从0开始
	if root != nil {
		if depth+1 > len(*res) {
			*res = append(*res, []int{})
		}
		(*res)[depth] = append((*res)[depth], root.Val)
		for _, node := range root.Children {
			nBranchLevelOrder(node, res, depth+1)
		}
	}
}
