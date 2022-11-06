/*
@Time : 2022/11/6 15:05
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : linked-list_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package linked_list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	var list List
	list.Add(1)
	list.Append(2)
	list.Append(3)
	list.Add(4)
	list.Insert(1, 5)
	list.Print()
}