/*
@Time : 2022/11/6 15:05
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : linked-list
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	headNode *Node
}

func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	} else {
		return false
	}
}

func (l *List) Length() int {
	cur := l.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// 头部添加元素
func (l *List) Add(data int) {
	node := &Node{
		Data: data,
		Next: l.headNode,
	}
	l.headNode = node
}

// 尾部添加元素

func (l *List) Append(data int) {
	node := &Node{
		Data: data,
	}
	if l.IsEmpty() {
		l.headNode = node
	} else {
		cur := l.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

// 任意位置插入
func (l *List) Insert(index, data int) {
	if index < 0 {
		l.Add(data)
	} else if index > l.Length() {
		l.Append(data)
	} else {
		count := 0
		cur := l.headNode
		for count < index - 1 {
			cur = cur.Next
			count++
		}
		node := &Node{
			Data: data,
		}
		node.Next = cur.Next
		cur.Next = node
	}
}

func (l *List) Print() {
	for l.headNode != nil {
		fmt.Printf("list value: %d, pointer: %p \n",l.headNode.Data, l.headNode)
		l.headNode = l.headNode.Next
	}
}