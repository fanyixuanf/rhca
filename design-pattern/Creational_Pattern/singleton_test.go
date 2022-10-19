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