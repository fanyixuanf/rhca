/*
@Time : 2022/10/7 15:43
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : main
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

//type Clothes struct {}
//
//func (c *Clothes) OnWork(){
//	fmt.Println("工作装扮")
//}
//
//func (c *Clothes) OnShop() {
//	fmt.Println("逛街装扮")
//}
//
//func main() {
//	c := Clothes{}
//
//	c.OnWork()
//
//	c.OnShop()
//}

type ClothesWork struct {}

func (cw *ClothesWork) style() {
	fmt.Println("work...")
}

type ClothShop struct {}

func (cs *ClothShop) style() {
	fmt.Println("shop...")
}

func main() {

	cw := ClothesWork{}
	cw.style()

	cs := ClothShop{}
	cs.style()
}
