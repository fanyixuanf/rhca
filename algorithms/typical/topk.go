/*
@Time : 2023/4/28 21:40
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : topk
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package typical

import (
	"math/rand"
	"time"
)

// partition 将数组划分为小于 pivot 和大于等于 pivot 两部分，返回 pivot 最终的位置
func partition(arr []int, left, right, pivotIndex int) int {
	pivotVal := arr[pivotIndex]
	// 交换 pivotVal 和 right 位置上的元素，以便将 pivot 移动到最右侧
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	storeIndex := left
	for i := left; i < right; i++ {
		if arr[i] > pivotVal {
			arr[i], arr[storeIndex] = arr[storeIndex], arr[i]
			storeIndex++
		}
	}
	// 将 pivot 移回其最终的位置
	arr[right], arr[storeIndex] = arr[storeIndex], arr[right]

	return storeIndex
}

// randomSelect 使用随机选择算法查找数组中第 K 大的元素
func randomSelect(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}

	// 选取一个随机 pivot
	rand.Seed(time.Now().UnixNano())
	pivotIndex := left + rand.Intn(right-left)

	// 对数组进行 partition
	pivotIndex = partition(arr, left, right, pivotIndex)

	// 根据 pivot 的位置判断所需的下一步操作
	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return randomSelect(arr, left, pivotIndex-1, k)
	} else {
		return randomSelect(arr, pivotIndex+1, right, k)
	}
}

// topK 使用 randomSelect 实现 topK 问题
func topK(arr []int, k int) []int {
	if k <= 0 || k > len(arr) {
		return []int{}
	}

	// 查找第 K 大的元素
	pivotVal := randomSelect(arr, 0, len(arr)-1, k-1)

	// 从数组中筛选出前 K 大的元素
	topKArr := make([]int, k)
	topKIndex := 0
	for _, val := range arr {
		if val >= pivotVal {
			topKArr[topKIndex] = val
			topKIndex++
			if topKIndex >= k {
				break
			}
		}
	}

	return topKArr
}

// 用数组第一个元素分组

// partition 将数组划分为小于 pivot 和大于等于 pivot 两部分，返回 pivot 最终的位置
func partition1(arr []int, left, right int) int {
	pivotVal := arr[left]
	storeIndex := left + 1
	for i := left + 1; i <= right; i++ {
		if arr[i] > pivotVal {
			arr[i], arr[storeIndex] = arr[storeIndex], arr[i]
			storeIndex++
		}
	}
	// 将 pivot 移回其最终的位置
	pivotIndex := storeIndex - 1
	arr[left], arr[pivotIndex] = arr[pivotIndex], arr[left]

	return pivotIndex
}

// randomSelect 使用随机选择算法查找数组中第 K 大的元素
func randomSelect1(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}

	// 选取一个随机 pivot
	randIndex := rand.Intn(right-left+1) + left
	arr[left], arr[randIndex] = arr[randIndex], arr[left]

	// 对数组进行 partition
	pivotIndex := partition1(arr, left, right)

	// 根据 pivot 的位置判断所需的下一步操作
	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return randomSelect1(arr, left, pivotIndex-1, k)
	} else {
		return randomSelect1(arr, pivotIndex+1, right, k)
	}
}

// topK 使用 randomSelect 实现 topK 问题
func topK1(arr []int, k int) []int {
	if k <= 0 || k > len(arr) {
		return []int{}
	}

	// 查找第 K 大的元素
	pivotVal := randomSelect1(arr, 0, len(arr)-1, k-1)

	// 从数组中筛选出前 K 大的元素
	topKArr := make([]int, k)
	topKIndex := 0
	for _, val := range arr {
		if val >= pivotVal {
			topKArr[topKIndex] = val
			topKIndex++
			if topKIndex >= k {
				break
			}
		}
	}

	return topKArr
}
