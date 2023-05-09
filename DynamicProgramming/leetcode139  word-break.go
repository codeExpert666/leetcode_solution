package DynamicProgramming

/**
 * 完全背包问题应用（dp）
 * 这里dp[i]表示字符串s[:i]能否由当前单词表构成
 * 这里注意三点：1、完全背包 2、排列问题 3、递推公式较为特殊，有一个判断尾部的过程
 */
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1) // 构建dp数组及初始化
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		dp[i] = false
	}
	// 注意遍历顺序，因为是排列问题
	for j := 0; j < len(dp); j++ {
		for i := 0; i < len(wordDict); i++ {
			if j >= len(wordDict[i]) {
				// 递推公式，注意需要判断当前单词是否是当前字符串的尾部
				if !dp[j] && isTail(s[:j], wordDict[i]) && dp[j-len(wordDict[i])] {
					dp[j] = true
				}
			}
		}
	}
	return dp[len(s)]
}

// 判断字符串t是否是s的结尾
func isTail(s, t string) bool {
	d := len(s) - len(t)
	for i := d; i < len(s); i++ {
		if s[i] != t[i-d] {
			return false
		}
	}
	return true
}

/**
 * 与上述方法思路完全一致，但实现细节有所不同，比上一方法更好
 */
func wordBreak1(s string, wordDict []string) bool {
	// 将单词表转换成了map结构，从而实现了上一种方法的isTail功能
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	// dp初始化
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// 遍历（完全背包+排列）
	for i := 1; i <= len(s); i++ {
		// 这里是与上一种方法的不同之处
		// 这种不同主要体现在对物品（单词）的遍历顺序不同：上一种方法就是无脑按照输入所给的顺序进行遍历
		// 而本方法则是有选择有条理的进行遍历：直接按照当前字符串的可能结尾进行遍历（无需对全部单词遍历），从而省去了上一方法isTail的时间复杂度
		// 由于本身dp[i]就为false，故判断处dp[i]为false时便不再处理
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
