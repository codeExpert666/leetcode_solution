package Greedy

/**
 * gas与cost数组之差反映了汽车在每个加油站的油量得失情况，先利用总体得失情况判断是否有解
 * 若有解，则每个站都有可能是起始站， 利用起始站的油量得失一定为正进行剪枝，再对剩余情况
 * 循环模拟判断，找出解
 * 该方法本质上还是暴力法，很遗憾，lc超时
 * ！！！for循环适合模拟从头到尾的遍历，而while循环适合模拟环形遍历，要善于使用while！！！
 */
func canCompleteCircuit(gas []int, cost []int) int {
	var sum int // 记录总体油量得失情况
	for i := 0; i < len(gas); i++ {
		gas[i] = gas[i] - cost[i] // 将每一站的油量得失存在gas中
		sum += gas[i]
	}
	if sum < 0 {
		// 若总体油量得失为负，则无解
		return -1
	}
	// 若有解
	var res int
	for ; res < len(gas); res++ {
		if gas[res] >= 0 && isSolution(gas, res) {
			// 起始站的油量得失一定为正，利用此结论进行剪枝
			break
		}
	}
	return res
}

// 判断给定站能否成为起始站
func isSolution(gasAdded []int, start int) bool {
	// 注意加油站站都在环上，处理方式类似于循环队列
	var curGas int // 记录当前油量
	index := start // 当前所在站
	l := len(gasAdded)
	for ; index != (start-1)%l; index = (index + 1) % l {
		curGas += gasAdded[index]
		if curGas < 0 {
			// 一旦经过该站油量为负，则该站不能成为起始站
			return false
		}
	}
	// 上面循环剩余终点站未处理
	curGas += gasAdded[index]
	return curGas >= 0
}

/**
 * 方法二：贪心
 * 首先如果总油量减去总消耗大于等于零那么一定可以跑完一圈，说明各个站点的加油站剩油量rest[i]相加一定是大于等于零的。
 * 每个加油站的剩余量rest[i]为gas[i] - cost[i]。
 * i从0开始累加rest[i]，和记为curSum，一旦curSum小于零，说明[0, i]区间都不能作为起始位置，因为这个区间选择任何一个位置作为起点，
 * 到i这里都会断油，那么起始位置从i+1算起，再从0计算curSum。
 * ！！！！好好体会这个解法，就一点没想到“一旦curSum小于零，说明[0, i]区间都不能作为起始位置”！！！！
 */
func canCompleteCircuit1(gas []int, cost []int) int {
	curSum := 0   // 当前剩余油量
	totalSum := 0 // 总体油量得失情况
	start := 0    // 起点加油站
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			// 出现断油，则起点start到当前加油站之间的所有加油站都不能作为真正起点
			// 重新记录
			start = i + 1
			curSum = 0
		}
	}
	// 上面的部分给出了当解存在时，起点的可能位置[start, len(gas)-1]
	// 但解不一定存在，所以下面进行判断
	if totalSum < 0 {
		// 总体油量得失为负值，一定无解
		return -1
	}
	// 有解，则解一定在[start, len(gas)-1]之中，反正法：若start不是解，则说明其在[0. start)之间断油
	// 由上面的循环可知，在[start, len(gas)-1]中，curSum一直大于0，故[start, len(gas)-1]均不可能
	// 是解，这与有解矛盾！故start就是解。
	return start
}
