/*
@Time : 2022/11/6 16:20
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : doubly-linked-list_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package doubly_linked_list

import (
	"container/list"
	"fmt"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	// Create a new list and put some numbers in it.
	l := list.New()
	fmt.Printf("List: %v\n", l)
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Printf("List: %v\n", l)
}