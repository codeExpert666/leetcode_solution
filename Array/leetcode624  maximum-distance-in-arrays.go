package Array

import (
	"math"
	"slices"
)

// 按照首元素升序排列后，最大距离只可能出现在一些固定位置上
// 该方法的性能排序占大头，所以不算高效
func MaxDistance1(arrays [][]int) int {
	m, n := len(arrays), len(arrays[0])
	// 按照首元素排序
	slices.SortFunc(arrays, func(a, b []int) int {
		return a[0] - b[0]
	})

	// 其中一个数为第一行最后一个元素
	res := max(abs(arrays[1][0]-arrays[0][n-1]), abs(arrays[m-1][0]-arrays[0][n-1]))

	// 其中一个数为第一行第一个元素
	maxValue := -10001
	for i := 1; i < m; i++ {
		l := len(arrays[i])
		maxValue = max(maxValue, arrays[i][l-1])
	}
	res = max(res, maxValue-arrays[0][0])

	return res
}

// 灵神解法：遍历右，维护左
// 具体见链接：https://leetcode.cn/problems/maximum-distance-in-arrays/solutions/3067679/mei-ju-you-wei-hu-zuo-pythonjavaccgojsru-wtgb/
func MaxDistance2(arrays [][]int) (ans int) {
	// 这里最大值最小值的初始值设计非常巧妙，仔细体会
	mn, mx := math.MaxInt/2, math.MinInt/2 // 防止减法溢出
	for _, a := range arrays {
		x, y := a[0], a[len(a)-1]
		ans = max(ans, y-mn, mx-x)
		// 这里 mn 与 mx 的更新体现了“维护左”，具体解释：
		// 遍历右意味着从当前遍历数组取出一个数（也即从 a 中取出一个数），剩余一个数从
		// a 之前的任意一个数组中取一个，而这等价与从 a 之前的数组的并集中取一个。
		// 于是这又转化为灵神题解中两个有序集合的最大距离理论。
		// 而下面的更新就保证了 mn 与 mx 中始终维护着左侧并集的最大最小值
		mn = min(mn, x)
		mx = max(mx, y)
	}
	return
}
