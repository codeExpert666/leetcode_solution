package DynamicProgramming

// 数位DP适用于从一系列连续的内容（内容一般由一系列元素组成，元素按规律变化形成不同内容）中找出符合条件的内容，尤其适合进制以及字典顺序所构成的连续数字或字符串
// 数位DP的难点在于 DP 状态（DFS 参数）的定义，我总结如下规律：
// 要找到一系列内容中符合条件的内容，就将 dfs 定义为：前 i 个元素构成的内容在（部分）满足条件的情况下，后续能构造出的符合条件的完整内容个数

// next[i]的值本质上就是 s[:i] 的最长公共前后缀的长度
func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1 // 方便后续比较时的判断
	for i := 1; i < len(next); i++ {
		for j := i - 1; j > 0; j = next[j] {
			if s[i-1] == s[next[j]] {
				next[i] = next[j] + 1
				break
			}
		}
	}
	return next
}

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	m := len(evil)
	mod := int(1e9 + 7)
	evilNext := getNext(evil) // KMP next数组
	// dp 是对递归状态的记忆，避免重复搜索。理论上，需要记忆 dfs 四个参数以及他们对应的结果
	// 但由于 lowLimit 和 upLimit 一旦为 true，前 i 个字符便完全确定，这样的状态不可能再经过第二次，故无需记录
	// 保留 i 与 cmp 状态以及其对应结果即可
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ { // 初始化
			dp[i][j] = -1
		}
	}
	// dfs 函数返回前 i 个字母不包含 evil 串时能够构造出的不包含 evil 的字符串数量
	// i：匹配串当前下标；cmp：模式串当前待比较下标
	// lowLimit：前 i 个字母是否到达字典序下限；upLimit：前 i 个字母是否到达字典序上限
	var dfs func(int, int, bool, bool) int
	dfs = func(i, cmp int, lowLimit, upLimit bool) (res int) {
		// 递归退出条件
		if i == n { // 无法再添加字符，能构造出的不含 evil 的字符串只有 s[:i] 一个
			return 1
		}
		if !lowLimit && !upLimit { // 记录的状态只对应于前面字符未达上下限的情况
			p := &dp[i][cmp]
			if *p != -1 { // 该状态已被记录
				return *p
			}
			defer func() { *p = res }()
		}
		// 获取当前位字符上下限
		var low, up byte = 'a', 'z'
		if lowLimit {
			low = s1[i]
		}
		if upLimit {
			up = s2[i]
		}
		for c := low; c <= up; c++ {
			// KMP 匹配
			tmp := cmp
			for tmp >= 0 && c != evil[tmp] {
				tmp = evilNext[tmp]
			}
			tmp++         // 指向模式串下一待比较字符
			if tmp == m { // 模式串匹配完毕，说明当前主串已经包含 evil，无需继续递归
				continue
			}
			res = (res + dfs(i+1, tmp, lowLimit && c == s1[i], upLimit && c == s2[i])) % mod
		}
		return
	}
	return dfs(0, 0, true, true) % mod
}
