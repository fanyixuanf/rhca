/*
@Time : 2023/5/3 17:11
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : linearSearch
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package demo

func linearSearch(array []int, target int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return i
		}
	}
	return -1
}
