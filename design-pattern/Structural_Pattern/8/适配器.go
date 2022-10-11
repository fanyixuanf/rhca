/*
@Time : 2022/10/11 20:59
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 适配器
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

//type V220 struct {}
//
//func (v2 V220) Use220() {
//	fmt.Println("use 220v 充电")
//}
//
//type V5 interface {
//	Use5v()
//}
//
//type Phone struct {
//	v V5
//}
//
//func NewPhone(v V5) *Phone{
//	return &Phone{v: v}
//}
//
//type Adapter struct {
//	v2 *V220
//}
//
//func (a *Adapter) Use5v() {
//	fmt.Println("Adapter ....")
//	a.v2.Use220()
//}
//
//func NewAdapter(v220 *V220) *Adapter {
//	return &Adapter{v2: v220}
//}
//
//func (p *Phone) Charge() {
//	fmt.Println("Phone 充电")
//	p.v.Use5v()
//}
//
//func main() {
//	phone := NewPhone(NewAdapter(new(V220)))
//	phone.Charge()
//}



type Attack interface {
	Fight()
}

type Hero struct {
	Name string
	attack Attack
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能")
	h.attack.Fight()
}

type Dabaojian struct {

}

func (dd *Dabaojian)Fight() {
	fmt.Println("使用了大保健")
}

type Poweroff struct {}

func (p *Poweroff)Shutdown() {
	fmt.Println("shut down .....")
}

type Adapter struct {
	Poweroff *Poweroff
}

func (a *Adapter)Fight() {
	fmt.Println("Adapter .....")
	a.Poweroff.Shutdown()
}

func NewAdapter(power *Poweroff) *Adapter {
	return &Adapter{Poweroff: power}
}

func main() {
	gai := Hero{
		Name: "gailun",
		attack: new(Dabaojian),
	}
	gai.Skill()

	fmt.Println("----------")
	gai1 := Hero{
		Name:   "gay",
		attack: NewAdapter(new(Poweroff)),
	}
	gai1.Skill()
}
