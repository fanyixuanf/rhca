/*
@Time : 2022/11/16 20:11
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : shellSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package shellSort

func ShellSort(arr []int) []int {
	len := len(arr)
	if  len <= 1 {
		return arr
	}
	// 每次减半, 一直到step = 1
	for step := len/2; step >= 1; step /= 2 {
		// 插入排序, 步长为step
		for i := step; i < len; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入条件, 交换元素
				if arr[j + step] < arr[j] {
					arr[j], arr[j + step] = arr[j + step], arr[j]
					continue
				}
				break
			}
		}
	}
	return arr
}
