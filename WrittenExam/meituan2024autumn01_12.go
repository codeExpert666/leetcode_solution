package main

/*
	链接：https://www.nowcoder.com/exam/test/85715094/detail?pid=58083953&testCallback=https%3A%2F%2Fwww.nowcoder.com%2Fexam%2Fcompany%3FquestionJobId%3D10%26subTabName%3Dwritten_page
*/
import (
	"fmt"
)

func MT12() {
	var testNum int
	_, _ = fmt.Scan(&testNum)
	for ; testNum > 0; testNum-- {
		var n, k, x int
		_, _ = fmt.Scan(&n, &k, &x)

		array := make([]int, n)
		countMap := make([]int, n+2)
		for i := 0; i < n; i++ {
			var elem int
			_, _ = fmt.Scan(&elem)
			array[i] = elem
			countMap[elem]++
		}

		var res, mex int
		for i, v := range countMap {
			if v == 0 {
				mex = i
				break
			}
		}
		res = k * mex

		for i, v := range array {
			countMap[v]--
			if countMap[v] == 0 && v < mex {
				mex = v
			}
			res = min(res, (i+1)*x+k*mex)
		}

		fmt.Println(res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
