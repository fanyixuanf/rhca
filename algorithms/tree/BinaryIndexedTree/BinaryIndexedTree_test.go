/*
@Time : 2022/11/19 12:13
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : BinaryIndexedTree_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package BinaryIndexedTree

import (
	"fmt"
	"testing"
)

func TestBinaryIndexedTree_Init(t *testing.T) {
	var arr = []int{2,4,6,8,10}
	var tree BinaryIndexedTree
	tree.Init(arr)
	fmt.Println(tree)
}
