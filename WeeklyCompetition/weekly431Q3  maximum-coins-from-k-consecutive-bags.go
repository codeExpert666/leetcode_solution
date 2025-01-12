package weeklycompetition

import (
	"slices"
)

// 线段树超时
func MaximumCoins(coins [][]int, k int) int64 {
	n := int(1e9)
	s, e := int(1e9), 1 // 有硬币的袋子范围（两个边界）

	// 构建线段树
	root := &node{}
	for _, interval := range coins {
		s = min(s, interval[0])
		e = max(e, interval[1])
		update(root, 1, n, interval[0], interval[1], interval[2])
	}

	// 查询线段树
	var res int
	for i := s; ; i++ {
		res = max(res, query(root, 1, n, i, i+k-1))
		if i+k-1 > e {
			break
		}
	}

	return int64(res)
}

// 线段树
type node struct {
	left, right *node
	val, add    int
}

func update(root *node, start, end, l, r, val int) {
	if start >= l && end <= r {
		root.val += (end - start + 1) * val
		root.add = val
		return
	}
	mid := start + (end-start)>>1
	pushDown(root, mid-start+1, end-mid)
	if l <= mid {
		update(root.left, start, mid, l, r, val)
	}
	if r > mid {
		update(root.right, mid+1, end, l, r, val)
	}
	pushUp(root)
}

func query(root *node, start, end, l, r int) int {
	if start >= l && end <= r {
		return root.val
	}
	mid := start + (end-start)>>1
	pushDown(root, mid-start+1, end-mid)
	var res int
	if l <= mid {
		res += query(root.left, start, mid, l, r)
	}
	if r > mid {
		res += query(root.right, mid+1, end, l, r)
	}
	return res
}

func pushUp(root *node) {
	root.val = root.left.val + root.right.val
}

func pushDown(root *node, leftNum, rightNum int) {
	if root.left == nil {
		root.left = &node{}
	}
	if root.right == nil {
		root.right = &node{}
	}
	if root.add == 0 {
		return
	}
	root.left.val += leftNum * root.add
	root.right.val += rightNum * root.add
	root.left.add = root.add
	root.right.add = root.add
	root.add = 0
}

// 前缀和，超出内存限制
func maximumCoins(coins [][]int, k int) int64 {
	slices.SortFunc(coins, func(a, b []int) int { return a[0] - b[0] })

	// 构建前缀和数组
	prefixSum := make([]int, 0)
	sum := 0 // 当前遍历到的袋子的钱币和
	j := 0   // 指向当前遍历到的 coins 区间
	for i := coins[0][0]; i <= coins[len(coins)-1][1]; i++ {
		if i < coins[j][0] {
			prefixSum = append(prefixSum, sum)
		} else {
			sum += coins[j][2]
			prefixSum = append(prefixSum, sum)
			if i == coins[j][1] {
				j++
				if j >= len(coins) {
					break
				}
			}
		}
	}

	// 计算连续 k 个袋子的最大硬币数量
	l := len(prefixSum)
	if l <= k {
		return int64(prefixSum[l-1])
	}

	var res int
	for i := l - 1; i-k >= 0; i-- {
		res = max(res, prefixSum[i]-prefixSum[i-k])
	}
	return int64(max(res, prefixSum[k-1]))
}
