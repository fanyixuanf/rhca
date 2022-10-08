/*
@Time : 2022/10/7 17:25
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 简单工厂
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

// ----> 抽象层 <----
type Fruit interface {
	Show()
}

// ----> 实现层 <----
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("Apple")
}

type Banana struct {
	Fruit
}

func (b *Banana)Show() {
	fmt.Println("Banana")
}

type Pear struct {
	Fruit
}

func (pear *Pear)Show() {
	fmt.Println("Pear")
}

// ----> 工厂模块 <----
type Facory struct {

}

func (f *Facory)CreateFruit(name string) Fruit {
	var fruit Fruit
	if name == "apple" {
		fruit = new(Apple)
	} else if name == "banana" {
		fruit = new(Banana)
	} else if name == "pear" {
		fruit = new(Pear)
	}
	return fruit
}


// ----> 业务逻辑层 <----
func main() {
	factory := new(Facory)
	apple := factory.CreateFruit("apple")
	apple.Show()

	ba := factory.CreateFruit("banana")
	ba.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
