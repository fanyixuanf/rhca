/*
@Time : 2023/4/26 21:47
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : list_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package List

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	n1 := Node{val: 1}
	n2 := Node{val: 2}
	n3 := Node{val: 3}

	n1.Next = &n2
	n2.Next = &n3

	printList(&n1)
	fmt.Println()
	r := reverseList(&n1)
	printList(r)
}