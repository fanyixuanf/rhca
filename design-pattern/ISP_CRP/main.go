/*
@Time : 2022/10/7 16:53
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : main
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

type Cat struct {

}

func (c *Cat) Eat () {
	fmt.Println("Eat...")
}

// 添加方法（继承）

type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("Sleep...")
}

// 添加方法（组合）

type CatC struct {
	// cat *Cat
}

func (cc *CatC) Eat(c *Cat) {
	//cc.cat.Eat()
	c.Eat()
}

func (cc *CatC) Sleep () {
	fmt.Println("Sleep....")
}

func main () {
	cat := &Cat{}
	cat.Eat()

	fmt.Println(".........")
	cb := &CatB{}
	cb.Sleep()

	fmt.Println("--------------")
	cc := CatC{}
	cc.Eat(cat)
	cc.Sleep()

}