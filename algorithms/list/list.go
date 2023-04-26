/*
@Time : 2023/4/26 21:47
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : list.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package List

import "fmt"

type Node struct {
	val int
	Next *Node
}

func printList(l *Node) {
	for l != nil {
		fmt.Printf("%d", l.val)
		l = l.Next
	}
}

func reverseList(l *Node) *Node {
	var prev *Node
	for l != nil {
		curr := l
		l = l.Next
		curr.Next = prev
		prev = curr
	}
	return prev
}