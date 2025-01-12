package StringFunc

import (
	"fmt"
	"strconv"
)

func LargestGoodInteger(num string) string {
	var res = -1

	var count = 1 // 连续出现的字符数量
	for i := 1; i < len(num); i++ {
		if num[i] != num[i-1] {
			count = 1
		} else if count == 1 {
			count++
		} else if count == 2 {
			count++
			if n, _ := strconv.Atoi(num[i-2 : i+1]); n > res {
				res = n
			}
		}
	}

	if res == -1 {
		return ""
	}
	if res == 0 {
		return "000"
	}
	return fmt.Sprintf("%d", res)
}

// LargestGoodInteger1 别人的优雅写法
func LargestGoodInteger1(num string) string {
	ansStr := ""
	for i := len(num) - 3; i >= 0; i-- {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if len(ansStr) == 0 || ansStr[0] < num[i] {
				ansStr = num[i : i+3]
			}
		}
	}
	return ansStr
}
