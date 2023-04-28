/*
@Time : 2022/11/19 13:50
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : Graph
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Graph

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	name     string
	visited  bool
	neightbours []*Vertex
}
// 添加顶点
func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}
// 添加边
func (v *Vertex) AddNeightbour(neighbour *Vertex) {
	v.neightbours = append(v.neightbours, neighbour)
}
// 栈的数据结构
type Stack struct {
	items []*Vertex
}
// 入栈
func (s *Stack) Push(v *Vertex) {
	s.items = append(s.items, v)
}
// 出栈
func (s *Stack) Pop() *Vertex {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
// 图的深度优先遍历
func dfs(g *Graph, start *Vertex) {
	stack := &Stack{}
	stack.Push(start)
	for !stack.IsEmpty() {
		currentVertex := stack.Pop()
		if !currentVertex.visited {
			fmt.Printf("%s ", currentVertex.name)
			currentVertex.visited = true
		}
		for _, neighbour := range currentVertex.neightbours {
			if !neighbour.visited {
				stack.Push(neighbour)
			}
		}
	}
}


// 队列的数据结构
type Queue struct {
	items []*Vertex
}
// 入队
func (q *Queue) Enqueue(v *Vertex) {
	q.items = append(q.items, v)
}
// 出队
func (q *Queue) Dequeue() *Vertex {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
// 图的广度优先遍历
func bfs(g *Graph, start *Vertex) {
	queue := &Queue{}
	queue.Enqueue(start)
	start.visited = true
	for !queue.IsEmpty() {
		currentVertex := queue.Dequeue()
		fmt.Printf("%s ", currentVertex.name)
		for _, neighbour := range currentVertex.neightbours {
			if !neighbour.visited {
				queue.Enqueue(neighbour)
				neighbour.visited = true
			}
		}
	}
}
