/*
@Time : 2022/11/5 17:42
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : heap.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package heap

// min heap
type IntHeap []int

func (h IntHeap) Len () int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] <= h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


// max heap

type MaxHeap []int

func (m MaxHeap) Len () int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] >= m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
