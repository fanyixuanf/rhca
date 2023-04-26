/*
@Time : 2023/4/26 21:47
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : list.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

type ListNode struct {
	val int
	Next *ListNode
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d", l.val)
		l = l.Next
	}
}

func reverseList(l *ListNode) *ListNode {
	var prev *ListNode
	for l != nil {
		curr := l
		l = l.Next
		curr.Next = prev
		prev = curr
	}
	return prev
}