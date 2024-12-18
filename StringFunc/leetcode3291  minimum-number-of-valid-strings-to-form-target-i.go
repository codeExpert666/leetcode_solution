package StringFunc

import "math"

// 字典树数据结构
// 注意理解这个存储结构，一个节点所对应的字符实际上存储在父节点的孩子列表中
type TireNode struct {
	children [26]*TireNode // 存储子节点
	isEnd    bool          // 标记单词结束
}

func NewTrieNode() *TireNode {
	return &TireNode{
		children: [26]*TireNode{},
		isEnd:    false,
	}
}

// 字典树+动态规划
// 字典树(Trie)是一种树形数据结构,主要用于高效地存储和检索字符串数据集。
// 特点是共享前缀的字符串可以共享存储空间
func MinValidStrings(words []string, target string) int {
	// 遍历单词集合，构建字典树
	root := NewTrieNode()
	for _, w := range words {
		cur := root
		for i := 0; i < len(w); i++ {
			charIndex := w[i] - 'a'
			if cur.children[charIndex] == nil { // 字典树中没有 w[i] 字符
				cur.children[charIndex] = NewTrieNode()
			}
			cur = cur.children[charIndex] // 指向 w[i] 字符
		}
		cur.isEnd = true // 单词 w 到达末尾
	}
	// 动态规划：dp[i] 表示 target[i:] 的最少拼接次数
	l := len(target)
	dp := make([]int, l+1)
	// dp 数组初始化
	for i := 0; i <= l; i++ {
		dp[i] = math.MaxInt
	}
	dp[l] = 0
	// 倒序状态转移
	for i := l - 1; i >= 0; i-- {
		cur := root
		for j := i; j <= l-1; j++ {
			charIndex := target[j] - 'a'
			if cur.children[charIndex] == nil {
				break // 组成 target[i：] 的第一个前缀最大只能是 target[i:j]
			}
			if dp[j+1] != math.MaxInt { // target[i：] 除去前缀 target[i:j] 的部分有拼接方式
				dp[i] = min(dp[i], dp[j+1]+1)
			}
			cur = cur.children[charIndex]
		}
	}
	// 输出结果
	if dp[0] == math.MaxInt {
		return -1
	}
	return dp[0]
}
