/*
@Time : 2022/11/14 21:10
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : heapSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package heapSort

import (
	"container/heap"
	"fmt"
	myheap "rhca/algorithms/heap"
	"testing"
)

func TestHeapSort(t *testing.T) {
	//arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	intHeap := myheap.IntHeap{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	heap.Init(&intHeap)
	fmt.Println(intHeap)
}