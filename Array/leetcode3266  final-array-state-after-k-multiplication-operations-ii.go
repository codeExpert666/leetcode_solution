package Array

import (
	"container/heap"
	"sort"
)

// 快速幂运算，核心思想是将指数 n 转化为二进制，只计算二进制中为 1 的部分，从而将幂运算分解，减少了乘法运算的次数。
// 复杂度为 log(n)
func pow(a, n, mod int) int {
	result := 1 // 存储最终结果
	for n > 0 {
		if n&1 == 1 { // 如果当前二进制位为1
			result = (result * a) % mod // 将当前的a累乘到结果中
		}
		a = (a * a) % mod // a自乘，相当于每次把a变成a²
		n >>= 1           // n右移一位，相当于除以2
	}
	return result
}

// 在小根堆的基础上进行优化，当某个元素乘以 multiplier 后大于原数组中的最大值时，
// 后续的操作不再需要小根堆，转变为数组所有元素轮流乘以 multiplier，直到 k 次结束
// 注意这题数据很大，特别容易溢出，所以一切围绕不溢出的思路去写
func GetFinalState1(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 特殊情况
		return nums
	}
	n := len(nums)
	mod := int(1e9 + 7)
	// 获取数组元素最大值，同时构造小根堆
	maxValue := 0
	hp := make(minHeapv2, n)
	for i, v := range nums {
		if v > maxValue { // 更新最大值
			maxValue = v
		}
		hp[i] = pair{i, v} // 构建小根堆
	}
	heap.Init(&hp)
	for ; k > 0; k-- {
		if tmp := hp[0].value * multiplier; tmp <= maxValue {
			// 处理最小元素
			nums[hp[0].index] = tmp
			hp[0].value = tmp
			heap.Fix(&hp, 0)
		} else { // 不再需要小根堆，转为数学计算
			break
		}
	}
	if k > 0 { // 数学计算
		sort.Sort(&hp)
		round, extra := k/n, k%n
		p := pow(multiplier, round, mod)
		for i, pr := range hp {
			nums[pr.index] = nums[pr.index] % mod * p % mod
			if i < extra {
				nums[pr.index] = nums[pr.index] * multiplier % mod
			}
		}
	}
	return nums
}

type pair struct {
	index int
	value int
}

func less(p1, p2 pair) bool {
	return p1.value < p2.value || (p1.value == p2.value && p1.index < p2.index)
}

type minHeapv2 []pair

func (hp *minHeapv2) Len() int {
	return len(*hp)
}

func (hp *minHeapv2) Less(i, j int) bool {
	return less((*hp)[i], (*hp)[j])
}

func (hp *minHeapv2) Swap(i, j int) {
	(*hp)[i], (*hp)[j] = (*hp)[j], (*hp)[i]
}

func (hp *minHeapv2) Push(any) {}

func (hp *minHeapv2) Pop() (_ any) { return }
