/*
@Time : 2022/11/4 21:26
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : queue
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package queue

var queue []int

func enqueue(num []int) {
	queue = append(num, queue...)
}

func dequeue() {
	if len(queue) == 0 {
		return
	}
	queue = queue[:len(queue) - 1]
}