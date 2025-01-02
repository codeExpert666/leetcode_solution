package Array

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendar struct {
	orderedEvent *redblacktree.Tree // 将不交叉的日程安装起始时间排序
}

func Constructor() MyCalendar {
	return MyCalendar{
		orderedEvent: redblacktree.NewWithIntComparator(),
	}
}

func (mc *MyCalendar) Book(startTime int, endTime int) bool {
	cal := mc.orderedEvent
	// 在日历中找到输入日程前的最近日程，输入日程的起始时间不能与前一个最近日程交叉
	if pre, exist := cal.Floor(startTime); exist && pre.Value.(int) > startTime {
		return false
	}
	// 在日历中找到输入日程后的最近日程，输入日程的结束时间不能超过后一个最近日程的起始时间
	if after, exist := cal.Ceiling(startTime); exist && after.Key.(int) < endTime {
		return false
	}
	// 插入该日程
	cal.Put(startTime, endTime)
	return true
}
