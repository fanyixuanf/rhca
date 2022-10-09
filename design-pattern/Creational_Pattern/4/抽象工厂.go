/*
@Time : 2022/10/9 20:58
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 抽象工厂
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import (
	"fmt"
)

// ----> 抽象层 <----
type AbstractApple interface {
	ShowApple()
}
type AbstractBanana interface {
	ShowBanana()
}
type AbstractPear interface {
	ShowPear()
}

type AbstractFactroy interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// ----> 实现层 <----
 // china
type ChinaApple struct {}
func (ca *ChinaApple)ShowApple() {
	fmt.Println("china apple")
}
type ChinaBanana struct {}

func (cb *ChinaBanana)ShowBanana() {
	fmt.Println("china Banana")
}

type ChinaPear struct {}

func (cp *ChinaPear)ShowPear() {
	fmt.Println("china Pear")
}

type ChinaFactory struct {}

func (cf *ChinaFactory)CreateApple() AbstractApple {
	var app AbstractApple
	app = new(ChinaApple)
	return  app
}

func (cf *ChinaFactory)CreateBanana() AbstractBanana {
	var ba AbstractBanana
	ba = new(ChinaBanana)
	return  ba
}

func (cf *ChinaFactory)CreatePear() AbstractPear {
	var pe AbstractPear
	pe = new(ChinaPear)
	return  pe
}

// japan

type JapanApple struct {}
func (ja *JapanApple)ShowApple() {
	fmt.Println("Japan Apple")
}
type JapanBanana struct {}
func (ja *JapanBanana)ShowBanana() {
	fmt.Println("Japan Banana")
}

type JapanPear struct {}
func (ja *JapanPear)ShowPear() {
	fmt.Println("Japan Pear")
}

type JapanFactory struct {}

func (jf *JapanFactory)CreateApple() AbstractApple {
	var ja AbstractApple
	ja = new(JapanApple)
	return ja
}

func (jf *JapanFactory)CreateBanana() AbstractBanana {
	var jb AbstractBanana
	jb = new(JapanBanana)
	return jb
}

func (jf *JapanFactory)CreatePear() AbstractPear {
	var jp AbstractPear
	jp = new(JapanPear)
	return jp
}

// ----> 业务逻辑 <----
func main1() {

	var cf AbstractFactroy
	cf = new(ChinaFactory)

	var capp AbstractApple
	capp = cf.CreateApple()
	capp.ShowApple()

	var cban AbstractBanana
	cban = cf.CreateBanana()
	cban.ShowBanana()

	var cp AbstractPear
	cp = cf.CreatePear()
	cp.ShowPear()

}
