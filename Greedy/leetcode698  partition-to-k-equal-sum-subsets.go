package Greedy

import "sort"

// CanPartitionKSubsets 回溯+记忆搜索，高难度题目，说实话，即使看了题解，我现在仍然是有点模糊的状态，再让我写一次我一定还是写不出来
// 所以了解即可，考到算我倒霉，官解很笼统，参考这个解答：https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/solutions/2892519/javapython3cpai-xu-ji-yi-hua-sou-suo-hui-glsu/
// 代码读懂程度 95%，但一定写不出来
func CanPartitionKSubsets(nums []int, k int) bool {
	all := 0
	for _, v := range nums {
		all += v
	}
	if all%k > 0 {
		return false
	}
	per := all / k
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] > per {
		return false
	}

	dp := make([]bool, 1<<n)
	for i := range dp {
		dp[i] = true
	}
	var dfs func(int, int) bool
	dfs = func(s, p int) bool {
		if s == 0 {
			return true
		}
		if !dp[s] {
			return dp[s]
		}
		dp[s] = false
		for i, num := range nums {
			if num+p > per {
				break
			}
			if s>>i&1 > 0 && dfs(s^1<<i, (p+nums[i])%per) {
				return true
			}
		}
		return false
	}
	return dfs(1<<n-1, 0)
}
