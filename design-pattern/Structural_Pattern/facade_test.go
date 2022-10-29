/*
@Time : 2022/10/29 20:18
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : facade_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Structural_Pattern

import (
	"fmt"
	"testing"
)

type SystemA struct {}

func (a *SystemA)methodA() {
	fmt.Println("method A ...")
}

type SystemB struct {}

func (b *SystemB) methodB() {
	fmt.Println("method B ...")
}

type SystemC struct {}

func (c *SystemC)methodC() {
	fmt.Println("method C ....")
}

type SystemD struct {}

func (d *SystemD) methodD () {
	fmt.Println("method D ...")
}

type Facade struct {
	a *SystemA
	b *SystemB
	c *SystemC
	d *SystemD
}

func (f *Facade)MethodOne() {
	f.a.methodA()
	f.b.methodB()
	f.c.methodC()
	f.d.methodD()
}

func (f *Facade)MethodTwo() {
	f.c.methodC()
	f.d.methodD()
	f.a.methodA()
	f.b.methodB()
}

func TestMyMethod(t *testing.T) {
	f := &Facade{
		a: new(SystemA),
		b: new(SystemB),
		c: new(SystemC),
		d: new(SystemD),
	}
	f.MethodOne()
	fmt.Println("------")
	f.MethodTwo()
}