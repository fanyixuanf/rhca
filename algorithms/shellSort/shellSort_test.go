/*
@Time : 2022/11/16 20:11
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : shellSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package shellSort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(ShellSort(arr))
}