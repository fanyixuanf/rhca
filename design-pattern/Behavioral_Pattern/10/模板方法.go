/*
@Time : 2022/10/11 22:02
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 模板方法
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

// 抽象层
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()
	WantAdd() bool
}

type Template struct {
	b Beverage
}

func (t *Template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAdd() == true {
		t.b.AddThings()
	}
}

// 实现层
type MakeCoffee struct {
	Template
}

func (mc *MakeCoffee)BoilWater() {
	fmt.Println("MakeCoffee BoilWater ...")
}

func (mc *MakeCoffee)Brew() {
	fmt.Println("MakeCoffee Brew ...")
}

func (mc *MakeCoffee)PourInCup() {
	fmt.Println("MakeCoffee PourInCup ...")
}

func (mc *MakeCoffee)AddThings() {
	fmt.Println("MakeCoffee AddThings ...")
}

func (mc *MakeCoffee)WantAdd() bool {
	return true
}

func NewMakecoffee() *MakeCoffee{
	makecoffee := new(MakeCoffee)
	makecoffee.b = makecoffee
	return makecoffee
}

type MakeTee struct {
	Template
}

func (mt *MakeTee)BoilWater() {
	fmt.Println("MakeTee BoilWater.....")
}

func (mt *MakeTee)Brew() {
	fmt.Println("MakeTee Brew.....")
}

func (mt *MakeTee)PourInCup() {
	fmt.Println("MakeTee PourInCup.....")
}

func (mt *MakeTee)AddThings() {
	fmt.Println("MakeTee AddThings.....")
}

func (mt *MakeTee)WantAdd() bool {
	return false
}

func NewMakeTee() *MakeTee {
	mt := new(MakeTee)
	mt.b = mt
	return mt
}


// 业务逻辑层
func main() {
	makeCoffee := NewMakecoffee()
	makeCoffee.MakeBeverage()
	fmt.Println("------------")
	mt := NewMakeTee()
	mt.MakeBeverage()
}