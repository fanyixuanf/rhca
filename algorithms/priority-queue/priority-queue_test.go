/*
@Time : 2022/11/5 19:11
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : priority-queue_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package priority_queue

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// 初始化数据
	items := map[string]int{
		"A job": 3,
		"B job": 2,
		"C job": 4,
		"D job": 1,
		"F job": 1,
	}
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	// 初始化最大堆
	heap.Init(&pq)
	// 添加数据到队列中
	item := &Item{
		value:    "E job",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.Print(pq)
	fmt.Println("-------------------")
	// 更改优先级
	pq.update(item, item.value, 5)
	pq.Print(pq)
	fmt.Println("-------------------")
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s \n", item.priority, item.value)
	}

}