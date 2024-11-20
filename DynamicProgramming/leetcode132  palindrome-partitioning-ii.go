package DynamicProgramming

import "math"

// 判断输入字符串是否为回文串
func isSymmetric(s string) bool {
	for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
		if s[low] != s[high] {
			return false
		}
	}
	return true
}

// 时间复杂度较高，O(n^3)
func minCut(s string) int {
	// dp[i][j]表示s[i:j+1]分解成回文串的最少切割次数
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- { // 注意遍历顺序
		for j := i; j < len(s); j++ {
			if j != i && !isSymmetric(s[i:j+1]) { // 当前串不是回文串时才需要切割
				dp[i][j] = j - i              // 最差切割次数
				for k := i + 1; k <= j; k++ { // 有j-i个切割位置，逐个尝试
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k][j]+1)
				}
			}
		}
	}
	return dp[0][len(s)-1]
}

// 优化思路：减少状态数，将状态降为一维，dp[i]表示s[:i]分解为回文串的最少切割次数
// 对于串s[:i]，倘若自身不是回文串，则需要切割，共有i-1个切割位置，对于切割位置j<i
// 倘若s[j:i]是回文串，则容易处理，符合动态规划思路；如若s[j:i]不是回文串，无法根据已有状态推导
// 但需要注意的是，这种情况完全不必处理，因为，如果s[j:i]不是回文串，则其需要分割成回文串，那么关注其最右侧的回文串，
// 发现这种情况已经考虑过，这就是我上面的方法疏忽的一点。
// 其次，判断是否为回文串这个方法可以提前利用dp方法求出所有子串是否为回文串，这个题目之前做过。这样对于主算法中涉及的
// 回文串判断，就可以直接判断查看dp数组结果。这种方法相当于将复杂度分离了出来，这样，判断所有子串是否为回文串是O(n^2)
// 主算法也是O(n^2)，总体为O(n^2)
func minCut1(s string) int {
	isValid := make([][]bool, len(s))
	for i := 0; i < len(isValid); i++ {
		isValid[i] = make([]bool, len(s))
		isValid[i][i] = true
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && (isValid[i+1][j-1] || j-i == 1) {
				isValid[i][j] = true
			}
		}
	}

	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len(s); i++ {
		if isValid[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isValid[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(s)-1]
}
