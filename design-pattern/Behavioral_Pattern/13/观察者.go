/*
@Time : 2022/10/13 21:23
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 观察者
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

//import "fmt"
//
//// 抽象
//type Listener interface {
//	OnTeacherComing()
//	DoBadThing()
//}
//
//type Notifier interface {
//	AddListener(listener Listener)
//	RemoveListener(listener Listener)
//	Notify()
//	NotifyBadThing()
//}
//
//// 实现
//
//type Zhang3 struct {
//	BadThing string
//}
//
//func (z3 *Zhang3) OnTeacherComing() {
//	fmt.Println("zhang3 停止....",z3.BadThing)
//}
//
//func (z3 *Zhang3) DoBadThing() {
//	fmt.Println("zhang3 开始....",z3.BadThing)
//}
//
//type Zhao4 struct {
//	BadThing string
//}
//
//func (z4 *Zhao4) OnTeacherComing() {
//	fmt.Println("zhao4 停止...",z4.BadThing)
//}
//
//func (z4 *Zhao4) DoBadThing() {
//	fmt.Println("zhao4 开始...",z4.BadThing)
//}
//
//type Wang5 struct {
//	BadThing string
//}
//
//func (w5 *Wang5) OnTeacherComing() {
//	fmt.Println("wang5 停止....", w5.BadThing)
//}
//
//func (w5 *Wang5) DoBadThing() {
//	fmt.Println("wang5 开始....", w5.BadThing)
//}
//
//type ClassMoitor struct {
//	listener []Listener
//}
//
//func (cm *ClassMoitor) AddListener(listener Listener) {
//	cm.listener = append(cm.listener, listener)
//}
//
//func (cm *ClassMoitor) RemoveListener(listener Listener) {
//	for i, l := range cm.listener {
//		if listener == l {
//			cm.listener = append(cm.listener[:i], cm.listener[i+1:]...)
//		}
//	}
//}
//
//func (cm *ClassMoitor) Notify() {
//	for _, l := range cm.listener {
//		l.OnTeacherComing()
//	}
//}
//
//func (cm *ClassMoitor) NotifyBadThing() {
//	for _, l := range cm.listener {
//		l.DoBadThing()
//	}
//}
//
//// 业务逻辑
//func main1() {
//	z3 := &Zhang3{BadThing: "打架"}
//	z4 := &Zhao4{BadThing: "打游戏"}
//	w5 := &Wang5{BadThing: "吹牛"}
//	m := ClassMoitor{}
//	m.AddListener(z3)
//	m.AddListener(z4)
//	m.AddListener(w5)
//	fmt.Println("----干坏事----")
//	m.NotifyBadThing()
//	fmt.Println("----停止干坏事-----")
//	m.Notify()
//}
