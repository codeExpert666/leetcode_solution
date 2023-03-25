package BinaryTree

import "container/list"

/**
 * 递归法一，本质是将第二棵树合并到了第一棵树上，当遇到第一棵树上不存在的部分时，本质上是直接指向了第二棵树多出来的部分
 * 这种方法并没有创建任何一个新的结点，只是修改了原有结点的指针，缺点是破坏了原有树的结构
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		// 当两棵树均不为空时，将第二棵树合并到第一棵树上
		// 根节点值相加
		root1.Val += root2.Val
		// 递归合并两课树的左右子树
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
		// 此时第一棵树合并好的树
		return root1
	}
	// 当两棵树中有一颗非空时，返回非空树的根节点
	if root1 != nil {
		return root1
	}
	if root2 != nil {
		return root2
	}
	// 当两棵树均为空时，返回nil
	return nil
}

/**
 * 递归法二：不破坏原有两棵树的结构，合并后树的每个结点都是单独创建的
 */
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
		root.Right = mergeTrees1(root1.Right, root2.Right)
	}
	return root
}

/**
 * 递归法三：对于两棵树都有的结点，合并树新建结点，对于一方有，一方没有的结点，直接引用
 * 这种方式属于部分新建结点
 */
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		// 当两棵树均不为空时，创建新结点
		root := &TreeNode{}
		// 根节点值相加
		root.Val = root1.Val + root2.Val
		// 递归合并两棵树的左右子树
		root.Left = mergeTrees2(root1.Left, root2.Left)
		root.Right = mergeTrees2(root1.Right, root2.Right)

		return root
	}
	// 当两棵树中有一颗非空时，返回非空树的根节点
	if root1 != nil {
		return root1
	}
	if root2 != nil {
		return root2
	}
	// 当两棵树均为空时，返回nil
	return nil
}

/**
 * 迭代法：层序遍历模板，和leetcode101题中判断二叉树对称的思路一致
 * 一般如果题目牵涉到两棵树的结构相关问题时，可以考虑使用层序遍历的方法
 * 在本题中，这种迭代法的本质是将两棵树合并到了第一棵树上，当遇到第一棵树上不存在的部分时，本质上是直接指向了第二棵树多出来的部分
 */
func mergeTrees3(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 因为结点出队后一定需要对孩子结点进行判断，所以结点一定不能为空，故需要先进行判空处理
	// 后序的结点在迭代的过程中可以保证一定不空

	// 当有空树存在时
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 当两棵树不空时
	queue := list.New()
	queue.PushBack(root1)
	queue.PushBack(root2)

	for queue.Len() > 0 {
		leftNode := queue.Remove(queue.Front()).(*TreeNode)
		rightNode := queue.Remove(queue.Front()).(*TreeNode)
		// 此时两个节点一定不为空，val相加
		leftNode.Val += rightNode.Val
		// 如果两棵树左节点都不为空，加入队列
		if leftNode.Left != nil && rightNode.Left != nil {
			queue.PushBack(leftNode.Left)
			queue.PushBack(rightNode.Left)
		}
		// 如果两棵树右节点都不为空，加入队列
		if leftNode.Right != nil && rightNode.Right != nil {
			queue.PushBack(leftNode.Right)
			queue.PushBack(rightNode.Right)
		}
		// 当root1的左节点 为空 root2左节点不为空，就赋值过去
		if leftNode.Left == nil && rightNode.Left != nil {
			leftNode.Left = rightNode.Left
		}
		// 当root1的右节点 为空 root2右节点不为空，就赋值过去
		if leftNode.Right == nil && rightNode.Right != nil {
			leftNode.Right = rightNode.Right
		}
	}
	return root1
}
