/*
@Time : 2022/11/18 21:07
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : BinarySearchTree
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package BinarySearchTree

import "sync"

type node struct {
	left, right *node
	key uint
	value interface{}
}

func (n *node) remove() {
	n.key = 0
	n.value = nil
	n.left = nil
	n.right = nil
}

type BinarySearchTree struct {
	root *node
	size uint
	lock sync.Mutex
}

func NewBST() *BinarySearchTree {
	return new(BinarySearchTree)
}

// insert
func (t *BinarySearchTree) insert(at, new *node) {
	if new.key == at.key {
		return
	}
	if new.key < at.key {
		if at.left == nil {
			at.left = new
		} else {
			t.insert(at.left, new)
		}
	} else {
		if at.right == nil {
			at.right = new
		} else {
			t.insert(at.right, new)
		}
	}
}
