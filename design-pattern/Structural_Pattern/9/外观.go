/*
@Time : 2022/10/11 21:27
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 外观
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

type SubSysA struct {}

func (sa *SubSysA) MethodA() {
	fmt.Println("MethodA")
}

type SubSysB struct {}

func (sb *SubSysB) MethodB() {
	fmt.Println("MethodB")
}

type SubSysC struct {}

func (sc *SubSysC) MethodC() {
	fmt.Println("MethodC")
}

type SubSysD struct {}

func (sd *SubSysD) MethodD() {
	fmt.Println("SubSysD")
}

type Facade struct {
	a *SubSysA
	b *SubSysB
	c *SubSysC
	d *SubSysD
}

func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func main() {
	f := Facade{
		a: new(SubSysA),
		b: new(SubSysB),
		c: new(SubSysC),
		d: new(SubSysD),
	}
	f.MethodOne()
}
