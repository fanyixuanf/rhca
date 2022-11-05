/*
@Time : 2022/11/5 17:26
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : deque_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package deque

import (
	"fmt"
	"testing"
)

func TestDeque(t *testing.T) {
	fmt.Printf("init deque, %p, %v\n", deque, deque)
	lendeque([]int{1})
	fmt.Printf("lendeque, %p, %v\n", deque, deque)
	lendeque([]int{0})
	fmt.Printf("lendeque, %p, %v\n", deque, deque)
	rendeque([]int{2})
	fmt.Printf("rendeque, %p, %v\n", deque, deque)
	rendeque([]int{3})
	fmt.Printf("rendeque, %p, %v\n", deque, deque)
	fmt.Println("de queue ....")
	ldedeque()
	fmt.Printf("ldedeque, %p, %v\n", deque, deque)
	rdedeque()
	fmt.Printf("rdedeque, %p, %v\n", deque, deque)

}