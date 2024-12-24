package Array

import "slices"

func sortTheStudents(score [][]int, k int) [][]int {
    slices.SortFunc(score, func(a, b []int) int {
		return b[k] - a[k]
	})
	return score
}