/*
@Time : 2022/11/14 20:49
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : selectSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package selectSort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(SelectSort(arr))
}