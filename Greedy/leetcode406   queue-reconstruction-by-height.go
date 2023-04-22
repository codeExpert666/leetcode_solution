package Greedy

import (
	"container/list"
	"sort"
)

/**
 * Tips：具体来说，当当前slice的长度小于1024时，执行append操作，新slice的capacity会变成当前的2倍；
 * 而当slice长度大于等于1024时，slice的扩容变成了每次增加当前slice长度的1/4。
 * 在Go Slice的底层实现中，如果capacity不够时，会做一个reslice的操作，底层数组也会重新被复制到另一块内存区域中，
 * 所以append一个元素，不一定是O(1), 也可能是O(n)哦。
 */

/**
 * 贪心法
 * 本题有两个维度，如果两个维度同时考虑，会顾此失彼，所以只能先考虑一个维度，再在当前
 * 维度的基础上考虑另一个维度。本题首先考虑身高维度，先按照身高h由大到小进行排序，再按照people
 * 的k属性（前面只能有k个大于等于当前身高的人）进行排序调整，不断由局部最优调整为全局最优。
 * 注意到，这种调整是按照身高由大到小逐人调整的，由于每个人前面的人都比自己高，所以可以直接把k
 * 当作调整后的下标，同时身高的由大到小保证了后面调整上来的人不会影响之前的调整结果（因为后面的人身高更矮，不影响k）
 */
func reconstructQueue(people [][]int) [][]int {
	// 先按照身高由大到小进行排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	// 遍历数组，按照k属性进行调整
	for i, p := range people {
		if p[1] != i {
			copy(people[p[1]+1:i+1], people[p[1]:i]) // 前面的元素后移
			people[p[1]] = p                         // 调整到正确位置
		}
	}
	return people
}

/**
 * 贪心法二
 * 与上面思路完全一致，只是实现方式不同，上面用的顺序表，本方法使用链表
 */
func reconstructQueue1(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] //当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0]
	})
	l := list.New() //创建链表
	for i := 0; i < len(people); i++ {
		position := people[i][1]
		mark := l.PushBack(people[i]) //插入元素
		e := l.Front()
		for position != 0 { //获取相对位置
			position--
			e = e.Next()
		}
		l.MoveBefore(mark, e) //移动位置

	}
	var res [][]int
	for e := l.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.([]int))
	}
	return res
}
