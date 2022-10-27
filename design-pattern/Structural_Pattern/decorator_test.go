/*
@Time : 2022/10/27 21:17
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : decorator_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Structural_Pattern

import "testing"

type ComponentDecorator interface {
	Calculation() int
}

type ConcreateComponentDecorator struct {}

func (*ConcreateComponentDecorator)Calculation() int {
	return 0
}

type ADecorator struct {
	ComponentDecorator
	num int
}

func NewADecorator(c ComponentDecorator, num int) ComponentDecorator {
	return &ADecorator{
		ComponentDecorator: c,
		num: num,
	}
}

func (a *ADecorator)Calculation() int {
	return a.ComponentDecorator.Calculation() + a.num
}

type MDecorator struct {
	ComponentDecorator
	num int
}

func NewMDecorator(c ComponentDecorator, num int) ComponentDecorator {
	return &MDecorator{
		ComponentDecorator: c,
		num: num,
	}
}

func (m *MDecorator) Calculation() int {
	return m.ComponentDecorator.Calculation() * m.num
}

func TestMyDecorator(t *testing.T) {
	var c ComponentDecorator
	c = &ConcreateComponentDecorator{}
	c = NewADecorator(c, 10)
	c = NewMDecorator(c, 20)
	res := c.Calculation()
	println(res)

}

