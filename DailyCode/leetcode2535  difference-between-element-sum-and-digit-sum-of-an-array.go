package DailyCode

func DifferenceOfSum(nums []int) int {
	var elementSum, numSum int
	for _, num := range nums {
		elementSum += num // 元素和
		for num > 0 {     // 数字和
			numSum += num % 10
			num /= 10
		}
	}
	// 考虑到正整数一定大于等于其各个位上的数字和，所以无需判断
	// if elementSum > numSum {
	// 	return elementSum - numSum
	// } else {
	// 	return numSum - elementSum
	// }
	return elementSum - numSum
}
