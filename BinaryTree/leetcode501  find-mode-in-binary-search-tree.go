package BinaryTree

import "math"

/**
 * 递归法，先递归统计结点值出现情况，再寻找众数，这种方法用什么遍历方式都可以，需要一个hash表记录出现频率
 * 空间复杂度较高
 */
func findMode1(root *TreeNode) []int {
	var res []int
	maxValue := math.MinInt
	// 获取节点值出现情况
	countBST(root)
	// 遍历hash表，寻找value值最大的key值
	for k, v := range countTable {
		if v > maxValue {
			// 更新结果列表
			res = res[:0]
			res = append(res, k)
			// 更新当前最大值
			maxValue = v
		} else if v == maxValue {
			// 最大值不止一个
			res = append(res, k)
		}
	}
	return res
}

// 中序遍历记录结点值出现情况
var countTable = make(map[int]int, 0)

func countBST(root *TreeNode) {
	if root == nil {
		return
	}
	countBST(root.Left)
	countTable[root.Val]++
	countBST(root.Right)
}

/**
 * 递归法二：中序遍历过程中统计众数
 */
func findMode(root *TreeNode) []int {
	var res []int

	var pre *TreeNode          // 指向前一结点
	var curCount, maxCount int // 分别记录当前遍历结点的出现次数以及最大出现次数

	// 中序遍历过程，二叉搜索树的中序遍历相当于再访问有序表
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		//if pre != nil {
		//	if node.Val != pre.Val {
		//		// 当前结点的值发生改变时，可能需要更新众数相关结果
		//		if curCount > maxCount {
		//			res = res[:0]
		//			res = append(res, pre.Val)
		//			maxCount = curCount
		//		} else if curCount == maxCount {
		//			// 众数不止一个
		//			res = append(res, pre.Val)
		//		}
		//		// 重新统计新节点值
		//		curCount = 0
		//	}
		//}
		//curCount++
		//// 更新上一访问结点
		//pre = node

		// 对注释部分进行改进(改进在：现在每个结点都处理自己的计数并判断是否更新自己的信息到结果中，
		// 而之前是处理自己的计数并在结点值发生改变时判断上一结点的信息是否更新到结果中，这样会造成最后一种结点的信息无法更新，因为节点值不会再改变)
		// 可以避免代码重复（下方注释部分可省去）
		// 先处理当前结点值的计数
		if pre != nil && node.Val == pre.Val {
			curCount++
		} else {
			curCount = 1
		}
		// 再将当前值的计数与当前众数比较，进而更新众数
		if curCount == maxCount {
			// 众数不止一个
			res = append(res, node.Val)
		}
		if curCount > maxCount {
			// 结果与众数更新为当前结点的信息
			res = res[:0] // 清空已有结果列表
			res = append(res, node.Val)
			maxCount = curCount
		}
		pre = node
		travel(node.Right)
	}

	travel(root)
	//// 此时需要处理一个特殊情况，中序遍历过程中最后一种结点值（最大值Val）不会更新到结果中
	//if curCount > maxCount {
	//	res = res[:0]
	//	res = append(res, pre.Val)
	//} else if curCount == maxCount {
	//	res = append(res, pre.Val)
	//}

	return res
}

/**
 * 迭代法也是可以的，仍然采用中序遍历模板，与上述递归法在核心处理上是完全一直的，区别就是遍历方式的不同，故不再赘述
 * 实际上中序遍历迭代法不满足题目进阶要求（不使用额外空间），因为会用到栈
 */
