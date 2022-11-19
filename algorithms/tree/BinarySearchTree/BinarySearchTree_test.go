/*
@Time : 2022/11/18 21:08
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : BinarySearchTree_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package BinarySearchTree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	tree := NewBST()
	tree.Insert(1, "test1")
	tree.Insert(2, "test2")

	fmt.Println(tree.Search(2))
	fmt.Println(tree.Len())
}