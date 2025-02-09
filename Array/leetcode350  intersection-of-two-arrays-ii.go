package Array

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		m2[v]++
	}

	res := make([]int, 0)
	for k, v := range m1 {
		tmp := min(v, m2[k])
		for i := 1; i <= tmp; i++ {
			res = append(res, k)
		}
	}

	return res
}
