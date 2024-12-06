package MathFunc

import (
	"sort"
)

// minimumRemoval 计算需要移除的最少魔法豆数量，使得剩下的非空袋子中每个袋子的魔法豆数量相等。
// 解答本题抓住一点：最终所有袋子中保留的豆子数量一定是初始所有袋子中某个袋子中的豆子数量。
// 于是，我们可以先对 beans 数组进行升序排序，然后遍历 beans 数组，对于当前袋子中的豆子数量 bean[i]，要想保留它，
// 那么其左侧的所有豆子均要移除，其右侧的豆子移除部分以保证所有袋子中的豆子数量相等。
func MinimumRemoval(beans []int) int64 {
	// 对 beans 数组进行排序，便于后续计算
	sort.Ints(beans)
	n := len(beans)
	var total int64 = 0

	// 计算所有袋子中魔法豆的总数
	for _, bean := range beans {
		total += int64(bean)
	}

	// 初始化最小移除数量为总魔法豆数量
	minRemoval := total

	// 遍历排序后的 beans 数组
	for i, bean := range beans {
		// 计算保留从当前索引 i 到末尾所有袋子，每个袋子保留 bean 个魔法豆所需的总数
		remaining := int64(bean) * int64(n-i)

		// 计算需要移除的魔法豆数量
		removal := total - remaining

		// 更新最小移除数量
		if removal < minRemoval {
			minRemoval = removal
		}
	}

	// 返回最小需要移除的魔法豆数量
	return minRemoval
}
