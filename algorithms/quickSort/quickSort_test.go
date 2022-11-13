/*
@Time : 2022/11/13 17:01
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : quickSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package quickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}
