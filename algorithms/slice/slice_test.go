/*
@Time : 2022/11/2 21:44
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : slice_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package slice

import (
	"fmt"
	"testing"
)

// Go 中数组赋值和函数传参都是值复制的
/**
	arrayA : 0xc00000c370 , [100 200]
	arrayB : 0xc00000c380 , [100 200]
	func Array : 0xc00000c3b0 , [100 200]
 */
func TestArray1(t *testing.T) {

	arrayA := [2]int{100, 200}
	var arrayB [2]int
	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

/**

func Array : 0xc00009e2a0 , [100 200]
func Array : 0xc00009e2c0 , [100 300]
arrayA : 0xc00009e2c0 , [100 400]

 */
func TestArray2(t *testing.T) {
	arrayA := [2]int{100, 200}
	testArrayPoint(&arrayA)
	arrayB := arrayA
	testArrayPoint(&arrayB)
	fmt.Printf("arrayA : %p , %v\n", &arrayB, arrayB)
}

func testArrayPoint(a *[2]int) {
	fmt.Printf("func Array : %p , %v\n", a, *a)
	(*a)[1] += 100
}