/*
@Time : 2022/11/12 22:18
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : topk
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package typical

import (
	"fmt"
	"testing"
)

// 取出 arr{5,8,2,3,6,4,1,7,9,4,6,2}中最大的5个数

func TestTopK1(t *testing.T) {
	arr := [12]int{5,8,2,3,6,4,1,7,9,4,6,2}
	fmt.Println(arr)

}

func TestForeach(t *testing.T) {

	year := 22
	month := 8
	for {
		if month <= 0 {
			month = 12
			year--
		}
		if year == 19 {
			break
		}
		fmt.Println(year, "-------", month)
		for i := year - 1; i < year + 1; i++ {
			if i < year {
				for j := month; j < 13; j++ {
					fmt.Println(i, j)
				}
			} else if i == year {
				for j := 1; j < month; j++ {
					fmt.Println(i, j)
				}
			}
		}
		month--
	}

}
