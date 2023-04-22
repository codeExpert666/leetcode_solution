package Greedy

/**
 * 贪心法，自己实现的版本，代码有冗余，但逻辑清晰，时间复杂度O(n)，空间复杂度O(1)
 */
func candy(ratings []int) int {
	var res int
	pre := 1 // 前一孩子(目前指向第一个孩子)分配到的糖果数量，每个孩子至少有一个糖果
	res += pre

	// 评分递减时需要特殊处理
	firstCandy := 1 // 下降序列第一个孩子分到的糖果
	count := 1      // 下降序列中的孩子个数，一个孩子也是下降序列，故初始化为1

	for i := 1; i < len(ratings); i++ { // 从第二个孩子开始遍历
		if ratings[i] < ratings[i-1] { // 当前孩子评分低于前一孩子
			count++
		} else {
			// 先处理前面的递减序列
			if count > 1 { // 判断是否存在递减序列
				if firstCandy < count {
					res -= firstCandy
					res += (1 + count) * count / 2
				} else {
					res += count * (count - 1) / 2
				}
				pre = 1
			}
			// 处理当前相邻两个孩子
			if ratings[i] > ratings[i-1] { // 当前孩子评分高于前一孩子
				pre += 1 // 多分一个糖
			} else { // 当前孩子评分等于前一孩子
				pre = 1 // 分最少的糖
			}
			res += pre // 更新总分配糖数
			// 更新递减序列信息
			count = 1
			firstCandy = pre
		}
	}
	// 若评分数组末尾出现递减序列，则上述循环无法处理
	// 需单独处理
	if count > 1 {
		if firstCandy < count {
			res -= firstCandy
			res += (1 + count) * count / 2
		} else {
			res += count * (count - 1) / 2
		}
	}
	return res
}
