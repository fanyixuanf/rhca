/*
@Time : 2022/11/16 20:56
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : countingSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package countingSort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(CountingSort(arr))
}
