package StringFunc

import (
	"cmp"
	"maps"
	"slices"
	"strings"
)

type team struct {
	name      byte  // 队伍名
	rankVotes []int // 队伍在每个排位获得的票数
}

func less(t1, t2 team) int {
	for i := 0; i < len(t1.rankVotes); i++ {
		if t1.rankVotes[i] != t2.rankVotes[i] {
			return t2.rankVotes[i] - t1.rankVotes[i]
		}
	}
	// 注意先转 int 再做减法，因为先减结果是 byte 类型，而如果结果是负数，会溢出。
	return int(t1.name) - int(t2.name)
}

func RankTeams(votes []string) string {
	if len(votes) == 1 { // 特殊情况
		return votes[0]
	}
	teamNum := len(votes[0])          // 队伍数量
	teams := make([]team, 0, teamNum) // 存储所有队伍
	teamIndex := [26]byte{}           // 记录每个队伍在 teams 中的下标
	for i := 0; i < 26; i++ {
		teamIndex[i] = 26
	}
	// 统计票数
	for _, v := range votes {
		for i := 0; i < teamNum; i++ {
			if index := teamIndex[v[i]-'A']; index != 26 { // v[i] 队伍已有存储记录
				teams[index].rankVotes[i]++
			} else { // 创建 v[i] 队伍的存储记录，添加到 teams 的末尾
				teamIndex[v[i]-'A'] = byte(len(teams))
				teams = append(teams, team{
					name:      v[i],
					rankVotes: make([]int, teamNum),
				})
				teams[len(teams)-1].rankVotes[i]++
			}
		}
	}
	// 队伍排序
	slices.SortFunc(teams, func(t1, t2 team) int {
		return less(t1, t2)
	})
	// 处理队伍排序字符串
	var builder strings.Builder
	for _, t := range teams {
		builder.WriteByte(t.name)
	}
	return builder.String()
}

// 灵神大佬优雅写法，尤其是对 cmp 包的使用！
// PS：我总是觉得 map 的各种操作都太耗时，所以没敢用。但对于小数据量的 map，
// 该用还得用，性能其实也是可以的，用了 map 会减少很多心智负担。
func RankTeams1(votes []string) string {
	cnts := map[rune][]int{}
	for _, ch := range votes[0] {
		cnts[ch] = make([]int, len(votes[0]))
	}
	for _, vote := range votes {
		for i, ch := range vote {
			cnts[ch][i]++
		}
	}

	ans := slices.SortedFunc(maps.Keys(cnts), func(a rune, b rune) int {
		return cmp.Or(slices.Compare(cnts[b], cnts[a]), cmp.Compare(a, b))
	})
	return string(ans)
}
