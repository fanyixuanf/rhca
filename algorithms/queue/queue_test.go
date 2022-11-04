/*
@Time : 2022/11/4 21:26
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : queue_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	var q = []int{1,2}
	q = append([]int{0}, q...)
	fmt.Printf("queue : %p, %v\n", q,q)
}

func TestQueue1(t *testing.T) {
	fmt.Printf("queue %p, %v\n", queue, queue)
	enqueue([]int{1})
	fmt.Printf("enqueue %p, %v\n", queue, queue)
	enqueue([]int{2})
	fmt.Printf("enqueue %p, %v\n", queue, queue)
	enqueue([]int{3})
	fmt.Printf("enqueue %p, %v\n", queue, queue)
	dequeue()
	fmt.Printf("dequeue %p, %v\n", queue, queue)
	dequeue()
	fmt.Printf("dequeue %p, %v\n", queue, queue)
	dequeue()
	fmt.Printf("dequeue %p, %v\n", queue, queue)
	dequeue()
	fmt.Printf("dequeue %p, %v\n", queue, queue)
}