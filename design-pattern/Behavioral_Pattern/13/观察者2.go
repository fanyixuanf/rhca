/*
@Time : 2022/10/13 21:48
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 观察者2
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

// 抽象层
type Event struct {
	Noti Notifier
	One Listener
	Another Listener
	Msg string
}

type Listener interface {
	OnFriendBeFight(event *Event)
	Title() string
	GetName() string
	GetParty() string
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify(event *Event)
}

// 实现层
type Hero struct {
	Name string
	Party string
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) GetParty() string {
	return h.Party
}

func (h *Hero) OnFriendBeFight(event *Event) {
	if event.One.GetName() == h.Name || event.Another.GetName() == h.Name{
		return
	}

	if h.Party == event.One.GetParty() {
		fmt.Println(h.Title(), "得知消息, 拍手....")
	}

	if h.Party == event.Another.GetParty() {
		fmt.Println(h.Title(), "得知消息, 复仇....")
		//h.Fight(event.One, event.Noti)
	}

}

func (h *Hero) Title() string {
	return fmt.Sprintf("[%s]%s", h.Party, h.Name)
}

func (h *Hero) Fight(an Listener, baixiao Notifier) {
	event := new(Event)
	event.Msg = fmt.Sprintf("%s 将 %s 揍了 ...", h.Title(), an.Title())
	event.Noti = baixiao
	event.One = h
	event.Another = an
	baixiao.Notify(event)

}



type Baixiao struct {
	heroList []Listener
}

func (bh *Baixiao)AddListener(listener Listener) {
	bh.heroList = append(bh.heroList, listener)
}

func (bh *Baixiao)RemoveListener(listener Listener) {
	for i, l := range bh.heroList {
		if listener == l {
			bh.heroList = append(bh.heroList[:i], bh.heroList[i + 1:]...)
		}
	}
}

func (bh *Baixiao)Notify(event *Event) {
	fmt.Println("[世界消息] 百晓生广播消息:", event.Msg)
	for _, l := range bh.heroList {
		l.OnFriendBeFight(event)
	}
}

// 业务逻辑层
func main() {
	h1 := Hero{
		Name:  "黄蓉",
		Party: "丐帮",
	}
	h2 := Hero{
		Name:  "洪七公",
		Party: "丐帮",
	}
	h3 := Hero{
		Name:  "乔峰",
		Party: "丐帮",
	}
	h4 := Hero{
		Name:  "张无忌",
		Party: "明教",
	}
	h5 := Hero{
		Name:  "灭绝师太",
		Party: "明教",
	}
	h6 := Hero{
		Name:  "金毛狮王",
		Party: "明教",
	}
	baixiao := Baixiao{}
	baixiao.AddListener(&h1)
	baixiao.AddListener(&h2)
	baixiao.AddListener(&h3)
	baixiao.AddListener(&h4)
	baixiao.AddListener(&h5)
	baixiao.AddListener(&h6)

	fmt.Println("----武林一片平静----")

	h1.Fight(&h4, &baixiao)

}