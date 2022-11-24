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
	tree.Insert(8, "8")
	tree.Insert(4, "4")
	tree.Insert(10, "10")
	tree.Insert(2, "2")
	tree.Insert(6, "6")
	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(5, "5")
	tree.Insert(7, "7")
	tree.Insert(9, "9")
	//fmt.Println(tree.Search(2).value)
	//fmt.Println(tree.Len())
	tree.String()
}

func TestBinarySearchTree_InOrderTraversal(t *testing.T) {
	tree := NewBST()
	tree.Insert(8, "8")
	tree.Insert(4, "4")
	tree.Insert(10, "10")
	tree.Insert(2, "2")
	tree.Insert(6, "6")
	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(5, "5")
	tree.Insert(7, "7")
	tree.Insert(9, "9")
	a := ""
	tree.InOrderTraversal(func(n *node){
		a += fmt.Sprintf("%s\t", n.value)
	})
	println(a)
	tree.String()
}

func TestBinarySearchTree_BackOrderTraversal(t *testing.T) {
	tree := NewBST()
	tree.Insert(8, "8")
	tree.Insert(4, "4")
	tree.Insert(10, "10")
	tree.Insert(2, "2")
	tree.Insert(6, "6")
	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(5, "5")
	tree.Insert(7, "7")
	tree.Insert(9, "9")
	a := ""
	tree.BackOrderTraversal(func(n *node){
		a += fmt.Sprintf("%s\t", n.value)
	})
	println(a)
	tree.String()
}

func TestBinarySearchTree_FrontOrderTraversal(t *testing.T) {
	tree := NewBST()
	tree.Insert(8, "8")
	tree.Insert(4, "4")
	tree.Insert(10, "10")
	tree.Insert(2, "2")
	tree.Insert(6, "6")
	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(5, "5")
	tree.Insert(7, "7")
	tree.Insert(9, "9")
	a := ""
	tree.FrontOrderTraversal(func(n *node){
		a += fmt.Sprintf("%s\t", n.value)
	})
	println(a)
	tree.String()
}
