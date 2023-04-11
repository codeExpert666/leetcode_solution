package Greedy

/**
 * 一次遍历，不断判断是否可以扩充子序列
 * 遍历序列：会遇到以下两种情况：1、差值正负交替，这就是所求序列，增加当前子序列长度即可
 * 2、差值不交替（正正、负负、正0、负0），保持当前子序列长度不变
 * 此方法不断寻找当前所遍历序列的最长交替子序列，也可以将其理解为贪心的思想
 */
func wiggleMaxLength(nums []int) int {
	res := 1
	var flag int // 用于判断差值是否交替，记录的是前一组有效数据的差值
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if d == 0 {
			// 遇到差值为0，直接跳过
			continue
		}
		//tmp := d * flag
		// 乘法复杂度高，改用正负判断
		if flag == 0 || (flag < 0 && d > 0) || (flag > 0 && d < 0) {
			// 两种情况：1、差值交替 2、第一次遇到差值不为零的一组数据
			res++
			flag = d
		}
	}
	return res
}
