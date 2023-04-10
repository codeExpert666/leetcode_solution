package BackTracking

import (
	"sort"
	"strings"
)

/**
 * 回溯，这种方法思路与过程都是正确的，问题在于时间复杂度太高，lc提交超出时间限制
 * 仔细分析后，个人任务复杂度高的原因是我把本题当作全排列去做的，每层都是从零开始遍历，这没错，
 * 但没有必要，因为本题的输入本质上是一张图，一些遍历明显是没有意义的，当path的最后一站确定后
 * 它的下一站只能是出边所对应的地点
 */
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

/**
 * 回溯，先构造出图结构，再全排列(从图的角度是深度优先搜索)
 * 仔细体会这里的写法，很考验基本功
 */
type pair struct {
	target  string
	visited bool // 一个pair实际上对应一张机票，该bool值反映该票是否被使用
}
type pairs []*pair

// 实现了下面三个函数等价于实现了sort接口，便可以进行排序
// 这部分是不熟悉的地方，多看多学

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairs) Less(i, j int) bool {
	return p[i].target < p[j].target
}

func findItinerary1(tickets [][]string) []string {
	var result []string
	// map[出发机场] pair{目的地,是否被访问过}
	targets := make(map[string]pairs)
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(pairs, 0)
		}
		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
	}
	for k, _ := range targets {
		sort.Sort(targets[k])
	}
	result = append(result, "JFK")
	var backtracking func() bool
	backtracking = func() bool {
		if len(tickets)+1 == len(result) {
			return true
		}
		// 取出起飞航班对应的目的地
		for _, pair := range targets[result[len(result)-1]] {
			if pair.visited == false {
				result = append(result, pair.target)
				pair.visited = true
				if backtracking() {
					return true
				}
				result = result[:len(result)-1]
				pair.visited = false
			}
		}
		return false
	}

	backtracking()

	return result
}
