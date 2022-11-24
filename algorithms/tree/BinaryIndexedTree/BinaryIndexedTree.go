/*
@Time : 2022/11/19 12:13
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : BinaryIndexedTree
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package BinaryIndexedTree

type BinaryIndexedTree struct {
	tree []int
	capacity int
}

func (bit *BinaryIndexedTree) Init(nums []int) {
	bit.tree, bit.capacity = make([]int, len(nums) + 1), len(nums) + 1
	for i := 1; i < len(nums); i++ {
		bit.tree[i] += nums[i - 1]
		for j := i - 2; j >= i - lowbit(i); j-- {
			bit.tree[i] += nums[j]
		}
	}
}

func lowbit(x int) int {
	return x & -x
}