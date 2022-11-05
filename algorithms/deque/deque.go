/*
@Time : 2022/11/5 17:26
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : deque.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package deque

var deque []int

func lendeque(q []int) {
	deque = append(q, deque...)
}

func rendeque(q []int) {
	deque = append(deque, q...)
}

func ldedeque() {
	deque = deque[1:]
}

func rdedeque() {
	deque = deque[:len(deque) - 1]
}
