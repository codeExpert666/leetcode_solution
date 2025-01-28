package weeklycompetition

import "sort"

func maxFrequency(nums []int, k int) int {
	n := len(nums)

	kPos := make([]int, 0)
	first := make([]int, 51)
	for i := 1; i <= 50; i++ {
		first[i] = n
	}
	last := make([]int, 51)
	for i := 1; i <= 50; i++ {
		last[i] = -1
	}
	count := make([]int, 51)

	for i, v := range nums {
		if i < first[v] {
			first[v] = i
		}
		if i > last[v] {
			last[v] = i
		}
		if v == k {
			kPos = append(kPos, i)
		}
		count[v]++
	}

	kNum, res := len(kPos), len(kPos)
	for i := 1; i <= 50; i++ {
		if i != k && first[i] != n {
			tmp := countInRange(kPos, first[i], last[i])
			res = max(res, kNum+count[i]-tmp)
		}
	}

	return res
}

func countInRange(nums []int, start, end int) int {
	left := binarySearch(nums, start)
	right := binarySearch(nums, end)
	return right - left
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func maxFrequency1(nums []int, k int) int {
	m := 0
	for _, num := range nums {
		if num == k {
			m++
		}
	}

	countV := make(map[int]int)
	for _, num := range nums {
		if num != k {
			countV[num]++
		}
	}

	type vCount struct {
		v     int
		count int
	}
	var vList []vCount
	for v, c := range countV {
		vList = append(vList, vCount{v, c})
	}

	sort.Slice(vList, func(i, j int) bool {
		return vList[i].count > vList[j].count
	})

	maxGain := 0
	for _, vc := range vList {
		v := vc.v
		currentCount := vc.count
		if currentCount <= maxGain {
			continue
		}

		currentMax := 0
		currentSum := 0
		for _, num := range nums {
			var val int
			if num == v {
				val = 1
			} else if num == k {
				val = -1
			} else {
				val = 0
			}
			currentSum += val
			if currentSum < 0 {
				currentSum = 0
			}
			if currentSum > currentMax {
				currentMax = currentSum
			}
		}

		if currentMax > maxGain {
			maxGain = currentMax
		}
	}

	total := m + maxGain
	if total > m {
		return total
	}
	return m
}
