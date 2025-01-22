package Array

import "math"

// 前后缀分解 + 二维 0-1 背包 + 优化所选元素个数 + 试填法
// 以往的动态规划用的都是查表法，也即根据转移状态计算状态。
// 当难以获取转移来源时，可以考虑刷表法，也即，用当前状态更新其他状态。
// 本题难度较大，以下解法来自于灵神题解。并未理解。
const bitWidth = 7
const mx = 1 << bitWidth

func maxValue(nums []int, k int) (ans int) {
	n := len(nums)
	k2 := min(k, bitWidth) // 至多选 k2 个数
	suf := make([][mx]bool, n-k+1)
	f := make([][mx]bool, k2+1)
	f[0][0] = true
	maxI := [mx]int{}
	cnt := [mx]int{}
	for i := n - 1; i >= k; i-- {
		v := nums[i]
		for j := min(k2-1, n-1-i); j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		if i <= n-k {
			suf[i] = f[k2]
		}
		// 枚举 v 的超集
		for s := v; s < mx; s = (s + 1) | v {
			cnt[s]++
			if cnt[s] == k {
				// 从 n-1 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
				maxI[s] = i
			}
		}
	}

	pre := make([][mx]bool, k2+1)
	pre[0][0] = true
	minI := [mx]int{}
	for i := range minI {
		minI[i] = math.MaxInt
	}
	cnt = [mx]int{}
	for i, v := range nums[:n-k] {
		for j := min(k2-1, i); j >= 0; j-- {
			for x, hasX := range pre[j] {
				if hasX {
					pre[j+1][x|v] = true
				}
			}
		}
		// 枚举 v 的超集
		for s := v; s < mx; s = (s + 1) | v {
			cnt[s]++
			if cnt[s] == k {
				// 从 0 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
				minI[s] = i
			}
		}
		if i < k-1 {
			continue
		}
		a := []int{}
		b := []int{}
		for x, has := range pre[k2] {
			if has && minI[x] <= i {
				a = append(a, x)
			}
			if suf[i+1][x] && maxI[x] > i {
				b = append(b, x)
			}
		}
		ans = max(ans, findMaximumXOR(a, b))
		if ans == mx-1 {
			return
		}
	}
	return
}

// 421. 数组中两个数的最大异或值
// 改成两个数组的最大异或值，做法是类似的，仍然可以用【试填法】解决
func findMaximumXOR(a, b []int) (ans int) {
	mask := 0
	for i := bitWidth - 1; i >= 0; i-- { // 从最高位开始枚举
		mask |= 1 << i
		newAns := ans | 1<<i // 这个比特位可以是 1 吗？
		seen := [mx]bool{}
		for _, x := range a {
			seen[x&mask] = true // 低于 i 的比特位置为 0
		}
		for _, x := range b {
			x &= mask // 低于 i 的比特位置为 0
			if seen[newAns^x] {
				ans = newAns
				break
			}
		}
	}
	return
}
