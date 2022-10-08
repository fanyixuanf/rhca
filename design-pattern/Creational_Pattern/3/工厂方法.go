/*
@Time : 2022/10/8 21:43
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 工厂方法
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

type AbstractFactory interface {
	CreateFruit() Fruit
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

// +++++
type Oooo struct {

}

func (o *Oooo)Show() {
	fmt.Println("Oooo")
}

// 工厂
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var f Fruit
	f = new(Apple)
	return f
}

type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var f Fruit
	f = new(Banana)
	return f
}

type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var f Fruit
	f = new(Pear)
	return f
}

// ++++++
type OoooFactory struct {}

func (oo *OoooFactory) CreateFruit() Fruit {
	return new(Oooo)
}

// ----> 业务逻辑层 <----
func main () {
	var fapp AbstractFactory
	fapp = new(AppleFactory)
	var app Fruit
	app = fapp.CreateFruit()
	app.Show()

	fbana := new(BananaFactory)
	bana := fbana.CreateFruit()
	bana.Show()

	var fpear AbstractFactory
	fpear = new(PearFactory)
	var pear Fruit
	pear = fpear.CreateFruit()
	pear.Show()

	var foo AbstractFactory
	foo = new(OoooFactory)
	var oo Fruit
	oo = foo.CreateFruit()
	oo.Show()

}
