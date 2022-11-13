/*
@Time : 2022/11/13 19:53
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : bubbleSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package bubbleSort

func BubbleSortAsc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func BubbleSortDesc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}