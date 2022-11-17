/*
@Time : 2022/11/16 21:44
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : radixSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package radixSort

func RadixSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 获取最大值
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}

	for bit := 1; bit < max/bit; bit *= 10 {
		arr = bitSort(arr, bit)
	}

	return arr
}

func bitSort(arr []int, bit int) []int {
	n := len(arr)
	// 统计各个位相同的数到bitCounts[]中
	bitCounts := make([]int, 11)
	for i := 0; i < n; i++ {
		num := (arr[i] / bit) % 10
		bitCounts[num]++
	}
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i - 1]
	}
	tmp := make([]int, 11)
	for i := n - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitCounts[num] - 1] = arr[i]
		bitCounts[num]--
	}
	for i := 0; i < n; i++ {
		arr[i] = tmp[i]
	}

	return arr
}