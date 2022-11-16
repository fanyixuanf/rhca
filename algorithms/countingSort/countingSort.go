/*
@Time : 2022/11/16 20:56
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : countingSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package countingSort

func CountingSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	// 默认最大值
	max := arr[0]
	// 找出最大值
	for i := 1; i < len; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	// 记录出现次数
	count := make([]int, max + 1)
	for i := 0; i < len; i++ {
		count[arr[i]]++
	}

	pos := 0
	for index, v := range count {
		for x := 0; x < v; x++ {
			arr[pos] = index
			pos++
		}
	}

	return arr
}
