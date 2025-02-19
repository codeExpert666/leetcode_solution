package Array

func ReplaceElements(arr []int) []int {
	maxRight := -1
	tmp := 0
	for i := len(arr) - 1; i >= 0; i-- {
		tmp = arr[i]
		arr[i] = maxRight
		maxRight = max(maxRight, tmp)
	}
	return arr
}
