package Array

// 利用栈累积错误位置，每当遇到不同类错误便消除栈顶元素
func SortArrayByParityII(nums []int) []int {
	var stack []int  // 栈中存储错误位置
	var errType bool // 标记栈中错误类型，false表示奇数位存偶数，true表示偶数位存奇数
	for i, v := range nums {
		if i%2 == 1 && v%2 == 0 { // 奇数位存偶数
			if !errType || len(stack) == 0 {
				stack = append(stack, i)
				errType = false
			} else {
				// 不同类型错误位置交换元素
				nums[stack[len(stack)-1]], nums[i] = v, nums[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
			}
		} else { // 偶数位存奇数
			if i%2 == 0 && v%2 == 1 {
				if errType || len(stack) == 0 {
					stack = append(stack, i)
					errType = true
				} else {
					// 不同类型错误位置交换元素
					nums[stack[len(stack)-1]], nums[i] = v, nums[stack[len(stack)-1]]
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	return nums
}

// 双指针法，无需栈，两个指针分别指向偶数位和奇数位
// 偶数位遇到奇数，奇数位指针移动，反之亦然
func SortArrayByParityII1(nums []int) []int {
	for i, j := 0, 1; i < len(nums); {
		if nums[i]%2 == 0 {
			i += 2
		} else if nums[j]%2 == 1 {
			j += 2
		} else {
			nums[i], nums[j] = nums[j], nums[i] // 交换
			i += 2
		}
	}
	return nums
}
