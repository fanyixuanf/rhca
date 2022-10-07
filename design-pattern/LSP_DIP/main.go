/*
@Time : 2022/10/7 16:18
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : main
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

// 耦合度极高的设计

//type Benz struct {
//
//}
//
//func (b *Benz) Run () {
//	fmt.Println("Benz run...")
//}
//
//type BMW struct {
//
//}
//
//func (b *BMW) Run() {
//	fmt.Println("BMW run...")
//}
//
//type zhang3 struct {
//
//}
//
//func (z3 *zhang3) DriveBenz(benz *Benz) {
//	benz.Run()
//	fmt.Println("zhang3 Benz....")
//}
//
//// +++++
//func (z3 *zhang3) DriveBmw(bmw *BMW) {
//	bmw.Run()
//	fmt.Println("zhang3 BMW....")
//}
//
//type li4 struct {
//
//}
//
//func (l *li4) DirveBMW(bmw *BMW) {
//	bmw.Run()
//	fmt.Println("li4 BMW...")
//}
//
//// +++++
//func (l *li4) DirveBenz(benz *Benz) {
//	benz.Run()
//	fmt.Println("li4 Benz...")
//}
//
//func main() {
//	benz := &Benz{}
//	z3 := &zhang3{}
//	z3.DriveBenz(benz)
//	li4 := &li4{}
//	bmw := &BMW{}
//	li4.DirveBMW(bmw)
//
//	z3.DriveBmw(bmw)
//	li4.DirveBenz(benz)
//}

// 依赖倒转

// 抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层

type Benz struct {

}

func (b *Benz) Run () {
	fmt.Println("Benz running...")
}

type BMW struct {

}

func (bm *BMW) Run() {
	fmt.Println("BMW running...")
}

// ++++++++


type Zhang3 struct {

}

func (z3 *Zhang3)Drive(car Car) {
	fmt.Println("z3 drive car")
	car.Run()
}

type Li4 struct {

}

func (l4 *Li4) Drive(car Car) {
	fmt.Println("l4 drive car")
	car.Run()
}

// +++++++++++++

type Wang5 struct {

}

func (w *Wang5) Drive(car Car) {
	fmt.Println("wang5 drive car")
	car.Run()
}



// 业务逻辑层
func main() {
	var benz Car
	benz = new(Benz)
	var z3 Driver
	z3 = new(Zhang3)

	z3.Drive(benz)

	var bmw Car
	bmw = new(BMW)
	var l4 Driver
	l4 = new(Li4)
	l4.Drive(bmw)

	// ++++
	z3.Drive(bmw)

	var wang5 Driver
	wang5 = new(Wang5)
	wang5.Drive(bmw)
}