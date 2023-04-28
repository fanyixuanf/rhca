/*
@Time : 2022/11/18 21:00
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : tree_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	// 构造一棵二叉树
	tree := NewTree()
	tree.AddNode(1)
	tree.AddNode(2)
	tree.AddNode(3)
	tree.AddNode(4)
	tree.AddNode(5)
	tree.AddNode(6)
	tree.AddNode(7)
	// 删除节点5
	//tree.DeleteNode(5)
	// 深度优先搜索
	fmt.Println("DFS:")
	tree.dfs(tree.Root)
	fmt.Println()
	// 广度优先搜索
	fmt.Println("BFS:")
	tree.bfs(tree.Root)
	fmt.Println()
}