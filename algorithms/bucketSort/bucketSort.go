/*
@Time : 2022/11/17 20:36
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : bucketSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package bucketSort

import (
	"rhca/algorithms/insertionSort"
)

func BucketSort(arr []int, size int) []int {
	alen := len(arr)
	if alen <= 1 {
		return arr
	}
	min := arr[0]
	max := arr[0]

	for i := 0; i < alen; i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}

	// 桶初始化
	count := make([][]int, (max-min)/size + 1)

	// 数据插入桶中
	for i := 0; i < alen; i++ {
		count[(arr[i] - min)/size] = append(count[(arr[i] - min)/size], arr[i])
	}
	// 排序
	k := 0
	for _, bucket := range count {
		if len(bucket) <= 0 {
			continue
		}
		// 插入排序
		insertionSort.InsertionSort(bucket)
		for _, v := range bucket {
			arr[k] = v
			k += 1
		}
	}

	return arr
}