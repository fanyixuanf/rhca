/*
@Time : 2022/11/19 13:42
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : BloomFilter
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package BloomFilter

import (
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	size uint
	bits []bool
	funcs []func([]byte) uint32
}

func NewBloomFilter(items uint, falsePositiveRate float64) *BloomFilter {
	// 计算位数组大小
	size := math.Ceil(-float64(items) * math.Log(falsePositiveRate) / math.Pow(math.Log(2), 2))
	// 计算哈希函数数量
	n := math.Ceil(math.Log(2) * size / float64(items))
	// 初始化布隆过滤器
	filter := &BloomFilter{
		size:  uint(size),
		bits:  make([]bool, int(size)),
		funcs: make([]func([]byte) uint32, int(n)),
	}
	// 设置哈希函数
	for i := 0; i < int(n); i++ {
		filter.funcs[i] = hashFunc(i)
	}
	return filter
}

func (b *BloomFilter) Add(data []byte) {
	for _, fn := range b.funcs {
		i := fn(data) % uint32(b.size) // 根据哈希函数计算哈希值
		b.bits[i] = true // 将对应的位设为 true
	}
}

func (b *BloomFilter) Test(data []byte) bool {
	for _, fn := range b.funcs {
		i := fn(data) % uint32(b.size) // 根据哈希函数计算哈希值
		if b.bits[i] == false {
			return false // 如果对应位为 false，说明元素不存在
		}
	}
	return true
}

// 哈希函数
func hashFunc(seed int) func([]byte) uint32 {
	return func(data []byte) uint32 {
		h := fnv.New32a()
		h.Write(data)
		return uint32(seed) * h.Sum32()
	}
}
