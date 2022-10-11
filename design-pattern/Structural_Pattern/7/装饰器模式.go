/*
@Time : 2022/10/11 20:38
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 装饰器模式
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

// 抽象层
type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

func (d *Decorator)Show() {

}

// 实现层
type Huawei struct {

}

func (hw *Huawei)Show() {
	fmt.Println("huawei phone ....")
}
type Xiaomi struct {

}

func (xm *Xiaomi)Show() {
	fmt.Println("Xiaomi phone ....")
}

type MeDecorator struct {
	Decorator
}

func (md *MeDecorator)Show() {
	md.phone.Show()
	fmt.Println("me me ....")
}

func NewMeDecorator(phone Phone) Phone {
	return &MeDecorator{Decorator{phone: phone}}
}

type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator)Show() {
	kd.phone.Show()
	fmt.Println("ke ke ....")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}


// 业务逻辑
func main() {
	var huawei Phone
	huawei = new(Huawei)
	//huawei.Show()

	var mohua Phone
	mohua = NewMeDecorator(huawei)
	//mohua.Show()

	var kehua Phone
	kehua = NewKeDecorator(mohua)
	kehua.Show()


}
