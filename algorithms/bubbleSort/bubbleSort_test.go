/*
@Time : 2022/11/13 19:53
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : bubbleSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package bubbleSort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(BubbleSortAsc(arr))
	fmt.Println("------------")
	fmt.Println(BubbleSortDesc(arr))
}
