/*
@Time : 2022/11/17 20:36
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : bucketSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package bucketSort

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(BucketSort(arr, 5))
}