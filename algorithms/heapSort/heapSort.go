/*
@Time : 2022/11/14 21:10
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : heapSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package heapSort

func HeapSort(arr []int) []int {

	length := len(arr)
	for i := 0; i < length; i++ {
		last := length - 1
		heapsotMax(arr, last)
		if i < length {
			arr[0], arr[last - 1] = arr[last - 1], arr[0]
		}
	}
	return arr
}

func heapsotMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}

	depth := length/2 - 1 // 深度

	for i := depth; i >= 0; i-- {
		topMax := i
		left := 2 * i + 1
		right := 2 * i + 2
		if left <= length - 1 && arr[left] > arr[topMax] {
			topMax = left
		}
		if right <= length - 1 && arr[right] > arr[topMax] {
			topMax = right
		}

		if topMax != i {
			arr[i], arr[topMax] = arr[topMax], arr[i]
		}
	}
	return arr
}
