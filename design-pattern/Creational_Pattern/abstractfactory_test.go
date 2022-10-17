/*
@Time : 2022/10/17 21:29
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : abstractfactory
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package creation_pattern

import (
	"fmt"
	"testing"
)

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type AbstractProductA interface {
	CreateA()
}

type AbstractProductB interface {
	CreateB()
}

type ProductA1 struct {}

func (a1 *ProductA1) CreateA() {
	fmt.Println("CreateA 1....")
}

type ProductB1 struct {}

func (b1 *ProductB1) CreateB() {
	fmt.Println("CreateB 1....")
}

type ProductA2 struct {}

func (a2 *ProductA2)CreateA() {
	fmt.Println("CreateA 2....")
}

type ProductB2 struct {}

func (b2 *ProductB2) CreateB() {
	fmt.Println("CreateB 2....")
}

type ConcreteFactory1 struct {}

func (f1 *ConcreteFactory1) CreateProductA() AbstractProductA {
	var a AbstractProductA
	a = new(ProductA1)
	return a
}

func (f1 *ConcreteFactory1) CreateProductB() AbstractProductB {
	var b AbstractProductB
	b = new(ProductB1)
	return b
}

type ConcreteFactory2 struct {}

func (f2 *ConcreteFactory2) CreateProductA() AbstractProductA {
	var a AbstractProductA
	a = new(ProductA2)
	return a
}

func (f2 *ConcreteFactory2) CreateProductB () AbstractProductB {
	var b AbstractProductB
	b = new(ProductB2)
	return b
}

func TestAbstractFactory(t *testing.T) {
	var f1 AbstractFactory
	f1 = new(ConcreteFactory1)
	var a1 AbstractProductA
	a1 = f1.CreateProductA()
	a1.CreateA()
	t.Log("-----------")
	var f2 AbstractFactory
	f2 = new(ConcreteFactory2)
	var b2 = f2.CreateProductB()
	b2.CreateB()

}

