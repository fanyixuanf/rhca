/*
@Time : 2022/10/31 20:35
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : state_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type Week interface {
	Today()
	Next(ctx *DayContext)
}

type DayContext struct {
	today Week
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

func NewDayContext() *DayContext {
	return &DayContext{today: &sunday{}}
}

type sunday struct {}

func (*sunday) Today() {
	fmt.Println("Sunday")
}

func (*sunday)Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

type Monday struct {}

func (*Monday) Today() {
	fmt.Println("Monday")
}

func (*Monday)Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

type Tuesday struct {}

func(*Tuesday)Today() {
	fmt.Println("Tuesday")
}

func (*Tuesday)Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

type Wednesday struct {}

func (*Wednesday) Today() {
	fmt.Println("Wednesday")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct {}

func (*Thursday) Today () {
	fmt.Println("Thursday")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct {}

func (*Friday) Today () {
	fmt.Println("Friday")
}

func (*Friday) Next (ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct {}

func (*Saturday) Today () {
	fmt.Println("Saturday")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = &sunday{}
}

func TestState(t *testing.T) {
	days := NewDayContext()
	today := func() {
		days.Today()
		days.Next()
	}

	for i := 0; i < 8; i++ {
		today()
	}

}