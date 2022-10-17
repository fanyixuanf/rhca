/*
@Time : 2022/10/17 21:09
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : simplefactory
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package simplefactory

import (
	"fmt"
	"testing"
)

type Product interface {
	Use()
}

type A struct {
	Product
}

func (a *A)Use() {
	fmt.Println("AAAAAAAAAA...")
}

type B struct {
	Product
}

func (b *B)Use() {
	fmt.Println("BBBBBBBBB....")
}

type Factory struct {

}

func (f *Factory)CreateProduct(name string) Product {
	var product Product
	if name == "A" {
		product =  new(A)
	} else if name == "B" {
		product = new(B)
	}
	return product
}

func TestCreateProduct(t *testing.T) {
	factory := new(Factory)
	a :=factory.CreateProduct("A")
	a.Use()

	b := factory.CreateProduct("B")
	b.Use()
}
