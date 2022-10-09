/*
@Time : 2022/10/9 21:34
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 练习
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

// 抽象层
type AbstractGPU interface {
	Display()
}

type AbstractRAM interface {
	Storage()
}

type AbstractCPU interface {
	Calculate()
}

type AbstractFactory interface {
	CreateGPU() AbstractGPU
	CreateRAM() AbstractRAM
	CreateCPU() AbstractCPU
}
// 实现层
// Intel
type IntelCUP struct {}
func (ic *IntelCUP) Calculate() {
	fmt.Println("IntelCUP Calculate .....")
}
type IntelGPU struct {}
func (ig *IntelGPU)Display() {
	fmt.Println("IntelGPU Display .....")
}
type IntelRAM struct {}
func (ia *IntelRAM) Storage() {
	fmt.Println("IntelRAM Storage ....")
}

type IntelFactory struct {}
func (i * IntelFactory)CreateGPU() AbstractGPU {
	var inf AbstractGPU
	inf = new(IntelGPU)
	return inf
}
func (i * IntelFactory)CreateRAM() AbstractRAM{
	var inf AbstractRAM
	inf = new(IntelRAM)
	return inf
}
func (i * IntelFactory)CreateCPU() AbstractCPU{
	var inf AbstractCPU
	inf = new(IntelCUP)
	return inf
}
// nvidia
type NvidiaGPU struct {}
func (ng *NvidiaGPU)Display() {
	fmt.Println("NvidiaGPU Display...")
}
type NvidiaRAM struct {}
func (ng *NvidiaRAM)Storage() {
	fmt.Println("NvidiaRAM Storage...")
}
type NvidiaCPU struct {}
func (ng *NvidiaCPU)Calculate() {
	fmt.Println("NvidiaCPU Calculate...")
}
type NvidiaFactory struct {}
func (n *NvidiaFactory)CreateGPU() AbstractGPU {
	var nf AbstractGPU
	nf = new(NvidiaGPU)
	return nf
}
func (n *NvidiaFactory)CreateRAM() AbstractRAM {
	var nf AbstractRAM
	nf = new(NvidiaRAM)
	return nf
}
func (n *NvidiaFactory)CreateCPU() AbstractCPU {
	var nf AbstractCPU
	nf = new(NvidiaCPU)
	return nf
}
// Kingston

type KingstonGPU struct {}
func (ng *KingstonGPU)Display() {
	fmt.Println("KingstonGPU Display...")
}
type KingstonRAM struct {}
func (ng *KingstonRAM)Storage() {
	fmt.Println("KingstonRAM Storage...")
}
type KingstonCPU struct {}
func (ng *KingstonCPU)Calculate() {
	fmt.Println("KingstonCPU Calculate...")
}
type KingstonFactory struct {}
func (n *KingstonFactory)CreateGPU() AbstractGPU {
	var nf AbstractGPU
	nf = new(KingstonGPU)
	return nf
}
func (n *KingstonFactory)CreateRAM() AbstractRAM {
	var nf AbstractRAM
	nf = new(KingstonRAM)
	return nf
}
func (n *KingstonFactory)CreateCPU() AbstractCPU {
	var nf AbstractCPU
	nf = new(KingstonCPU)
	return nf
}

// 业务逻辑层
func main() {
	fmt.Println("  1台（Intel的CPU，Intel的显卡，Intel的内存）")
	var inf AbstractFactory
	inf = new(IntelFactory)
	var ing AbstractGPU
	ing = inf.CreateGPU()
	ing.Display()
	var inr AbstractRAM
	inr = inf.CreateRAM()
	inr.Storage()
	var inc AbstractCPU
	inc = inf.CreateCPU()
	inc.Calculate()

	fmt.Println("----------------")
	fmt.Println("  1台（Intel的CPU， nvidia的显卡，Kingston的内存）")
	inc.Calculate()
	nf := new(NvidiaFactory)
	ng :=nf.CreateGPU()
	ng.Display()
	kf := new(KingstonFactory)
	kr := kf.CreateRAM()
	kr.Storage()
}
