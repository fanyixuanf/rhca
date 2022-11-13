/*
@Time : 2022/11/13 20:13
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : insertionSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package insertionSort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(InsertionSort(arr))
}
