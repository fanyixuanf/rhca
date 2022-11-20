/*
@Time : 2022/11/20 13:41
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : BitHacking
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package BitHacking

import (
	"fmt"
	"math/rand"
)

func and() {
	var x uint8 = 0xAC
	fmt.Printf("%b\n", x)
	x = x & 0xF0
	fmt.Printf("%b\n", x)
}

func and1() {
	for x := 0; x < 10; x++ {
		num := rand.Int()
		if num & 1 == 1 {
			fmt.Printf("%v is odd\n", num)
		} else {
			fmt.Printf("%v is even\n", num)
		}
	}
}

func or() {
	var a uint8 = 0
	fmt.Printf("%b\n", 3)
	a |= 196
	fmt.Printf("%b\n", a)
	a |= 3
	fmt.Printf("%b\n", a)
}

func xor() {
	n, b := -12, 25
	fmt.Println("a and b have same sign?", (n ^ b) >= 0)
	var a byte = 0x0F
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", ^a)
}