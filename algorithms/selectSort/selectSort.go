/*
@Time : 2022/11/14 20:49
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : selectSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package selectSort

func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
