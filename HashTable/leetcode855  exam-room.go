package HashTable

// 总结：如果在解题过程中需要 1、操作优先级最高元素 2、操作优先级非最高的特定元素（查、删、改）。那么必然选择红黑树，
// 因为红黑树是有序集合，其对任意元素的任意操作复杂度都是 log(n)
// 如果说只需要 1 功能，那么优先队列和红黑树都可以，且复杂度相当
// 所以，总的来说，无脑红黑树就完事了，肯定不差。

import "github.com/emirpasic/gods/trees/redblacktree"

// ExamRoom 考场座位分配结构
type ExamRoom struct {
	rbt   *redblacktree.Tree // 红黑树，用于维护座位区间
	left  map[int]int        // 记录每个就坐座位左边的就坐座位
	right map[int]int        // 记录每个就坐座位右边的就坐座位
	n     int                // 考场座位总数
}

func Constructor(n int) ExamRoom {
	// dist 计算区间的优先级（也即学生在该区间内就坐后，与离他最近的人之间的最大距离）
	dist := func(s []int) int {
		// 如果是边界区间（最左或最右），直接返回区间长度-1
		if s[0] == -1 || s[1] == n {
			return s[1] - s[0] - 1
		}
		// 对于中间区间，返回区间长度的一半
		return (s[1] - s[0]) >> 1
	}

	// cmp 比较两个区间的优先级
	cmp := func(a, b interface{}) int {
		x, y := a.([]int), b.([]int)
		d1, d2 := dist(x), dist(y)
		// 如果距离相等，选择左边界较小的
		if d1 == d2 {
			return x[0] - y[0]
		}
		// 否则选择距离较大的
		return d2 - d1
	}

	// 初始化考场，创建一个包含所有座位的区间 [-1, n]
	er := ExamRoom{redblacktree.NewWith(cmp), map[int]int{}, map[int]int{}, n}
	er.add([]int{-1, n})
	return er
}

// Seat 分配一个座位
func (er *ExamRoom) Seat() int {
	// 获取优先级最高的区间
	s := er.rbt.Left().Key.([]int)
	// 计算新座位位置（默认在区间中点）
	p := (s[0] + s[1]) >> 1
	// 特殊处理最左边界
	if s[0] == -1 {
		p = 0
		// 特殊处理最右边界
	} else if s[1] == er.n {
		p = er.n - 1
	}
	// 删除原区间，添加两个新区间
	er.del(s)
	er.add([]int{s[0], p})
	er.add([]int{p, s[1]})
	return p
}

// Leave 删除一个座位
func (er *ExamRoom) Leave(p int) {
	// 获取座位的左右边界
	l := er.left[p]
	r := er.right[p]
	// 删除包含p的两个区间
	er.del([]int{l, p})
	er.del([]int{p, r})
	// 添加合并后的新区间
	er.add([]int{l, r})
}

// add 添加一个区间
func (er *ExamRoom) add(s []int) {
	er.rbt.Put(s, struct{}{})
	er.left[s[1]] = s[0]  // 记录右端点的左边界
	er.right[s[0]] = s[1] // 记录左端点的右边界
}

// del 删除一个区间
func (er *ExamRoom) del(s []int) {
	er.rbt.Remove(s)
	// 两个 delete 可以省略，因为本题中 del 操作一定会配有 add 操作，
	// 而 add 操作中对边界的重新记录会覆盖原本记录，所以无需删除。
	// delete(er.left, s[1])  // 删除右端点的左边界记录
	// delete(er.right, s[0]) // 删除左端点的右边界记录
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
