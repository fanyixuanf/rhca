/*
@Time : 2022/10/10 21:20
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 代理模式
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

//type Goods struct {
//	Kind string
//	Fact bool
//}
//
//// ----> 抽象层 <----
//type Shoping interface {
//	Buy(goods *Goods)
//}
//
//type KoreaShopping struct {}
//
//func (k *KoreaShopping)Buy(goods *Goods) {
//	fmt.Println("KoreaShopping .....")
//}
//
//type AmericaShopping struct {}
//
//func (k *AmericaShopping)Buy(goods *Goods) {
//	fmt.Println("AmericaShopping .....")
//}
//type AfrikaShopping struct {}
//
//func (k *AfrikaShopping)Buy(goods *Goods) {
//	fmt.Println("AfrikaShopping .....")
//}
//
//type OverseasProsy struct {
//	shoping Shoping
//}
//
//func NewProxy(shoping Shoping) Shoping {
//	return &OverseasProsy{shoping: shoping}
//}
//
//func (op *OverseasProsy)Buy (goods *Goods) {
//	if op.checkGoods(goods) == true {
//		op.shoping.Buy(goods)
//		op.checkGoods2(goods)
//	}
//}
//
//func (op *OverseasProsy) checkGoods(goods *Goods) bool {
//	if goods.Fact == false {
//		fmt.Println("假货....")
//	}
//	return goods.Fact
//}
//
//func (op *OverseasProsy) checkGoods2(goods *Goods) {
//	fmt.Println("检查......")
//}
//
//
//// ----> 实现层 <----
//
//
//
//// ----> 业务逻辑层 <----
//func main() {
//	g1 := Goods{
//		Kind: "面膜",
//		Fact: true,
//	}
//	g2 := Goods{
//		Kind: "CET4",
//		Fact: false,
//	}
//
//	var kshoping = new(KoreaShopping)
//	var kproxy = NewProxy(kshoping)
//	kproxy.Buy(&g1)
//
//	var ashoping = new(AmericaShopping)
//	var aproxy = NewProxy(ashoping)
//	aproxy.Buy(&g2)
//	aproxy.Buy(&g1)
//
//}



// 抽象主题

type BeautyMoman interface {
	MakeEyesWithMan()
	HappyWithMan()
}

// 具体主题

type Panjinl struct {

}

func (P *Panjinl)MakeEyesWithMan() {
	fmt.Println("MakeEyesWithMan .....")
}
func (p *Panjinl)HappyWithMan() {
	fmt.Println("HappyWithMan ....")
}


type Wangp struct {
	woman BeautyMoman
}

func NewProxy(woman BeautyMoman) BeautyMoman {
	return &Wangp{woman: woman}
}

func (w *Wangp)MakeEyesWithMan() {
	w.woman.MakeEyesWithMan()
}
func (w *Wangp)HappyWithMan() {
	w.woman.HappyWithMan()
}


func main() {
	var pan BeautyMoman = new(Panjinl)
	 proxy := NewProxy(pan)
	 proxy.MakeEyesWithMan()
	 proxy.HappyWithMan()
}