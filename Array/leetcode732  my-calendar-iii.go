package Array

const maxTime = int(1e9)

type MyCalendarThree struct {
	root *node
}

func ConstructorMCTH() MyCalendarThree {
	return MyCalendarThree{
		root: &node{},
	}
}

func (mc *MyCalendarThree) Book(startTime int, endTime int) int {
	update(mc.root, 0, maxTime, startTime, endTime-1, 1)
	return mc.root.val
}

// 一个节点代表一个子区间
type node struct {
	left, right *node // 区间再分为左右子区间
	val         int   // 该区间中的最大预定次数
	add         int   // 懒惰标记（该区间最近被预定的次数）
}

// 向线段树中更新指定区间
func update(root *node, start, end, l, r, val int) {
	// 当前节点（区间）在输入日程范围内
	if start >= l && end <= r {
		root.val += val
		root.add += val
		return
	}
	// 向下查找子节点，直到找到符合日程范围的节点进行更新
	// 左节点对应区间为[start:mid+1]，右节点对应区间为[mid+1:end]
	mid := start + (end-start)>>1
	pushDown(root)
	if l <= mid { // 左节点对应区间有部分需要更新
		update(root.left, start, mid, l, r, val)
	}
	if r > mid { // 右节点对应区间有部分需要更新
		update(root.right, mid+1, end, l, r, val)
	}
	// 左右子区间更新完毕，向上汇总
	pushUp(root)
}

// 从线段树中查询指定区间
// func query(root *node, start, end, l, r int) int {
// 	// 整体逻辑与 update 操作相似
// 	if start >= l && end <= r {
// 		return root.val
// 	}
// 	mid := start + (end-start)>>1
// 	pushDown(root)
// 	var res int
// 	if l <= mid {
// 		res = query(root.left, start, mid, l, r)
// 	}
// 	if r > mid {
// 		res = max(res, query(root.right, mid+1, end, l, r))
// 	}
// 	return res
// }

func pushDown(root *node) {
	// 没有子节点先进行分区（因为每次都是二分，分区方案是固定的，所以不用记录区间端点）
	if root.left == nil {
		root.left = &node{}
	}
	if root.right == nil {
		root.right = &node{}
	}
	// 未累积懒惰操作
	if root.add == 0 {
		return
	}
	// 设置子节点值
	root.left.val += root.add
	root.left.add += root.add
	root.right.val += root.add
	root.right.add += root.add
	// 推送完毕，懒惰标记重置
	root.add = 0
}

func pushUp(root *node) {
	root.val = max(root.left.val, root.right.val)
}
