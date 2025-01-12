package weeklycompetition

import (
	"container/heap"
)

type TaskInfo struct {
	Priority int
	UserId   int
}

// 用于在优先级队列中存储的元素
type item struct {
	priority int
	taskId   int
	userId   int
}

// 实现heap.Interface
type maxHeap []item

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	// 优先比较优先级，优先级大的排前面
	// 若优先级相同，则taskId大的排前面
	if h[i].priority == h[j].priority {
		return h[i].taskId > h[j].taskId
	}
	return h[i].priority > h[j].priority
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type TaskManager struct {
	// 记录taskId -> (priority, userId)
	tasks map[int]TaskInfo
	// 优先级队列，存所有task
	pq maxHeap
}

// func Constructor
func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		tasks: make(map[int]TaskInfo),
		pq:    make(maxHeap, 0),
	}

	// 初始化
	for _, t := range tasks {
		userId, taskId, priority := t[0], t[1], t[2]
		tm.tasks[taskId] = TaskInfo{priority, userId}
		heap.Push(&tm.pq, item{
			priority: priority,
			taskId:   taskId,
			userId:   userId,
		})
	}
	return tm
}

// func (this *TaskManager) Add
func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.tasks[taskId] = TaskInfo{priority, userId}
	heap.Push(&this.pq, item{
		priority: priority,
		taskId:   taskId,
		userId:   userId,
	})
}

// func (this *TaskManager) Edit
func (this *TaskManager) Edit(taskId int, newPriority int) {
	// 更新任务优先级
	if info, ok := this.tasks[taskId]; ok {
		info.Priority = newPriority
		this.tasks[taskId] = info
		// 在堆中新push一个更新后的元素，旧元素“延迟删除”
		heap.Push(&this.pq, item{
			priority: newPriority,
			taskId:   taskId,
			userId:   info.UserId,
		})
	}
}

// func (this *TaskManager) Rmv
func (this *TaskManager) Rmv(taskId int) {
	if _, ok := this.tasks[taskId]; ok {
		delete(this.tasks, taskId)
	}
	// 堆中的元素同样采取“延迟删除”策略
}

// func (this *TaskManager) ExecTop
func (this *TaskManager) ExecTop() int {
	// 不断弹出堆顶，看是否合法
	for this.pq.Len() > 0 {
		top := heap.Pop(&this.pq).(item)
		info, ok := this.tasks[top.taskId]
		if !ok {
			// 说明已经被移除或更新，跳过
			continue
		}
		// 如果堆顶信息与当前存储信息一致，说明是最新的有效数据
		if info.Priority == top.priority && info.UserId == top.userId {
			// 找到最高优先级、且taskId最大的有效任务
			delete(this.tasks, top.taskId) // 执行后移除
			return top.userId
		}
		// 如果不一致，说明是旧数据，继续弹其他元素
	}
	// 没有任何任务
	return -1
}
