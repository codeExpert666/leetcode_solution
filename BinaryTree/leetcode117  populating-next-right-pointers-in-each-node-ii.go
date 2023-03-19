package BinaryTree

import "container/list"

/**
 * 层序遍历模板，实际上这种方法不符合题目要求的O(1)空间复杂度，方法二的递归符合
 */
func connect3(root *Node) *Node {
	// 特殊处理空树
	if root == nil {
		return root
	}

	// 初始化队列
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*Node)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
			if i < l-1 {
				// cur在本层有下一结点
				cur.Next = queue.Front().Value.(*Node)
			}
		}
	}
	return root
}

/**
 * 递归方法，思路就是对于树中的每个结点，都考虑两件事：1、连接本结点的左右孩子。2、将本结点的右孩子与兄弟结点的左孩子连接
 * 本题的递归方法与116题的递归思路完全一致，但由于本题的二叉树不是完全二叉树，而是任意的二叉树，所以需要考虑更多的细节
 * 具体包括：1、当前结点是否存在子树，如果有，哪个子树与兄弟结点的子树相连 2、当前结点是否需要连接左右子树
 * 3、当前结点的子树与哪个兄弟的结点的子树相连（从左往右第一个拥有子树的兄弟节点）
 */
func connect(root *Node) *Node {
	if root != nil {
		if root.Left != nil {
			if root.Right != nil {
				connectNext(root.Right, root.Next) // 右子树连接兄弟结点的子树
				root.Left.Next = root.Right        // 连接左右子树
			} else {
				connectNext(root.Left, root.Next) // 左子树连接兄弟结点的子树
			}
		} else {
			if root.Right != nil {
				connectNext(root.Right, root.Next) // 右子树连接兄弟结点的子树
			}
		}

		// 一定先处理右子树，因为这样next才不会断
		connect(root.Right) // 处理右子树
		connect(root.Left)  // 处理左子树
	}
	return root
}

/**
 * 连接当前结点的靠右子树与有子树的Next结点的相应子树
 */
func connectNext(rightNode, next *Node) {
	for next != nil {
		// 当兄弟结点存在时，且兄弟结点至少有一个子树，需要执行2
		if next.Left != nil {
			// 如果兄弟结点有左子树，则靠右结点的next指向该左子树
			rightNode.Next = next.Left
			break
		} else if next.Right != nil {
			// 如果兄弟节点没有左子树，则靠右结点直接连接兄弟结点的右子树
			rightNode.Next = next.Right
			break
		} else {
			// 如果兄弟结点没有子树，则寻找下一个兄弟结点
			next = next.Next
		}
	}
}
