package BinaryTree

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 方法一： 利用队列层序遍历，注意要求输出结果与常规层序遍历不同，
 * 需要对结果进行分层展示，所以与常规实现略有不同，需要统计每层长度，两种均需熟练掌握
 */
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int

	if root == nil { // 空树单独处理
		return res
	}

	// 利用队列处理二叉树的层序遍历
	var queue []*TreeNode

	// 初始化
	queue = append(queue, root)

	for l := len(queue); l > 0; l = len(queue) {
		var layer []int          // 记录每一层的结点
		for i := 0; i < l; i++ { // 分层访问
			// 出队并访问
			node := queue[0]
			queue = queue[1:]               // 出队
			layer = append(layer, node.Val) // 访问
			// 入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, layer)
	}
	return res
}

/**
 * 方法二：递归法，不常用，但可以理解一下
 */
func levelOrder2(root *TreeNode) [][]int {
	// 递归方法
	var res [][]int
	lOrder(root, &res, 0)
	return res
}

func lOrder(root *TreeNode, res *[][]int, depth int) {
	if root != nil {
		if depth == len(*res) { // 注意，depth从0开始，写成depth+1 > len(*res)也可以
			// 若当前访问元素所在的层还未被访问过，也即还未建立当前层对应的切片
			// 则建立该切片并添加相应元素
			*res = append(*res, []int{root.Val})
		} else {
			// 若所在层已被访问过，则将当前元素添加在该层对应切片末尾
			// layer := (*res)[depth]
			// layer = append(layer, root.Val)  这种写法是不对的，因为这样是值传递，只会改变layer，不会改变原切片
			// layer := &(*res)[depth]
			// *layer = append(*layer, root.Val)  这种写法是可以的，引用传递
			(*res)[depth] = append((*res)[depth], root.Val)
		}
		lOrder(root.Left, res, depth+1)  // 处理左子树
		lOrder(root.Right, res, depth+1) // 处理右子树
	}
}
