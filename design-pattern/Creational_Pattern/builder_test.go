/*
@Time : 2022/10/18 21:23
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : builder_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package creation_pattern

import (
	"fmt"
	"testing"
)

type Builder interface {
	part1()
	part2()
	part3()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director)Contract() {
	d.builder.part1()
	d.builder.part2()
	d.builder.part3()
}

type builder1 struct {
	result string
}

func (b1 *builder1)part1() {
	b1.result += "1"
}

func (b1 *builder1)part2() {
	b1.result += "2"
}

func (b1 *builder1)part3() {
	b1.result += "3"
}

func (b1 *builder1)Get()string {
	return b1.result
}

type builder2 struct {
	result int
}
func (b2 *builder2)part1() {
	b2.result += 1
}

func (b2 *builder2)part2() {
	b2.result += 2
}

func (b2 *builder2)part3() {
	b2.result += 3
}

func (b2 *builder2)Get() int {
	return b2.result
}


func TestBuilder(t *testing.T) {
	b1 := &builder1{}
	d1 := NewDirector(b1)
	d1.Contract()
	ret1 := b1.Get()
	fmt.Println(ret1)

	b2 := &builder2{}
	d2 := NewDirector(b2)
	d2.Contract()
	ret2 := b2.Get()
	fmt.Println(ret2)
}