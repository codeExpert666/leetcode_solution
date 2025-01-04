package Array

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type partInterval struct {
	startTime int
	repeat    bool
}

type MyCalendarTwo struct {
	// key 为 endTime，value 为 startTime 以及 bool (该时间区间是否已经重复出现)
	booked *redblacktree.Tree
}

func ConstructorMCT() MyCalendarTwo {
	rbt := redblacktree.NewWithIntComparator()
	rbt.Put(math.MaxInt, partInterval{math.MaxInt, false})
	return MyCalendarTwo{
		booked: rbt,
	}
}

func (mc *MyCalendarTwo) Check(startTime int, endTime int) bool {
	node, _ := mc.booked.Ceiling(startTime + 1) // 树中待处理区间
	it := mc.booked.IteratorAt(node)
	for it.Node().Value.(partInterval).startTime < endTime {
		if it.Node().Value.(partInterval).repeat { // 该区间已经重复预定
			return false
		}
		it.Next()
	}
	return true
}

// 从 startTime 开始，不断处理在其后结束的时间区间，直到区间的开始时间在 endTime 之后
func (mc *MyCalendarTwo) Book(startTime int, endTime int) bool {
	if !mc.Check(startTime, endTime) { // 失败直接返回，不做任何操作
		return false
	}
	// 如果成功，则需要对区间进行增删
	sTime := startTime                      // 待插入区间的起始时间
	node, _ := mc.booked.Ceiling(sTime + 1) // 树中待处理区间
	for {
		nodeStart := node.Value.(partInterval).startTime
		nodeEnd := node.Key.(int)
		if nodeStart >= endTime { // 后续区间不受影响
			// 添加区间（最后一个受影响的区间的靠后差集部分）
			if sTime != endTime { // 此时 sTime 的值就是上一个受影响区间的 nodeEnd
				mc.booked.Put(max(sTime, endTime), partInterval{
					startTime: min(sTime, endTime),
					repeat:    false,
				})
			}
			break
		}
		// 移除区间
		mc.booked.Remove(nodeEnd)
		// 添加区间（靠前差集部分）
		if nodeStart != sTime {
			mc.booked.Put(max(nodeStart, sTime), partInterval{
				startTime: min(nodeStart, sTime),
				repeat:    false,
			})
		}
		// 添加区间（相交部分）
		mc.booked.Put(min(nodeEnd, endTime), partInterval{
			startTime: max(nodeStart, sTime),
			repeat:    true,
		})

		sTime = nodeEnd                        // 更新待插入区间起始时间
		node, _ = mc.booked.Ceiling(sTime + 1) // 更新树中待处理区间（加 1 是为了不再把当前 node 包括进去）
	}

	// // test
	// fmt.Println("-------------------------------------")
	// for _, e := range mc.booked.Keys() {
	// 	info, _ := mc.booked.Get(e)
	// 	fmt.Println(info.(partInterval).startTime, e, info.(partInterval).repeat)
	// }

	return true
}
