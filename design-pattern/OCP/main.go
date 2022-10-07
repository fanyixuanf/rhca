/*
@Time : 2022/10/7 15:58
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : main
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

//func main() {
//	banker := Banker{}
//	banker.Save()
//	banker.Save()
//	banker.Pay()
//	banker.Share()
//}
//
//type Banker struct {}
//
//func (b *Banker) Save() {
//	fmt.Println("save...")
//}
//
//func (b *Banker) Transfer() {
//	fmt.Println("transfer...")
//}
//
//func (b *Banker) Pay() {
//	fmt.Println("Pay...")
//}
//
//// (+)
//func (b *Banker) Share() {
//	fmt.Println("share...")
//}

func main() {
	//sb := SaveBanker{}
	//sb.DoSomething()
	//
	//tb := TransferBanker{}
	//tb.DoSomething()
	//
	//ss := ShareBanker{}
	//ss.DoSomething()

	BankBusiness(&SaveBanker{})
	BankBusiness(&TransferBanker{})
	BankBusiness(&ShareBanker{})

}

type AbstractBanker interface {
	DoSomething()
}

type SaveBanker struct {
	// AbstractBanker
}

func (sb *SaveBanker) DoSomething() {
	fmt.Println("save...")
}

type TransferBanker struct {}

func (tb *TransferBanker) DoSomething() {
	fmt.Println("transfer...")
}

// +++++++

type ShareBanker struct {}

func (ss *ShareBanker) DoSomething() {
	fmt.Println("Share...")
}

// 优化
func BankBusiness(banker AbstractBanker) {
	banker.DoSomething()
}
