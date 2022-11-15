/*
@Time : 2022/11/15 21:03
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : mergeSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package mergeSort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(MergeSort(arr))
}
