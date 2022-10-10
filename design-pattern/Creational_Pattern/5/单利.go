/*
@Time : 2022/10/10 20:54
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 单利
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import (
	"fmt"
	"sync"
)

//// 非公有的
//type singelton struct {}
//// 指针不能修改
//var instance *singelton = new(singelton)
//
//// 不能变成成员方法, 逻辑矛盾了
//func GetInstance() *singelton {
//	return instance
//}
//
//func (s *singelton)Some() {
//	fmt.Println("singelton ....")
//}
//
//func main() {
//	f := GetInstance()
//	f.Some()
//}

type singelton struct {}
var instance *singelton
var lock sync.Mutex

var initialized uint32

//func GetInstance() *singelton {
//	// 性能问题
//	//lock.Lock()
//	//defer lock.Unlock()
//
//	if atomic.LoadUint32(&initialized) == 1 {
//		return instance
//	}
//
//	lock.Lock()
//	defer lock.Unlock()
//
//	if instance == nil {
//		instance = new(singelton)
//		atomic.StoreUint32(&initialized, 1)
//	}
//
//	return instance
//}
var once sync.Once
// once 线程安全的
func GetInstance() *singelton {
	once.Do(func() {
		instance = new(singelton)
	})
	return instance
}

func (s *singelton) Something() {
	fmt.Println("Something ....")
}

func main() {
	s1 := GetInstance()
	s1.Something()
}
