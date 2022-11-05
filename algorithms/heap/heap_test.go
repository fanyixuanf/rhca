/*
@Time : 2022/11/5 17:42
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : heap_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := &IntHeap{2,1,5,7,2}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("min : %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d ",heap.Pop(h))
	}
}

func TestMaxHeap(t *testing.T) {
	m := &MaxHeap{2,1,5,7,2}
	heap.Init(m)
	heap.Push(m, 3)
	fmt.Printf("max : %d\n", (*m)[0])
	for m.Len() > 0 {
		fmt.Printf("%d ",heap.Pop(m))
	}
}
