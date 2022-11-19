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
func (t *BinarySearchTree) Insert(key uint, val interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	nd := new(node)
	nd.key = key
	nd.value = val
	nd.left = nil
	nd.right = nil
	if t.root == nil {
		t.root = nd
	} else {
		t.insertNode(t.root, nd)
	}
	t.size++
}

func (t *BinarySearchTree) insertNode(at, new *node) {
	if new.key == at.key {
		return
	}
	if new.key < at.key {
		if at.left == nil {
			at.left = new
		} else {
			t.insertNode(at.left, new)
		}
	} else {
		if at.right == nil {
			at.right = new
		} else {
			t.insertNode(at.right, new)
		}
	}
}

// Search
func (t *BinarySearchTree)Search(key uint) *node {
	if t.root != nil {
		return t.searchNode(key, t.root)
	}
	return nil
}

func (t *BinarySearchTree) searchNode(key uint, at *node) *node {
	if at.key == key {
		return at
	} else if at.key > key {
		if at.left != nil {
			return t.searchNode(key, at.left)
		}
	} else {
		if at.right != nil {
			return t.searchNode(key, at.right)
		}
	}
	return nil
}

// 返回节点数
func (t *BinarySearchTree) Len() uint {
	return t.size
}

