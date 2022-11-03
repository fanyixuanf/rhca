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

// 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
/**

slice 0xc0000103c0, [1 2 3 4 5 6]
slice 0xc0000103c0, [1 2 3 4 5 6]
test Slice: 0xc0000103c0, [1 2 3 4 5 6]

 */
func TestSlice(t *testing.T) {
	slice := []int{1,2,3,4,5,6}
	fmt.Printf("slice %p, %v\n", slice, slice)
	c := slice
	fmt.Printf("slice %p, %v\n", c, c)
	testSlice(c)
}

func testSlice(s []int) {
	fmt.Printf("test Slice: %p, %v\n", s, s)
}

/**

sa 0xc000016240, [1 2 3 4], cap: 4, len: 4
sa 0xc000014380, [1 2 3 4 5], cap: 8, len: 5
sb 0xc000014380, [1 2 3 4], cap: 8, len: 4
sa 0xc000014380, [1 2 3 4 5], cap: 8, len: 5
sa 0xc000014380, [1 2 4 5], cap: 8, len: 4

 */
func TestSlice2(t *testing.T) {
	var sa = []int{1, 2, 3, 4}
	fmt.Printf("sa %p, %v, cap: %d, len: %d\n", sa, sa, cap(sa), len(sa))
	// 添加元素在后面
	sa = append(sa, 5)
	fmt.Printf("sa %p, %v, cap: %d, len: %d\n", sa, sa, cap(sa), len(sa))
	// 删除最后一个元素个sb
	sb := sa[:len(sa) - 1]
	fmt.Printf("sb %p, %v, cap: %d, len: %d\n", sb, sb, cap(sb), len(sb))
	fmt.Printf("sa %p, %v, cap: %d, len: %d\n", sa, sa, cap(sa), len(sa))
	// 删除中间元素给sc
	sa = append(sa[:2], sa[3:]...)
	//fmt.Printf("sc %p, %v, cap: %d, len: %d\n", sc, sc, cap(sc), len(sc))
	fmt.Printf("sa %p, %v, cap: %d, len: %d\n", sa, sa, cap(sa), len(sa))
}