package DynamicProgramming

/**
 * 法一：动态规划 O(n^2)   注意：子序列问题是动态规划解决的经典问题！！！！非常重要
 * 这里就一个难点：dp数组的含义，搞清楚这个，后续步骤都不是问题
 * 这里其实还是可以从“状态”去考虑：一个序列的递增子序列有哪些可能呢？无非就是以该序列的各个元素结尾
 * 于是这里就形成了n个状态，且状态之间存在递推关系。这里没有直接求解问题，而是通过对问题转换，间接在过程中解决了问题
 * !!!综上，dp[i]表示i之前包括i的以nums[i]结尾的最长递增子序列的长度!!!
 */
func lengthOfLIS(nums []int) int {
	var res = 1 // 子序列长度至少为1，而且这样初始化可以处理nums长度为1的情况
	// dp初始化
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		// 每个序列的递增子序列长度至少为1，因为一个数就是递增子序列
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// 更新输入的最长子序列长度
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

/**
 * 法二：二分法 O(n*log(n)) 非常巧妙的方法，欣赏即可
 * 具体做法就是：新建一个数组，然后第一个数先放进去，然后第二个数和第一个数比较，如果说大于第一个数，那么就接在他后面，
 * 如果小于第一个数，那么就替换，一般的，如果有i个数，那么每进来一个新的数，都要用二分查找法来得知要替换在哪个位置的数。
 * 因为有个for循环，所以是O(N),在加上循环里有个二分查找，所以最后是O(NlogN)的时间复杂度。
 */
func lengthOfLIS1(nums []int) int {
	// 存储当前最长递增子序列（只有长度是有效信息，里面存储的并不一定是真正的最长递增子序列）
	lis := make([]int, len(nums))
	// 初始化
	lis[0] = nums[0]
	var l = 1 // 记录子序列真实长度

	for i := 1; i < len(nums); i++ {
		if nums[i] > lis[l-1] {
			// 比当前递增子序列尾部元素大，加入子序列
			lis[l] = nums[i]
			l++
		} else {
			// 二分查找应插入的位置
			low, high := 0, l-1
			for low <= high {
				mid := low + (high-low)>>1
				if nums[i] <= lis[mid] {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
			// low为待插入位置，这里实际不用插入，替换low位置处的数即可，所以l不变
			// 这里有个小细节，二分查找时，low与high均可以通过上面if语句的细节调整进而作为待插入位置
			// 但这里用low作为插入位置更好，因为用high作为插入位置会出现high = -1的情况，从而出现数组越界错误
			lis[low] = nums[i]
		}
	}
	return l
}
