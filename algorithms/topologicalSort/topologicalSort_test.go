/*
@Time : 2022/11/17 21:04
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : topologicalSort_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package topologicalSort

import (
	"testing"
)

func TestGraph_DfsSort(t *testing.T) {
	g := NewGraph(8)
	g.addVertex(2, 1)
	g.addVertex(3, 1)
	g.addVertex(7, 1)
	g.addVertex(4, 2)
	g.addVertex(5, 2)
	g.addVertex(6, 7)
	g.DfsSort()
	//g.KhanSort()
}