/*
@Time : 2022/11/19 13:50
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : Graph_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	// 创建顶点
	vA := &Vertex{name: "A"}
	vB := &Vertex{name: "B"}
	vC := &Vertex{name: "C"}
	vD := &Vertex{name: "D"}
	vE := &Vertex{name: "E"}
	// 添加边
	vA.AddNeightbour(vB)
	vA.AddNeightbour(vD)
	vB.AddNeightbour(vC)
	vC.AddNeightbour(vD)
	vD.AddNeightbour(vE)
	// 创建图
	g := &Graph{}
	g.AddVertex(vA)
	g.AddVertex(vB)
	g.AddVertex(vC)
	g.AddVertex(vD)
	g.AddVertex(vE)
	// 深度优先搜索
	dfs(g, vA)

	vA.visited, vB.visited, vC.visited, vD.visited, vE.visited = false, false, false, false, false

	fmt.Println()
	// 广度优先
	bfs(g, vA)
}