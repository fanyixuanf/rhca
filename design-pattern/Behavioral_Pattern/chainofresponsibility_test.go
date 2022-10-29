/*
@Time : 2022/10/29 21:18
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : chainofresponsibility_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type Manager interface {
	HaveRight(money int) bool
	HandlerFeeRequest(name string, money int) bool
}

type RequsetChain struct {
	Manager
	successor *RequsetChain
}

func (r *RequsetChain) SetSuccessor(m *RequsetChain) {
	r.successor = m
}

func (r *RequsetChain) HandlerFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		return r.Manager.HandlerFeeRequest(name, money)
	}
	return false
}

func (r *RequsetChain)HaveRight(money int) bool {
	return true
}

type ProjectManager struct {}

func (*ProjectManager)HandlerFeeRequest(name string, money int) bool {
	if name == "abc" {
		fmt.Printf("Project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Project manager don't permit %s %d fee request\n", name, money)
	return false
}

func (*ProjectManager)HaveRight(money int) bool {
	return money < 500
}

func NewProjectManagerChain() *RequsetChain {
	return &RequsetChain{
		Manager: &ProjectManager{},
	}
}

type DepManager struct {}

func (*DepManager)HaveRight(money int) bool {
	return money < 5000
}

func (*DepManager) HandlerFeeRequest(name string, money int) bool {
	if name == "abcd" {
		fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep manager don't permit %s %d fee request\n", name, money)
	return false
}

func NewDepManagerChain() *RequsetChain {
	return &RequsetChain{Manager: &DepManager{}}
}

type GeneralManager struct {}

func (*GeneralManager) HaveRight (money int) bool {
	return true
}

func (*GeneralManager) HandlerFeeRequest(name string, money int) bool {
	if name == "abcde" {
		fmt.Printf("General manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request\n", name, money)
	return false
}

func NewGeneralManagerChain() *RequsetChain {
	return &RequsetChain{Manager: &GeneralManager{}}
}

func TestChainofresponsibility(t *testing.T) {
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	var c Manager = c3

	c.HandlerFeeRequest("abc", 400)
	c.HandlerFeeRequest("abcd", 1000)
	c.HandlerFeeRequest("abcde", 10000)
	c.HandlerFeeRequest("text", 500)

}
