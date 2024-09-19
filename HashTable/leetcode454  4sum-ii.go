package HashTable

// 两两一组，计算可能的数字和，存储在两个hash表中
// 对于hash表1中的一个值，如果能在hash表2中找到和为0的值，则记录
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var res int
	// 记录数字和的出现情况
	sumCount1 := make(map[int]int) // key为数字和，value为可能的组合数
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			sumCount1[v1+v2]++
		}
	}
	sumCount2 := make(map[int]int)
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			sumCount2[v3+v4]++
		}
	}
	// 遍历hash表
	for sum, count1 := range sumCount1 {
		if count2, exist := sumCount2[-sum]; exist {
			res += count1 * count2
		}
	}
	return res
}

// 优化版本：一个hash表即可。在处理nums3和nums4时，无需再存储，边遍历边查看nums1与nums2对应的hash表情况即可
func fourSumCount1(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumCount := make(map[int]int) // 用于记录num1与num2的和的出现情况与次数
	var res int
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			sumCount[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			res += sumCount[-n3-n4]
		}
	}
	return res
}
