/*
@Time : 2022/11/13 20:13
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : insertionSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package insertionSort

func InsertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		for j := i - 1; j >= 0 ; j-- {
			if arr[j] > v {
				arr[j + 1] = arr[j]
			} else {
				arr[j + 1] = v
				break
			}
		}
	}
	return arr
}
