/*
@Time : 2022/11/5 19:11
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : priority-queue
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package priority_queue

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value string // 值
	priority int // 优先级
	index int // 堆的索引
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority >= pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	// 交换值
	pq[i], pq[j] = pq[j], pq[i]
	// 索引不变
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop () interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[ : n - 1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// 打印目前数据

func (pq *PriorityQueue) Print(q PriorityQueue) {
	for i, v := range q {
		fmt.Printf("index: %d, value: %s, Priority: %d \n", i, v.value, v.priority)
	}
}
