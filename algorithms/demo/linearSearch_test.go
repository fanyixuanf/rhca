/*
@Time : 2023/5/3 17:13
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : linearSearch_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package demo

import (
	"fmt"
	"testing"
)

func TestLinsearSearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	target := 4

	index := linearSearch(numbers, target)
	if index != -1 {
		fmt.Printf("Target found at index %d\n", index)
	} else {
		fmt.Println("Target not found")
	}
}
