/*
@Time : 2022/11/3 21:04
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : stack_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package stack

import (
	"fmt"
	"testing"
)

/**
入栈前, [1 2 3 4 5]
入栈1, [1 2 3 4 5 6]
入栈2, [1 2 3 4 5 6 7]
出栈1, [1 2 3 4 5 6]
出栈2, [1 2 3 4 5]
出栈3, [1 2 3 4]
 */
func TestStack(t *testing.T) {
	var a = []int{1,2,3,4,5}
	fmt.Printf("入栈前, %v\n", a)
	// 入栈
	 a = append(a, 6)
	fmt.Printf("入栈1, %v\n", a)
	a = append(a, 7)
	fmt.Printf("入栈2, %v\n", a)

	// 出栈
	a = a[:len(a) - 1]
	fmt.Printf("出栈1, %v\n", a)
	a = a[:len(a) - 1]
	fmt.Printf("出栈2, %v\n", a)
	a = a[:len(a) - 1]
	fmt.Printf("出栈3, %v\n", a)

}

func TestPush(t *testing.T) {
	fmt.Printf("stack : %v\n", stack)
	push(1)
	fmt.Printf("stack+ : %v\n", stack)
	push(2)
	fmt.Printf("stack+ : %v\n", stack)
	push(3)
	fmt.Printf("stack+ : %v\n", stack)
	pop()
	fmt.Printf("stack- : %v\n", stack)
	pop()
	fmt.Printf("stack- : %v\n", stack)
	pop()
	fmt.Printf("stack+ : %v\n", stack)
	pop()
	fmt.Printf("stack- : %v\n", stack)

}