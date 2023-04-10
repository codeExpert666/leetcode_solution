package main

import (
	"fmt"
	"strings"
)

var (
	res12  []string
	path12 []string
)

func findItinerary(tickets [][]string) []string {
	res12, path12 = make([]string, 0, len(tickets)+1), make([]string, 0, len(tickets)+1)
	path12 = append(path12, "JFK")
	backtracking12(tickets, make([]bool, len(tickets)))
	return res12
}

func backtracking12(tickets [][]string, used []bool) {
	// 终止条件
	if len(path12) == len(tickets)+1 {
		if len(res12) == 0 {
			// 第一次找到行程
			res12 = append(res12, path12...)
		} else {
			// 第二次及以后找到行程，需要与之前的行程做比较
			if comparePath(path12, res12) {
				res12 = res12[:0]
				res12 = append(res12, path12...)
			}
		}
		return
	}

	for i := 0; i < len(tickets); i++ {
		if used[i] || tickets[i][0] != path12[len(path12)-1] {
			continue
		}
		path12 = append(path12, tickets[i][1])
		used[i] = true
		backtracking12(tickets, used)
		used[i] = false
		path12 = path12[:len(path12)-1]
	}
}

// 判断path1是否比path2更好
func comparePath(path1, path2 []string) bool {
	var r bool
	for i := 0; i < len(path1); i++ {
		v := strings.Compare(path1[i], path2[i])
		if v == -1 {
			r = true
			break
		}
		if v == 1 {
			r = false
			break
		}
	}
	return r
}

func main() {
	fmt.Println(findItinerary([][]string{
		[]string{"JFK", "KUL"},
		[]string{"JFK", "NRT"},
		[]string{"NRT", "JFK"},
	}))
}
