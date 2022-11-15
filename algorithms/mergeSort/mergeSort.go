/*
@Time : 2022/11/15 21:03
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : mergeSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package mergeSort

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	p := len(arr) / 2

	left := MergeSort(arr[:p])
	right := MergeSort(arr[p:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	l, r := len(left), len(right)
	var ret []int
	for {
		if i >= l || j >= r {
			break
		}
		if left[i] <= right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}

	if i != l {
		for ; i < l; i++ {
			ret = append(ret, left[i])
		}
	}

	if j != r {
		for ; j < r; j++ {
			ret = append(ret, right[j])
		}
	}

	return ret
}