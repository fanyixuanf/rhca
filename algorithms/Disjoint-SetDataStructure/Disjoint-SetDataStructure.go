/*
@Time : 2022/11/19 12:29
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : Disjoint-SetDataStructure
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Disjoint_SetDataStructure

type UnionFind struct {
	parent []int
	count int
}

// 初始化, 每个节点创建集合
func (uf *UnionFind) Init (n int) {
	uf.parent = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	uf.count = n
}

// 找到集合所在的根节点
func (uf *UnionFind) Find(p int) int {
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}

func (uf *UnionFind) Union(p, q int) {
	rootP, rootQ := uf.Find(p), uf.Find(q)
	if rootP == rootQ {
		return
	}
	uf.parent[rootP] = rootQ
	uf.count--
}

func (uf *UnionFind) Count() int {
	return uf.count
}
