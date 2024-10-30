package DailyCode

import "sort"

/**
 * 方法一：先对人与身高进行配对，再利用内置函数按照身高进行排序，最后获取排好序的人名
 */

type PHeight struct {
	name   string
	height int
}

func SortPeople(names []string, heights []int) []string {
	// 配对
	pairs := make([]*PHeight, 0, len(names))
	for i, n := range names {
		pairs = append(pairs, &PHeight{
			name:   n,
			height: heights[i],
		})
	}
	// 排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].height > pairs[j].height
	})
	// 获取人名
	res := make([]string, 0, len(names))
	for _, v := range pairs {
		res = append(res, v.name)
	}
	return res
}

/**
 * 方法二：自己实现排序，在对heights排序的过程中，对names也做同样操作
 * 不再赘述
 */
