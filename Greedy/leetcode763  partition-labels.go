package Greedy

/**
 * 在遍历的过程中相当于是要找每一个字母的边界，如果找到之前遍历过的所有字母的最远边界，说明这个边界就是分割点了。
 * 此时前面出现过所有字母，最远也就到这个边界了。
 * 可以分为如下两步：
 * 1、统计每一个字符最后出现的位置
 * 2、从头遍历字符，并更新字符的最远出现下标，如果找到字符最远出现位置下标和当前下标相等了，则找到了分割点
 * 这里的贪心应该主要体现在不断扩大最远下标
 */
func partitionLabels(s string) []int {
	// 利用hash表统计每个字符最后出现的位置
	m := make(map[byte]int) // 可以用长度为26的切片代替
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	var res []int
	var maxIndex int // 记录最远下标
	preIndex := -1   // 前一分割点
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 更新最远下标
		if m[s[i]] > maxIndex {
			maxIndex = m[s[i]]
		}
		// 判断是否到达分割点
		if i == maxIndex {
			// i是分割下标，需要转换为分段长度
			res = append(res, i-preIndex)
			preIndex = i
		}
	}
	return res
}
