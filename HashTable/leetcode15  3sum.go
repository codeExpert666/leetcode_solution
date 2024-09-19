package HashTable

import "sort"

// hash法相当困难，主要是因为题目要求三元组不可重复，同时至少有两层循环嵌套
// 去重与剪枝操作及其繁琐，hash表的使用本身就给剪枝操作带来极大困难
// 面试很难把所有情况考虑到位，所有换用双指针法
// 先对数组排序，这样可以通过下标判断大小，较为方便
// 反正是二重循环，排序的复杂度又是O(nlogn)，不影响
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 先对数组进行排序
	res := make([][]int, 0)

	// 固定第一个数的位置，剩余两个数利用双指针法寻找
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // 这种情况不可能再找到解
			break
		}
		if i == 0 || nums[i] != nums[i-1] { // 对第一个数去重, 这一步有一个细节：不能写成nums[i] != nums[i + 1]，这是因为可能有个解包含两个相同数。一定是先找解再去重
			left, right := i+1, len(nums)-1
			for left < right {
				l, r := nums[left], nums[right]
				sum := nums[i] + l + r
				if sum < 0 {
					left++
				} else if sum > 0 {
					right--
				} else {
					res = append(res, []int{nums[i], l, r})
					// 对第二个数和第三个数的去重应放在找到解之后
					// 对第二个数去重
					for left < right && nums[left] == l { // 这里left < right很重要，可以防止数组越界
						left++
					}
					// 对第三个数去重
					for left < right && nums[right] == r { // 同上
						right--
					}
				}
			}
		}
	}
	return res
}
