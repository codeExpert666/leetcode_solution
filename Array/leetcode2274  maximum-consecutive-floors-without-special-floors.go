package Array

import "slices"

func MaxConsecutive(bottom int, top int, special []int) int {
	slices.SortFunc(special, func(a, b int) int { return a - b })

	l := len(special)
	res := max(special[0]-bottom, top-special[l-1])

	for i := 1; i < l; i++ {
		res = max(res, special[i]-special[i-1]-1)
	}

	return res
}
