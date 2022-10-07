/*
@Time : 2022/10/7 17:13
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 工厂
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

type Fruit struct {

}

func NewFruit(name string) *Fruit{
	f := new(Fruit)
	if name == "apple" {

	} else if name == "banana" {

	}
	return f
}

func main() {

}