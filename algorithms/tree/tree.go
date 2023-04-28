/*
@Time : 2022/11/18 20:59
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : tree
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package tree

import (
"container/list"
"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}
type Tree struct {
	Root *Node
}
// 新建一个树
func NewTree() *Tree {
	return &Tree{}
}
// 添加新节点
func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &Node{Value: val}
		return
	}
	queue := list.New()
	queue.PushBack(t.Root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*Node)
		if node.Left == nil {
			node.Left = &Node{Value: val}
			return
		}
		if node.Right == nil {
			node.Right = &Node{Value: val}
			return
		}
		queue.PushBack(node.Left)
		queue.PushBack(node.Right)
	}
}
// 删除节点
func (t *Tree) DeleteNode(val int) {
	if t.Root == nil {
		return
	}
	if t.Root.Value == val {
		t.Root = nil
		return
	}
	parent, node := t.find(val)
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if parent.Left == node {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	}
	var replace *Node
	if node.Left != nil && node.Right != nil {
		replace = node.Right
		for replace.Left != nil {
			replace = replace.Left
		}
		node.Value = replace.Value
		val = replace.Value
		parent, node = t.find(val)
	}
	if node.Left == nil || node.Right == nil {
		var child *Node
		if node.Left != nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if parent.Left == node {
			parent.Left = child
		} else {
			parent.Right = child
		}
		return
	}
}
// 查找节点并返回父节点
func (t *Tree) find(val int) (*Node, *Node) {
	if t.Root == nil {
		return nil, nil
	}
	if t.Root.Value == val {
		return nil, t.Root
	}
	queue := list.New()
	queue.PushBack(t.Root)
	for queue.Len() > 0 {
		parent := queue.Front()
		node := parent.Value.(*Node)
		queue.Remove(parent)
		if node.Left != nil {
			if node.Left.Value == val {
				return node, node.Left
			}
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			if node.Right.Value == val {
				return node, node.Right
			}
			queue.PushBack(node.Right)
		}
	}
	return nil, nil
}
// dfs 深度优先搜索算法
func (t *Tree) dfs(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Value) // 先访问根节点
	t.dfs(root.Left)              // 再遍历左子树
	t.dfs(root.Right)             // 最后遍历右子树
}
// bfs 广度优先搜索算法
func (t *Tree) bfs(root *Node) {
	if root == nil {
		return
	}
	queue := list.New() // 使用列表模拟队列
	queue.PushBack(root) // 将根节点入队
	for queue.Len() > 0 { // 队列不为空时循环
		node := queue.Remove(queue.Front()).(*Node) // 取出队首节点
		fmt.Printf("%d ", node.Value)               // 访问节点
		if node.Left != nil {                       // 左子节点入队
			queue.PushBack(node.Left)
		}
		if node.Right != nil { // 右子节点入队
			queue.PushBack(node.Right)
		}
	}
}
