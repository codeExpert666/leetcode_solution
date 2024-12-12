package Greedy

import "container/heap"

var globalValues [][]int

// 写的有点啰嗦了，直接将 PQueue 的底层类型定义为二维切片更好。同时，也不需要定义 globalValues，
// 直接将 []int{price, i, j} 作为队列的元素即可。也就是说 PQueue 是一个二维切片，其每个元素是一个长度为 3 的数组（切片）
// 店铺 i 与其商品 j 组成的对
type pair [2]int

type PQueue []pair

func (lp *PQueue) Len() int {
	return len(*lp)
}

func (lp *PQueue) Less(i, j int) bool {
	return globalValues[(*lp)[i][0]][(*lp)[i][1]] < globalValues[(*lp)[j][0]][(*lp)[j][1]]
}

func (lp *PQueue) Swap(i, j int) {
	(*lp)[i], (*lp)[j] = (*lp)[j], (*lp)[i]
}

func (lp *PQueue) Push(x any) {
	*lp = append(*lp, x.(pair))
}

func (lp *PQueue) Pop() any {
	l := lp.Len()
	top := (*lp)[l-1]
	*lp = (*lp)[:l-1]
	return top
}

// 贪心：尽可能靠后购买价格更高的商品，也即价格越低的商品应优先购买
func MaxSpending(values [][]int) int64 {
	m, n := len(values), len(values[0])
	globalValues = values
	// 优先队列存储每家商店当前最低价格物品的下标
	lowestPrices := make(PQueue, m)
	for i := 0; i < m; i++ { // 队列初始化
		lowestPrices[i] = pair{i, n - 1}
	}
	heap.Init(&lowestPrices)
	// 逐个取出当前价格最低商品
	var res, day int
	for lowestPrices.Len() > 0 {
		// 累计当前商品开销
		day++
		lowestPair := heap.Pop(&lowestPrices).(pair)
		res += values[lowestPair[0]][lowestPair[1]] * day
		// 该店铺仍存在商品
		if lowestPair[1] > 0 {
			heap.Push(&lowestPrices, pair{lowestPair[0], lowestPair[1] - 1})
		}
	}
	return int64(res)
}
