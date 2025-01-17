package Simulation

func minOperations(nums []int, k int) int {
	var res int
	for _, v := range nums {
		if v < k {
			res++
		}
	}
	return res
}
