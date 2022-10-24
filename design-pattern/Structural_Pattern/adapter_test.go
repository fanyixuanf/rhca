/*
@Time : 2022/10/24 21:18
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : adapter_test.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Structural_Pattern

import (
	"fmt"
	"testing"
)

type Target interface {
	Request() string
}

type TargetImp struct {}

func (*TargetImp)Request() string {
	return "Target method"
}

func NewTarget() Target {
	return &TargetImp{}
}

type Adaptee interface {
	SpecificRequest() string
}

type AdapteeImp struct {}

func (*AdapteeImp) SpecificRequest() string {
	return "adaptee method"
}

func NewAdaptee() Adaptee {
	return &AdapteeImp{}
}

type Adapter struct {
	Adaptee
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{adaptee}
}

func TestAdapter(t *testing.T) {
	target := NewTarget()
	ret := target.Request()
	fmt.Println(ret)
	fmt.Println("-----------")
	adaptee := NewAdaptee()
	ret = adaptee.SpecificRequest()
	fmt.Println(ret)
	fmt.Println("-----------")
	adapter := NewAdapter(adaptee)
	ret = adapter.Request()
	fmt.Println(ret)

}
