package DailyCode

// 每一对相邻元素，如果出现降序或相等，如nums[i] >= nums[i+1]
// 那么一定范围内（以nums[j],i+2-k <= j <= i为首）的子数组能量值均为-1
// 降序如果比较接近，范围重叠，可通过记录上一范围的max(j)进行优化
// 我的方法时间也是O(n)，但本质上是O(2n-k)
func ResultsArray(nums []int, k int) []int {
	res := append([]int{}, nums[k-1:]...) // 初始值设置为不存在降序的情况
	bound := 0                            // 降序所影响数组的首元素最小位置
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] { // 降序或相等
			minj, maxj := max(i+2-k, bound), min(i, len(nums)-k) // 范围重叠优化
			for j := minj; j <= maxj; j++ {
				res[j] = -1 // 范围内子数组不可能递增
			}
			bound = i + 1 // 更新首元素最小位置
		}
	}
	return res
}

// 根据题意可知，由于子数组如果满足连续上升，此时相邻元素的差值一定为 1，
// 此时我们在遍历数组的同时，用一个计数器 cnt 统计以当前索引 i 为结尾时连续上升的元素个数，初始化 cnt=0
// 该方法更好，时间是真正的O(n)
func ResultsArray1(nums []int, k int) []int {
	n := len(nums)
	cnt := 0
	ans := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		if i == 0 || nums[i]-nums[i-1] != 1 {
			cnt = 1
		} else {
			cnt++
		}
		if cnt >= k {
			ans[i-k+1] = nums[i]
		} else if i-k+1 >= 0 {
			ans[i-k+1] = -1
		}
	}
	return ans
}
