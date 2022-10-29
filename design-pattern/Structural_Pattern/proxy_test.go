/*
@Time : 2022/10/29 21:00
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : proxy_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Structural_Pattern

import (
	"fmt"
	"testing"
)

type Subject interface {
	MethodA()
	MethodB()
}

type realSubject struct {}

func (r *realSubject)MethodA() {
	fmt.Println("method a ......")
}

func (r *realSubject)MethodB() {
	fmt.Println("method b ......")
}

type Proxy struct {
	sub Subject
}

func (p *Proxy)MethodA() {
	fmt.Println("Proxy A ....")
	 p.sub.MethodA()
}

func (p *Proxy)MethodB() {
	fmt.Println("Proxy B ....")
	p.sub.MethodB()
}

func NewProxy(s Subject) Subject {
	return &Proxy{sub: s}
}

func TestMyProxy(t *testing.T) {
	var r Subject = new(realSubject)
	proxy := NewProxy(r)
	proxy.MethodA()
	proxy.MethodB()
}