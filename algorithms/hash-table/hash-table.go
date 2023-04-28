/*
@Time : 2022/11/10 20:28
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : hash-table
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package hash_table

import "hash/fnv"

// Node 表示哈希表中的一个节点
type Node struct {
	Key   string      // 键
	Value interface{} // 值
	Next  *Node       // 链表中下一个节点的指针
}

// HashTable 表示哈希表
type HashTable struct {
	Table []*Node // 存储哈希表数据的数组
	Size  int     // 哈希表的大小
}

// NewHashTable 根据指定的大小创建一个新的哈希表
func NewHashTable(size int) *HashTable {
	return &HashTable{
		Table: make([]*Node, size),
		Size:  size,
	}
}

// hash 根据键值计算哈希值
func (h *HashTable) hash(key string) int {
	hval := fnv.New32()
	hval.Write([]byte(key))
	return int(hval.Sum32()) % h.Size
}

// Put 向哈希表中插入一个键值对
func (h *HashTable) Put(key string, value interface{}) {
	index := h.hash(key)
	if h.Table[index] == nil { // 若当前位置没有节点，则直接创建节点
		h.Table[index] = &Node{Key: key, Value: value}
	} else { // 否则在当前位置的链表中插入节点
		curr := h.Table[index]
		for curr.Next != nil { // 寻找链表末尾
			curr = curr.Next
		}
		curr.Next = &Node{Key: key, Value: value}
	}
}

// Get 根据键从哈希表中获取对应的值
func (h *HashTable) Get(key string) (interface{}, bool) {
	index := h.hash(key)
	curr := h.Table[index]
	for curr != nil {
		if curr.Key == key {
			return curr.Value, true
		}
		curr = curr.Next // 更新节点指针
	}
	return nil, false
}

// Remove 从哈希表中删除指定键对应的节点
func (h *HashTable) Remove(key string) bool {
	index := h.hash(key)
	curr := h.Table[index]
	if curr == nil {
		return false // 没有找到对应的节点
	}
	if curr.Key == key { // 若第一个节点就是需要删除的节点
		h.Table[index] = curr.Next
		return true
	}
	prev := curr
	curr = curr.Next
	for curr != nil {
		if curr.Key == key {
			prev.Next = curr.Next
			return true
		}
		prev = curr
		curr = curr.Next // 更新节点指针
	}
	return false // 没有找到对应的节点
}
