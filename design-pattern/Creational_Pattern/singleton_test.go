/*
@Time : 2022/10/19 21:48
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : singleton_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package creation_pattern

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type singleton struct {}

var instance *singleton

var once sync.Once

func GetInstance() *singleton{
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) DoSomething() {
	fmt.Println("singleton -> DoSomething ....")
}

func TestSingleton(t *testing.T) {
	s1 := GetInstance()
	s1.DoSomething()
}

var lock sync.Mutex
var initialized uint32
func GetInstance1() *singleton{
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(singleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func TestSingleton1(t *testing.T) {
	s1 := GetInstance1()
	s1.DoSomething()
}