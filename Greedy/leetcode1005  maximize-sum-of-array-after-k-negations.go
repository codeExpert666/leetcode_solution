package Greedy

import (
	"math"
	"sort"
)

/**
 * 先对数组进行排序（按绝对值大小排序），再利用贪心的思想：尽量使得绝对值大的数均调整为正数
 * 本题难在情况众多，容易考虑不全
 */
func largestSumAfterKNegations(nums []int, k int) int {
	// 先对数组进行排序
	sort.Ints(nums)
	// 遍历数组
	var sum int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			// 当k有剩余并且存在负数时，翻转一次该负数
			nums[i] = -nums[i]
			k--
		} else if nums[i] == 0 {
			// 遇到0，则可把所有的翻转次数用在0上
			k = 0
		} else if nums[i] > 0 && k > 0 {
			// 当k有剩余并且遍历到正数时，需要找到最小的正数进行翻转
			// 最小的正数只可能是其本身或其前一数字（如果存在的话，并且该正数是由负数翻转的）
			// 寻找最小正数
			if i > 0 && nums[i-1] <= nums[i] {
				// 判断k的奇偶性
				if k%2 == 1 {
					sum -= 2 * nums[i-1]
				}
			} else {
				// 判断k的奇偶性
				if k%2 == 1 {
					nums[i] = -nums[i]
				}
			}
			k = 0
		}
		sum += nums[i]
	}
	// 数组遍历完毕，但翻转次数仍然存在，这种情况只可能是数组全负，且翻转次数大于数组长度
	// 此时只需处理数组最后一个数
	if k%2 == 1 {
		sum -= 2 * nums[len(nums)-1]
	}
	return sum
}

/**
 * 贪心法二：
 * 第一步：将数组按照绝对值大小从大到小排序，注意要按照绝对值的大小
 * 第二步：从前向后遍历，遇到负数将其变为正数，同时K--
 * 第三步：如果K还大于0，那么反复转变数值最小的元素，将K用完
 * 第四步：求和
 * 这种写法较为简洁，无需像第一种方法一样考虑众多情况，不易出错，是更好的方法
 */
func largestSumAfterKNegations1(nums []int, K int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if K > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			K--
		}
	}

	if K%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}
