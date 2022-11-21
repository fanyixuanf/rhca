/*
@Time : 2022/11/18 21:07
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : BinarySearchTree
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package BinarySearchTree

import (
	"fmt"
	"sync"
)

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

// 中序遍历 => 左中右
func (t *BinarySearchTree) InOrderTraversal(handler func(*node)) {
	if t.root != nil {
		t.inOrderTraversal(t.root, handler)
	}
}

func (t *BinarySearchTree) inOrderTraversal(at *node, handler func(*node)) {
	if at.left != nil {
		t.inOrderTraversal(at.left, handler)
	}
	handler(at)
	if at.right != nil {
		t.inOrderTraversal(at.right, handler)
	}
}

// 前序遍历 中左右
func (t *BinarySearchTree)FrontOrderTraversal(handler func(*node)) {
	if t.root != nil {
		t.frontOrderTraversal(t.root, handler)
	}
}

func (t *BinarySearchTree)frontOrderTraversal(at *node, handler func(*node)) {
	handler(at)
	if at.left != nil {
		t.frontOrderTraversal(at.left, handler)
	}
	if at.right != nil {
		t.frontOrderTraversal(at.right, handler)
	}
}

// 后续遍历 左右中
func (t *BinarySearchTree) BackOrderTraversal(handler func(*node)) {
	if t.root != nil {
		t.backOrderTraversal(t.root, handler)
	}
}

func (t *BinarySearchTree)backOrderTraversal(at *node, handler func(*node)) {
	if at.left != nil {
		t.backOrderTraversal(at.left, handler)
	}
	if at.right != nil {
		t.backOrderTraversal(at.right, handler)
	}
	handler(at)
}

// 打印tree
func (t *BinarySearchTree) String() {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	stringify(t.root, 0)
	fmt.Println("-------------------")
}

func stringify(node *node, level int) {
	if node == nil {
		return
	}
	format := ""
	for i := 0; i < level; i++ {
		format += "\t"
	}
	format += "----[ "
	level++
	stringify(node.left, level)
	fmt.Printf(format + "%d\n", node.key)
	stringify(node.right, level)
}
