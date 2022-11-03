/*
@Time : 2022/11/3 21:04
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : stack
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package stack

var stack []int

func push(i int) {
	stack = append(stack, i)
}

func pop() {
	if len(stack) == 0 {
		return
	}
	stack = stack[:len(stack) - 1]
}
