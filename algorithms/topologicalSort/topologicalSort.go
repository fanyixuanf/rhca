/*
@Time : 2022/11/17 21:04
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : topologicalSort
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package topologicalSort

import "fmt"

type graph struct {
	vertex int // 顶点
	list map[int][]int // 连接表边
}

func (g *graph) addVertex(t, s int) {
	g.list[t] = push(g.list[t], s)
}

// push
func push(list []int, value int) []int {
	return append(list, value)
}

// pop
func pop(arr []int) (int, []int) {
	if len(arr) > 0 {
		return arr[0], arr[1:]
	} else {
		return -1, arr
	}
}

func NewGraph(v int) *graph {
	g := new(graph)
	g.vertex = v
	g.list = map[int][]int{}
	i := 0
	for i < v {
		g.list[i] = make([]int, 0)
		i++
	}
	return g
}

func (g *graph) KhanSort() {
	var inDegree = make(map[int]int)
	var queue []int
	for i := 1; i < g.vertex; i++ {
		for _, m := range g.list[i] {
			inDegree[m]++
		}
	}
	for i := 1; i < g.vertex; i++ {
		if inDegree[i] == 0 {
			queue = push(queue, i)
		}
	}

	for len(queue) > 0 {
		var c int
		c, queue = pop(queue)
		fmt.Println("->", c)
		for _, k := range g.list[c] {
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = push(queue, k)
			}
		}
	}

}

func (g *graph) DfsSort() {
	inverseList := make(map[int][]int)
	for i := 1; i < g.vertex; i++ {
		for _, k := range g.list[i] {
			inverseList[k] = push(inverseList[k], i)
		}
	}

	visited := make([]bool, g.vertex + 1)
	visited[0] = true
	for i := 1; i < g.vertex; i++ {
		if visited[i] == false {
			visited[i] = true
			dfs(i, inverseList, visited)
		}
	}
}

func dfs(vertex int, inverseList map[int][]int, visited []bool) {
	for _, w := range inverseList[vertex] {
		if visited[w] == true {
			continue
		} else {
			visited[w] = true
			dfs(w, inverseList, visited)
		}
	}
	fmt.Println("->", vertex)
}