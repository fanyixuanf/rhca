/*
@Time : 2022/11/19 12:30
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : Disjoint-SetDataStructure_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Disjoint_SetDataStructure

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	var n, m int
	fmt.Scan(&n, &m)

	uf := UnionFind{}
	uf.Init(n)

	for i := 0; i < m; i++ {
		var p, q int
		fmt.Scan(&p, &q)
		uf.Union(p, q)
	}

	fmt.Println(uf.count, "components")
}