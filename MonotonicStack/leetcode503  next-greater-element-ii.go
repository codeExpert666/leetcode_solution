package MonotonicStack

func NextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	stack := make([]int, 0)
	// 考虑到nums是循环数组，故遍历两次，相当于将两个nums拼接在一起
	for j := 0; j < 2; j++ {
		for i, v := range nums {
			for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
				res[stack[len(stack)-1]] = v
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	return res
}
