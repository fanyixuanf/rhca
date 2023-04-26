/*
@Time : 2022/11/19 13:42
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : BloomFilter_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package BloomFilter

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	// 创建布隆过滤器
	filter := NewBloomFilter(10000, 0.01)
	// 添加元素
	filter.Add([]byte("apple"))
	filter.Add([]byte("banana"))
	filter.Add([]byte("orange"))
	// 检查元素是否存在
	fmt.Println(filter.Test([]byte("apple")))   // true
	fmt.Println(filter.Test([]byte("banana")))  // true
	fmt.Println(filter.Test([]byte("orange")))  // true
	fmt.Println(filter.Test([]byte("pineapple")))  // false
}
