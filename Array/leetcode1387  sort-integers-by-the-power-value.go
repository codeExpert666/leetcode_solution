package Array

import "slices"

var memo = make(map[int]int) // 记忆化搜索

func getWeight(n int) int {
	if n == 1 {
		return 0
	}
	if w, ok := memo[n]; ok {
		return w
	}
	if n%2 == 0 {
		memo[n] = getWeight(n>>1) + 1
	} else {
		memo[n] = getWeight(3*n+1) + 1
	}
	return memo[n]
}

func GetKth(lo int, hi int, k int) int {
	l := hi - lo + 1
	array := make([]int, l)
	for i := 0; i < l; i++ {
		array[i] = lo + i
	}
	slices.SortFunc(array, func(a, b int) int {
		aw, bw := getWeight(a), getWeight(b)
		if aw == bw {
			return a - b
		}
		return aw - bw
	})
	return array[k-1]
}
