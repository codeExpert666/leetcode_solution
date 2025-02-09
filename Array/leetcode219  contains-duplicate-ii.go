package Array

func containsNearbyDuplicate(nums []int, k int) bool {
	numToLastIndex := make(map[int]int)

	for i, v := range nums {
		if lastIndex, exist := numToLastIndex[v]; exist && i-lastIndex <= k {
			return true
		} else {
			numToLastIndex[v] = i
		}
	}

	return false
}
