/*
@Time : 2022/11/16 21:44
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : radixSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package radixSort

import (
	"fmt"
	"testing"
)

func TestSort1(t *testing.T) {
	var arr [3][]int
	myarr := []int{1, 2, 3, 1, 1, 2, 2, 2, 2, 2, 3}

	for i := 0; i < len(myarr); i++ {
		arr[myarr[i] - 1] = append(arr[myarr[i] - 1], myarr[i])
	}
	fmt.Println(arr)
}

func TestRadixSort(t *testing.T) {
	arr := []int{1, 2, 3, 1, 1, 2, 2, 2, 2, 2, 3}
	fmt.Println(RadixSort(arr))
}