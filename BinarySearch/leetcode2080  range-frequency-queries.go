package binarysearch

type RangeFreqQuery struct {
	IndexArrays map[int][]int
}

// 离线统计：确定每个数在 arr 中出现的下标，形成下标数组
func Constructor(arr []int) RangeFreqQuery {
	rfq := RangeFreqQuery{
		IndexArrays: make(map[int][]int),
	}
	for i, v := range arr {
		rfq.IndexArrays[v] = append(rfq.IndexArrays[v], i)
	}
	return rfq
}

// 二分查找：查找 value 对应的下标数组中 [left, right] 包括的元素个数
func (rfq *RangeFreqQuery) Query(left int, right int, value int) int {
	iArray := rfq.IndexArrays[value] // 获取 value 的下标数组
	// 二分查找大于等于 left 的第一个数的下标
	p := getGreater(iArray, left)
	// 二分查找大于 right 的第一个数的下标，也即大于等于 right+1
	// 注意这里的下标是 iArray[p:] 的下标
	return getGreater(iArray[p:], right+1)
}

// 二分查找大于等于 value 的第一个数的下标，若不存在，返回 arr 长度
func getGreater(arr []int, value int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
