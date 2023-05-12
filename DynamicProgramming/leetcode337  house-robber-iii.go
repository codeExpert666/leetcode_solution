package DynamicProgramming

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 递归法，记忆化递推，换句话说就是把递归过程的中间结果进行存储，避免重复计算
 */
var record = make(map[*TreeNode]int) // 记录递归过程的中间结构，避免重复计算，功能类似于dp数组

func rob2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 第一种情况：抢劫根节点，此时不能抢劫左右孩子
	tmp1 := root.Val // 该种情况下的最大金额
	if root.Left != nil {
		// 计算左孩子的左子树的最大金额
		ll, e1 := record[root.Left.Left]
		if !e1 {
			ll = rob2(root.Left.Left)
			record[root.Left.Left] = ll // 记录中间结果
		}
		tmp1 += ll
		// 计算左孩子的右子树的最大金额
		lr, e2 := record[root.Left.Right]
		if !e2 {
			lr = rob2(root.Left.Right)
			record[root.Left.Right] = lr // 记录中间结果
		}
		tmp1 += lr
	}
	if root.Right != nil {
		// 计算右孩子的左子树的最大金额
		rl, e3 := record[root.Right.Left]
		if !e3 {
			rl = rob2(root.Right.Left)
			record[root.Right.Left] = rl
		}
		tmp1 += rl
		// 计算右孩子的右子树的最大金额
		rr, e4 := record[root.Right.Right]
		if !e4 {
			rr = rob2(root.Right.Right)
			record[root.Right.Right] = rr
		}
		tmp1 += rr
	}
	// 第二种情况：不抢劫根节点，抢劫左右子树
	tmp2 := 0
	// 计算根节点左子树最大金额
	l, e5 := record[root.Left]
	if !e5 {
		l = rob2(root.Left)
		record[root.Left] = l
	}
	tmp2 += l
	// 计算根节点右子树最大金额
	r, e6 := record[root.Right]
	if !e6 {
		r = rob2(root.Right)
		record[root.Right] = r
	}
	tmp2 += r

	return max(tmp1, tmp2)
}

/**
 * 思路仍与上一方法一致，但实现细节上有所改进，代码更加简洁
 * PS：还是对递归的理解不够深入
 */
func rob3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if v, e := record[root]; e { // 上一方法中的查表操作写在了这里，仔细体会
		return v
	}
	// 第一种情况：抢劫根节点，此时不能抢劫左右孩子
	tmp1 := root.Val
	if root.Left != nil {
		tmp1 += rob3(root.Left.Left) + rob3(root.Left.Right)
	}
	if root.Right != nil {
		tmp1 += rob3(root.Right.Left) + rob3(root.Right.Right)
	}
	// 第二种情况：不抢劫根节点，抢劫左右子树
	tmp2 := rob3(root.Left) + rob3(root.Right)

	res := max(tmp1, tmp2) // 取二者中的较大者
	record[root] = res     // 记录操作写在了这里
	return res
}

/**
 * 树形动态规划（dp）
 * 树形结构天然适合递归，而递归过程会保存一个递归调用栈，栈中记录了每一层节点的信息，这就可以类比为dp数组
 * 而递归出栈的过程就可以理解为dp数组的递推过程
 * 在本题中，递推方式应是由下至上的，对应着树的后序遍历，dp[i]代表着每一层节点的信息，这里将其构建为一个二维数组
 * 下标为0记录不偷该节点所得到的的最大金钱，下标为1记录偷该节点所得到的的最大金钱。
 */
func rob4(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

func robTree(root *TreeNode) [2]int {
	if root == nil { // dp初值
		return [2]int{0, 0}
	}
	// dp[i-1]
	left := robTree(root.Left)
	right := robTree(root.Right)
	// dp递推
	// 抢劫根节点
	val1 := root.Val + left[0] + right[0]
	// 不抢劫根节点
	val2 := max(left[0], left[1]) + max(right[0], right[1])

	return [2]int{val2, val1}
}
