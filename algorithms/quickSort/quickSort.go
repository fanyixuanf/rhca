/*
@Time : 2022/11/13 17:00
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : quickSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package quickSort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	partition := arr[0]

	left := make([]int, 0, 0)
	right := make([]int, 0, 0)
	mid := make([]int, 0, 0)

	mid = append(mid, partition)

	for i := 1; i < len(arr); i++ {
		if arr[i] < partition {
			left = append(left, arr[i])
		} else if arr[i] > partition {
			right = append(right, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	left, right = QuickSort(left), QuickSort(right)
	return append(append(left, mid...), right...)
}