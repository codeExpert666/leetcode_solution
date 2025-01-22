package weeklycompetition

import (
	"sort"
)

var mod = int(1e9 + 7)
var pow []int

func minMaxSums(nums []int, k int) int {
	n := len(nums)
	// 先对数组排序
	sort.Ints(nums)

	// Precompute powers of 2
	pow = make([]int, n+1)
	pow[0] = 1
	for i := 1; i <= n; i++ {
		pow[i] = pow[i-1] * 2 % mod
	}

	var res int
	// 每个元素既能作为最小值，也能作为最大值
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			res += ((nums[i] + nums[j]) * numOfAppearances(i, j, k)) % mod
		}
	}

	return res
}

func numOfAppearances(i, j, k int) int {
	if i == j {
		return 1
	} else if j-i+1 <= k {
		return pow[j-i-1]
	} else {
		return combination(j-i-1, k-2)
	}
}

func combination(n, k int) int {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	return (combination(n-1, k-1) + combination(n-1, k)) % mod
}
